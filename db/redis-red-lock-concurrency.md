# Redis를 활용해 동시성 문제 해결하기

- 해당 문서에서는 Redis와 동시성에 대해 아래 3 가지 방법을 중점적으로 작성한다.
  - 동시성 문제를 재현한다.
  - Redis의 set nx 명령어를 이용해 동시성 문제를 해결한다.
  - Redis의 Red Lock을 이용해 동시성 문제를 해결한다.
- 문서에 나오는 전체 코드는 [여기서](https://github.com/programmer-sjk/nestjs-redis) 확인할 수 있다.

## 동시성 문제

- 동시성 문제란 여러 쓰레드들이 공유 자원에 대한 경쟁을 벌여 실행 순서에 따라 의도하지 않은 결과를 뜻 한다.

### 동시성 문제 확인을 위한 함수 준비

- 테스트 해 볼 함수는 movie service에 영화의 추천 수를 업데이트 하는 함수이다.
- id에 해당하는 영화를 조회하고, 그 영화의 추천 수에 1을 증가시킨다.

  ```ts
    async increaseRecommendCount(id: number) {
      const movie = await this.movieRepository.findOne(id);
      await this.movieRepository.updateRecommendCount(
        id,
        movie.recommendCount + 1
      );
    }
  ```

### 동시성 문제 확인

- 동시성 문제를 확인하기 위해 아래와 같은 테스트 코드를 준비했다.
- 1번 영화의 추천 수를 증가시키는 기능을 Promise.all을 이용해 비동기로 함수를 10번 호출한다.

  ```ts
  describe('MovieService', () => {
    it('동시에 10개 요청', async () => {
      // given
      // DB에 추천 수가 0인 1번 영화를 수동으로 만들어 둠

      // when
      await Promise.all([
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
        service.increaseRecommendCount(1),
      ]);

      // then
      const movie = await movieRepository.findOne(1);
      expect(movie.recommendCount).toBe(10);
    });
  });
  ```

- 기대하는 결과는 순차적으로 실행되어 10을 바랬지만 실제 결과는 아래와 같다.
- 테스트를 여러 번 수행할 때 마다 결과는 조금씩 달랐다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/concurrency-test-fail.png" width="500">

### 동시성 문제가 발생한 이유

- 간절히 바라고 소망하고 염원했던 결과가 나오질 않았다. 왜 그랬는지 알아보자.
- 아마 우리는 아래와 같은 흐름처럼 쓰레드 1에서 조회 -> 업데이트 후 쓰레드 2가 작업을 수행하길 바랬을 것이다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/expect-db-concurrency.png" width="650">

- 하지만 실제로는 동시에 접근할 경우 아래와 같은 흐름이 충분히 발생할 수 있다.

  - 쓰레드 1에서 업데이트를 하기 전에 쓰레드 2가 현재의 추천 수 0을 획득 한다.
  - 쓰레드 1에서 추천 수를 1로 증가시켰지만 쓰레드 2도 마찬가지로 1로 업데이트 한다.

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/real-db-concurrency.png" width="650">

- DB의 Lock을 이용해 동시성을 해결할 수 있지만, 여기서는 Redis만 활용해 보기로 한다.

## Redis set에 NX 옵션을 활용하여 동시성 문제 해결하기

- Redis는 싱글 스레드로 동작하기 때문에 어떤 시점에 Redis에 접근해 작업을 수행할 수 있는 쓰레드는 1개 뿐이다.
- Redis key에 값을 set 할 때 NX 옵션을 줄 수 있는데, 이 옵션은 key가 없을 때만 set을 할 수 있다.

### redis-cli를 통해 NX 동작 확인

- 아래 명령어를 보면 myKey가 없었던 처음 경우에만 set 명령어가 성공한다.
- 두 번째 set에 NX 옵션을 주면 myKey가 존재하므로 명령이 정상적으로 수행되지 않는다.

  ```console
    127.0.0.1:6379> GET myKey
    (nil)
    127.0.0.1:6379> SET myKey value NX
    OK
    127.0.0.1:6379> GET myKey
    "value"
    127.0.0.1:6379> SET myKey updateValue NX
    (nil)
    127.0.0.1:6379> GET myKey
    "value"
  ```

### NestJS에서 set NX 옵션을 활용해 동시성 문제 해결하기

- 우선 Redis service에 setNx와 del 메서드를 준비한다.

  - `PX 1000`은 1000ms 다음에 expire됨을 뜻한다.

  ```ts
  @Injectable()
  export class RedisService {
    constructor(@InjectRedis() private redis: Redis) {}

    async setNx(key: string, value: string) {
      return this.redis.set(key, value, 'PX', 1000, 'NX');
    }

    async del(key: string) {
      return this.redis.del(key);
    }
  }
  ```

- 다음으로 영화 추천 수 업데이트에 set NX 명령어를 적용한 `increaseRecommendCountBySetNx` 함수를 작성한다.

  - set NX 명령을 성공하면 조회 / 업데이트 로직을 수행 후 키를 삭제한다.
  - set NX 명령에 실패한다면 100ms 후에 set NX 명령을 다시 시도한다.

    ```ts
    @Injectable()
    export class MovieService {
      constructor(
        private readonly redisService: RedisService,
        private readonly movieRepository: MovieRepository
      ) {}

      async increaseRecommendCountBySetNx(id: number) {
        const key = 'cacheKey';
        while (!(await this.redisService.setNx(key, 'value'))) {
          sleep(100);
        }

        const movie = await this.movieRepository.findOne(id);
        await this.movieRepository.updateRecommendCount(
          id,
          movie.recommendCount + 1
        );

        await this.redisService.del(key);
      }
    }
    ```

- 동시성 문제가 발생했던 것 처럼 Promise.all을 활용해 함수를 테스트 한다.

  ```ts
  it('set nx 동시에 10개 요청', async () => {
    // given
    // DB에 추천 수가 0인 1번 영화를 수동으로 만들어 둠

    // when
    await Promise.all([
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
      service.increaseRecommendCountBySetNx(1),
    ]);

    // then
    const movie = await movieRepository.findOne(1);
    expect(movie.recommendCount).toBe(10);
  });
  ```

- Redis에는 한 쓰레드만 접근이 가능하므로 아래와 같이 테스트 결과가 성공한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/concurrency-test-success-nx.png" width="500">

## RedLock을 활용해 동시성 문제 해결하기

- RedLock 알고리즘은 N대의 독립적인 Redis Node가 분산된 환경에서 Lock을 획득 및 해제하는 방법이다.
- Red Lock 개념에 대한 설명은 이 [문서](./redis-red-lock.md)를 참고하길 바란다.

### NestJS에서 RedLock을 사용해 동시성 문제 해결하기

- [Redis 공식문서](https://redis.io/docs/manual/patterns/distributed-locks/#implementations)를 보면 Node 구현으로 [node-redlock](https://github.com/mike-marcacci/node-redlock)을 언급하고 있다.
- `yarn add redlock` 명령어로 redlock을 설치한다.
- Redis Service에 redlock 알고리즘을 사용하려 락을 획득하는 메서드를 제공한다.

  ```ts
  import Redlock from 'redlock';

  @Injectable()
  export class RedisService {
    private readonly redlock: Redlock;
    private readonly lockDuration = 10_000;

    constructor(@InjectRedis() private redis: Redis) {
      this.redlock = new Redlock([redis]);
    }

    async acquireLock(key: string) {
      return this.redlock.acquire([`lock:${key}`], this.lockDuration);
    }
  }
  ```

- redlock을 활용해 영화의 추천 수를 업데이트하는 `increaseRecommendCountByRedLock` 함수를 추가한다.

  - `acquireLock` 메서드로 락을 획득한 쓰레드만 업데이트 후 락을 해제한다.
  - 락을 획득하지 못한 쓰레드는 일정 시간 대기 후 락을 획득하기 위한 재 시도를 요청한다.

    ```ts
    async increaseRecommendCountByRedLock(id: number) {
      let lock: Lock;
      try {
        lock = await this.redisService.acquireLock(
          `increase-recommend-count:${id}`
        );

        const movie = await this.movieRepository.findOne(id);
        await this.movieRepository.updateRecommendCount(
          id,
          movie.recommendCount + 1
        );
      } catch (err) {
        throw err;
      } finally {
        await lock.release();
      }
    }
    ```

- 위 함수를 부르는 테스트 코드를 작성한다.

  ```ts
  it('redlock 동시에 10개 요청', async () => {
    // given
    // DB에 추천 수가 0인 1번 영화를 수동으로 만들어 둠

    // when
    await Promise.all([
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
      service.increaseRecommendCountByRedLock(1),
    ]);

    // then
    const movie = await movieRepository.findOne(1);
    expect(movie.recommendCount).toBe(10);
  });
  ```

- lock을 획득한 쓰레드만 조회 및 업데이트를 동기적으로 수행해 성공한 것을 확인할 수 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/concurrency-test-success-lock.png" width="500">
