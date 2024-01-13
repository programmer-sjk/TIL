# Redis Red Lock

- 여러 프로세스가 공유된 자원에 접근하면 동시성 문제가 발생하기 마련이다.
- 동시성 이슈를 해결하기 위한 몇 가지 방안들이 있지만 이 문서에는 Redis가 제공하는 RedLock에 대해 정리한다.
- RedLock은 Redis가 제공하는 분산 락(Distributed Lock)이다.

## 왜 하필 RedLock 인가?

- Redis는 싱글 스레드로 동작하기 떄문에 아래의 명령어로도 잠금을 획득할 수 있다.
  - 아래 명렁어는 `key`가 없을 경우 value를 저장하고 30초 동안 유지한다.
  - `SET key value NX PX 30000`
- 따라서 동시성 문제가 발생하지 않는데 왜 분산락이 필요할까?
  - 만약 Redis가 하나의 서버로 동작할 경우 단일 장애 지점(SPOF)이 될 수 있다.
- 이를 위해 복제 모드(Master - Slave)로 구성할 수 있다. 하지만 Redis의 복제는 비동기라 문제가 발생할 수 있다.
  - 클라이언트 A가 마스터에서 잠금을 획득한다.
  - 키에 대한 쓰기가 복제본으로 전송되기 전 Master가 다운된다.
  - 키가 쓰여지지 않은 Slave가 Master로 승격한다.
  - 클라이언트 B가 새로운 마스터에서 동일한 키로 잠금을 획득한다.
- 이런 문제를 보완하기 위해 RedLock 알고리즘이 제안되었다.

## RedLock 알고리즘

- 위에서 언급한 문제를 보완하기 위해 Redis는 RedLock 알고리즘을 제안했다.
- 레드락은 N개의 레디스 노드들을 이용해 Quorum 이상의 노드에서 잠금을 획득하면 분산락을 획득한 것으로 판단한다
- 클라이언트가 분산 환경에서 락을 획득하기 위한 절차는 아래와 같다.
  - 현재 시간을 밀리초 단위로 가져온다.
  - 모든 인스턴스에서 동일한 키와 임의의 값을 사용하여 모든 N 인스턴스에서 순차적으로 잠금을 획득하려고 시도한다.
    - 각 인스턴스에 잠금을 설정할 때 클라이언트는 잠금을 획득하기 위해 잠금 자동 해제 시간에 비해 작은 제한 시간을 사용한다.
    - 예를 들어 자동 해제 시간이 10초인 경우 제한 시간은 ~ 5~50밀리초 범위일 수 있다.
    - 이렇게 하면 클라이언트가 다운된 Redis 노드와 통신하려고 오랫동안 차단되는 것을 방지할 수 있다.
  - 클라이언트는 잠금을 획득하는 데 경과된 시간(현재시간 - 1단계에서 구한 시간)을 계산한다.
    - 클라이언트가 Quorum 이상의 노드에서 잠금을 획득할 수 있었던 경우에만 해당된다.
    - 잠금을 획득하는데 소요된 총 시간이 잠금 유효 시간보다 작을 경우 잠금을 획득한 것으로 간주한다.
  - 잠금이 획득된 경우 잠금의 유효 시간은 전달한 타임아웃 - 3단계에서 구한 시간으로 간주한다.
  - 클라이언트가 Quorum 이상의 노드에서 잠금 획득을 실패한 경우, 모든 인스턴스의 잠금을 해제한다.
- 공식문서에 적힌 절차를 썼지만, 과반수 이상의 노드에서 잠금을 획득하면 잠금을 획득한다고 간주하는 알고리즘이다.

## RedLock은 완벽하지 않다

- 공식 문서에 언급된 `Martin Kleppmann`이 분석한 문서에 따르면 RedLock 알고리즘에도 문제가 발생할 수 있다.
- 아래와 같이 lock을 획득하고 파일에 데이터를 저장하는 코드가 있다고 가정하자

  ```js
  // THIS CODE IS BROKEN
  function writeData(filename, data) {
    var lock = lockService.acquireLock(filename);
    if (!lock) {
      throw "Failed to acquire lock";
    }

    try {
      var file = storage.readFile(filename);
      var updated = updateContents(file, data);
      storage.writeFile(filename, updated);
    } finally {
      lock.release();
    }
  }
  ```

- 위 코드에서 발생할 수 있는 문제를 시각화 한 다이어그램과 절차는 아래와 같다.

  - 클라이언트 A가 마스터에서 잠금을 획득한다.
  - 클라이언트 A에서 Stop-the-World GC로 인한 어플리케이션 코드 중지가 발생하고 그 사이에 잠금이 만료된다.
  - 클라이언트 B가 분산락을 획득하고 파일에 데이터를 쓴다.
  - 클라이언트 A가 GC가 끝난 후 파일에 데이터를 쓰면서 동시성 문제가 발생한다

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/broken-red-lock.png" width="600">

- 일반적으로 GC는 매우 빠르게 수행되지만 Stop-the-World GC는 드물게 잠금이 만료될 정도로 지속될 수 있다.
- GC 외에도 여러 이유들로 동일한 문제가 발생할 수 있다.

## 코드로 보는 NestJS에서 RedLock 사용

- 이 섹션에서는 전체 절차를 정리하지 않고 사용 예시로만 간단하게 정리하겠다.
- 사용하는 라이브러리는 찾아보니 [node-redlock](https://github.com/mike-marcacci/node-redlock)이 많이 사용된다.
- 특정 서비스 Layer에서 Redis의 redlock을 사용한다고 하면 코드는 아래와 같다.

```ts
// RedisService
@Injectable()
export class RedisService {
  private readonly redlock: Redlock;
  private readonly lockDuration = 10_000; // 10 seconds (unit: ms)

  constructor(@InjectRedis() private readonly redis: Redis) {
    this.redlock = new Redlock([this.redis]);
  }

  async acquireLock(key: string) {
    return this.redlock.acquire([`lock:${key}`], this.lockDuration);
  }
}

// lock을 획득하여 사용하는 서비스 코드
@Injectable()
export class ReviewService {
  constructor(private readonly redisService: RedisService) {}

  async doSomething(id: number) {
    let lock: Lock;
    try {
      lock = await this.redisService.acquireLock(`do-something:${id}`);
    } catch (err) {
      throw new Error("잠금 획득 실패");
    }

    // 분산 환경에서 한 프로세스만 실행해야 하는 코드 실행
    doSomething(id);

    await lock.release().catch(() => undefined);

    // 후 처리 코드
    return true;
  }
}
```

## 정리

- 분산 환경에서 동시성을 제어하기 위한 방안으로 Redis의 RedLock을 사용할 수 있다.
- 다만 RedLock은 완벽하지 않으며 문제가 발생할 수 있음을 인지하고 있어야 한다.
