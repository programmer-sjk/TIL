# MySQL Lock

- DB에서 Lock은 동시성을 제어하기 위해 사용한다.
- 다수의 트랜잭션에서 동시에 데이터를 조회, 수정하기 위해 접근할 때 락을 획득한 커넥션이 작업이 우선 처리된다.

## 비관적 락 vs 낙관적 락

- 락의 종류는 크게 비관적 락과 낙관적 락으로 나뉜다.
- 동시성을 비관적으로 보고 실제로 락을 잡는 비관적 락과, 낙관적으로 보고 락을 사용하지 않는 낙관적 락이다.
- 비관적 락은 실제로 공유 락이나 배타 락을 잡아서 동시성을 제어하는 방법이다. 두 락에 대해서는 아래에서 더 알아본다.
- 낙관적 락은 작업을 수행하고 충돌이 발생했다면 후에 실행된 쿼리를 실패시키는 방법이다.
- 주로 version 컬럼을 하나 추가하고 작업을 수행 후, version 값을 검사하는 방식을 많이 사용한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/optimistic-lock.png" width="400">

- 위 그림을 살펴보자.
  - 사용자 A, B가 동시에 트랜잭션을 시작하고 이때 버전 값은 1이다.
  - A의 작업이 먼저 commit 되어 2로 수정되면 A의 경우 버전 값이 정상적으로 1 증가했기에 정상 종료된다.
  - B의 작업이 이후 실행되며 처음 읽은 version 값 1이 수정되었기 때문에 실패한다.
- 경쟁이 치열한 환경에서 낙관적 락은 vesrion 값 체크와 rollback 작업이 빈번해 질 수 있어 성능이 좋지 않다.
- 따라서 동시성이 높다면 비관적 락을, 동시성이 낮다고 판단하면 낙관적 락을 고려해볼 수 있다.

## 공유 락 vs 배타 락

### 공유 락

- 공유 락은 읽기를 위해 잡는 락으로 다른 트랜잭션에서 공유 락을 잡을 순 있지만 수정을 위한 배타 락을 잡을 순 없다. 이를 달리 표현하면 공유 락을 잡는 순간 해당 데이터는 락이 해제되기 전까지 수정이 불가능하다.
- 아래와 같이 두 개의 DB 커넥션에서 작업을 수행해보면 어떻게 동작하는지 이해할 수 있을 것이다.

  ```sql
    // 트랜잭션 A
    begin;
    SELECT * FROM users WHERE id = 1 FOR SHARE;

    // 트랜잭션 B
    SELECT * FROM users WHERE id = 1; -- 가능
    SELECT * FROM users WHERE id = 1 FOR SHARE; -- 가능

    UPDATE user SET name = 'aaa' WHERE id = 1; -- 대기
    SELECT * FROM users WHERE id = 1 FOR UPDATE; -- 대기
  ```

- 배타 락을 잡거나 데이터의 수정은 트랜잭션 A에서 commit이나 rollback을 하게 되면 수행된다.

## AUTO INCREMENT 락

## InnoDB 락

- 레코드 락
- 갭락
- 넥스트 키 락
