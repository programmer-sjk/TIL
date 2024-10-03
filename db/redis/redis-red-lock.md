# Redis Red Lock

- **RedLock**은 분산 환경에서 Redis가 권장하는 **`분산 락(Distributed Lock)이다`**.
- 이 문서에서는 `다음 두 가지`를 중점적으로 정리한다.
  - Redis `Set 명령어에 NX 옵션`을 통한 Lock을 제공하는 방법과 한계
  - `RedLock`의 동작방식 및 한계에 대해 정리하겠다.

## Redis SET NX

- Redis는 2.6.12 버전 이전에는 `SETNX` 명령어를 제공했다.
  - 2.6.12 버전부터 **`SETNX 명령어 자체는 deprecated`** 하고 SET 명령어에 NX 옵션을 전달하는 방향으로 수정하였다.
- NX 옵션을 전달하면 SET 하려는 **`키가 없는 경우에만 SET이 성공한다`**.
  - Redis는 **싱글 스레드로 동작**하기 때문에 여러 프로세스가 공유 자원에 접근할 때 발생하는 동시성 문제를 이 명령어로 해결할 수 있다.
- 즉 먼저 접근한 쓰레드가 NX 옵션을 전달한 SET에 성공한다면 다른 쓰레드들은 대기한다.
  - 여기서 다른 쓰레드들이 대기하도록 while 문과 같은 루프와 Sleep같은 함수는 개발자가 직접 제공해야 한다.
- Lock을 획득해 먼저 자원을 선점한 쓰레드는 **`작업을 끝낸 후 키를 삭제한다`**.

  - 여기서 단순히 DEL 명령어로 키를 삭제하면 **`Lock을 획득하지 않은 다른 클라이언트들도 삭제가 가능하다`**.
  - 따라서 Key가 존재하고 값이 일치할 때만 삭제할 수 있도록 아래와 같은 **Lua 스크립트를** 통해 삭제할 것을 권고한다.

    ```lua
      if redis.call('get', KEYS[1]) == ARGV[1] then
        return redis.call('del', KEYS[1])
      else
        return 0
      end
    ```

- 여기서 **한 가지 문제점**은 Redis가 단일 서버로 동작한다면 **`단일 장애 지점(SPOF)`** 될 수 있다.
- 이를 위해 `Master-Slave`의 복제 구조를 고려 할 수 있지만 **`복제 구조가 비 동기라`** 아래와 같은 문제가 발생할 수 있다.
  - 클라이언트 A가 마스터에서 Lock을 획득한다.
  - 키에 대한 쓰기가 복제본으로 전송되기 전 Master가 다운된다.
  - 키가 쓰여지지 않은 Slave가 Master로 승격한다.
  - 클라이언트 B가 새로운 마스터에서 동일한 키로 Lock을 획득한다.
- 이런 문제를 보완하기 위해 **`RedLock`** 알고리즘이 제안되었다.

## RedLock 알고리즘

- 레드락은 N개의 레디스 노드들을 이용해 **`과반 수 이상의 노드에서 잠금을 획득하면`** 분산락을 획득한 것으로 판단한다
- 분산 환경에서 락을 획득하기 위한 **구체적인 절차**는 아래와 같다.
  - 현재 시간을 밀리초 단위로 가져온다.
  - 순차적으로 **`N대의 Redis 서버에 잠금을 요청한다`**. 이 때 timeout은 Lock의 유효시간 보다 훨씬 작은시간을 쓴다. 만약 Lock의 유효시간이 10초라면 각 Redis 서버에 잠금을 획득하기 위한 timeout은 5~50ms 이다. 이렇게 짧은 timeout을 사용해 장애가 발생한 Redis 서버와 통신에 많은 시간을 사용하지 않도록 방지할 수 있다.
  - Redis 서버가 5대라고 가정할 때 과반수(3대) 이상의 서버로부터 Lock을 획득했고 Lock을 획득하기 위해 사용한 시간이 Lock의 유효시간보다 작았다면 Lock을 획득했다고 간주한다. 즉 Lock의 유효시간이 10초인데 Lock을 얻기 위해 11초가 걸렸다면 실패한 것이다.
  - Lock을 획득 한 후 유효시간은 처음 **`Lock의 유효시간 - Lock을 얻기 위해 걸린 시간이다`**. Lock의 유효시간이 10초인데 획득에 3초가 걸렸다면 얻은 후부터 7초 뒤에 만료된다.
  - **`Lock을 얻지 못했다면 모든 Redis 서버에게 Lock 해제 요청을 보낸다`**. 예를 들어 5대의 Redis 서버 중 한 대의 서버에게만 Lock 획득을 성공했다. 과반 수 이상의 Lock을 획득하지 못했으므로 Lock 획득에 실패했고 모든 Redis 서버에 Lock 해제 요청을 보낸다.
- 공식문서에 적힌 절차를 썼지만, 쉽게 풀면 **분산 환경에서 과반수 이상의 Redis 노드에서 잠금을 획득하면 잠금을 획득한다고 간주하는 알고리즘**이다.

## RedLock은 완벽하지 않다

- 공식 문서에 언급된 `Martin Kleppmann`이 분석한 문서에 따르면 RedLock 알고리즘에도 문제가 발생할 수 있다.
- 아래와 같이 lock을 획득하고 파일에 데이터를 저장하는 코드가 있다고 가정하자

  ```js
  // THIS CODE IS BROKEN
  function writeData(filename, data) {
    var lock = lockService.acquireLock(filename);
    if (!lock) {
      throw 'Failed to acquire lock';
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
  - 클라이언트 A에서 **Stop-the-World GC**로 인한 어플리케이션 코드 중지가 발생하고 그 사이에 잠금이 만료된다.
  - 클라이언트 B가 분산락을 획득하고 파일에 데이터를 쓴다.
  - 클라이언트 A가 GC가 끝난 후 파일에 데이터를 쓰면서 동시성 문제가 발생한다

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/broken-red-lock.png" width="600">

- 일반적으로 GC는 매우 빠르게 수행되지만 **Stop-the-World** GC는 **`드물게 잠금이 만료될 정도로 지속될 수 있다`**.
- GC 외에도 여러 이유들로 동일한 문제가 발생할 수 있다.

## 코드로 보는 NestJS에서 RedLock 사용

- 원래 코드를 넣었다가, 코드 위주로 동시성 문제를 해결하는 문서를 따로 작성했다.
- [이 문서](./redis-red-lock-concurrency.md)에 NestJS에서 RedLock을 사용한 코드를 확인할 수 있다.

## 정리

- 분산 환경에서 동시성을 제어하기 위한 방안으로 Redis의 RedLock을 사용할 수 있다.
- 다만 RedLock은 완벽하지 않으며 문제가 발생할 수 있음을 인지하고 있어야 한다.

## 레퍼런스

- <https://redis.io/docs/manual/patterns/distributed-locks/>
- <https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html>
