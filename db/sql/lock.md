# MySQL Lock

- DB에서 Lock은 동시성을 제어하기 위해 사용한다.
- 다수의 트랜잭션에서 동시에 데이터를 조회, 수정하기 위해 접근할 때 락을 획득한 커넥션이 작업이 우선 처리된다.

## 비관적 락 vs 낙관적 락

- 락의 종류는 크게 **`비관적 락과 낙관적 락으로 나뉜다`**.
- **`동시성을 비관적으로 보고 실제로 락을 잡는 비관적 락과, 낙관적으로 보고 락을 사용하지 않는 낙관적 락이다`**.
- 비관적 락은 실제로 공유 락이나 배타 락을 잡아서 동시성을 제어하는 방법이다. 두 락에 대해서는 아래에서 더 알아본다.
- 낙관적 락은 작업을 수행하고 충돌이 발생했다면 후에 실행된 쿼리를 실패시키는 방법이다.
- 주로 version 컬럼을 하나 추가하고 작업을 수행 후, version 값을 검사하는 방식을 많이 사용한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/optimistic-lock.png" width="400">

- 위 그림을 살펴보자.
  - 사용자 A, B가 동시에 트랜잭션을 시작하고 이때 버전 값은 1이다.
  - A의 작업이 먼저 commit 되어 2로 수정되면 A의 경우 버전 값이 정상적으로 1 증가했기에 정상 종료된다.
  - B의 작업이 이후 실행되며 처음 읽은 version 값 1이 수정되었기 때문에 실패한다.
- 경쟁이 치열한 환경에서 낙관적 락은 vesrion 값 체크와 rollback 작업이 빈번해 질 수 있어 성능이 좋지 않다.
- **`따라서 동시성이 높다면 비관적 락을, 동시성이 낮다고 판단하면 낙관적 락을 고려해볼 수 있다`**.

## 공유 락 vs 배타 락

### 공유 락

- **`공유 락은 읽기를 위해 잡는 락으로`** 다른 트랜잭션에서 공유 락을 잡을 순 있지만 수정을 위한 배타 락을 잡을 순 없다. 이를 달리 표현하면 공유 락을 잡는 순간 해당 데이터는 락이 해제되기 전까지 수정이 불가능하다.
- 아래와 같이 두 개의 DB 커넥션에서 작업을 수행해보면 어떻게 동작하는지 이해할 수 있을 것이다.

  ```sql
    -- 트랜잭션 A
    begin;
    SELECT * FROM users WHERE id = 1 FOR SHARE;

    -- 트랜잭션 B
    SELECT * FROM users WHERE id = 1; -- 가능
    SELECT * FROM users WHERE id = 1 FOR SHARE; -- 가능

    UPDATE user SET name = 'aaa' WHERE id = 1; -- 대기
    SELECT * FROM users WHERE id = 1 FOR UPDATE; -- 대기
  ```

- 배타 락을 잡거나 데이터의 수정은 트랜잭션 A에서 commit이나 rollback을 하게 되면 수행된다.

### 배타 락

- **`배타 락은 수정을 위해 잡는 락으로`** 다른 트랜잭션에서 공유 락이나 배타 락을 잡을 순 없다. 달리 표현하면 배타 락을 잡는 순간 락이 해제되기 전까지 다른 트랜잭션은 대기해야 한다.

  ```sql
    -- 트랜잭션 A
    begin;
    SELECT * FROM users WHERE id = 1 FOR UPDATE;

    -- 트랜잭션 B
    SELECT * FROM users WHERE id = 1; -- 가능

    SELECT * FROM users WHERE id = 1 FOR SHARE; -- 대기
    UPDATE user SET name = 'aaa' WHERE id = 1; -- 대기
    SELECT * FROM users WHERE id = 1 FOR UPDATE; -- 대기
  ```

## AUTO INCREMENT 락

- 테이블 레벨의 락으로, 테이블에 AUTO_INCREMENT 컬럼이 있다면 데이터를 추가할 때 자동으로 잡게 되는 락이다.
- innodb_autoinc_lock_mode 변수로 아래와 같은 설정이 가능하다.
  - traditional(0)
  - consecutive(1)
  - interleaved(2)

### innodb_autoinc_lock_mode = traditional(0)

- 모든 INSERT 구문들(insert, insert...select, replace, load data)에 대해 auto-increment 락이 테이블 레벨로 동작
- 트랜잭션이 끝날때까지 적용되는 것이 아닌 해당 구문의 실행시까지만 유지되는 lock

### innodb_autoinc_lock_mode = consecutive(1)

- MySQL 버전 5.7까지 디폴트 값이다.
- bulk insert(insert...select, replace...select, load data)의 경우 테이블 수준에서 auto-increment 락을 잡는다.
- 간단한 insert 구문에 대해서는 테이블 레벨의 락을 잡지 않고 mutex를 활용하므로 동시성을 높일 수 있다.

### innodb_autoinc_lock_mode = interleaved(2)

- MySQL 버전 8부터 디폴트 값이다.
- 모든 insert 구문에서 테이블 수준에서 락을 잡지 않는다.
- 단순한 insert 구문에서는 증가 값에 gap이 존재하지 않지만 bulk insert의 경우 gap이 존재할 수 있다.
- 성능상 가장 빠르고 동시성이 좋지만, SQL 바이너리 로그 replay를 사용한 복구가 힘들다.

### AUTO_INCREMENT 컬럼이라고 항상 동일한 값으로 증가하지 않는다

- 테이블 수준에서 AUTO_INCREMENT 락을 잡는다고 해도 항상 동일한 값으로 증가하진 않는다.
- 트랜잭션에서 insert를 하고 auto increment 키를 할당받은 후 rollback 하는 경우가 그렇다.

  ```sql
    -- 트랜잭션 A
    begin;
    INSERT INTO users(name) values('하하');

    -- 트랜잭션 B
    begin;
    INSERT INTO users(name) values('유재석');

    -- 트랜잭션 A
    rollback;

    -- 트랜잭션 B
    commit;
  ```

- users 테이블에 id가 10까지 순차적으로 저장되었다고 가정하면 트랜잭션 A는 롤백이 발생했으므로 트랜잭션 B의 커밋은 id 12로 저장된다.

## InnoDB 락

- InnoDB에는 아래와 같이 3개의 락이 있는데 모두 인덱스를 잠근다는 특징을 가지고 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/innodb_lock_%EC%A2%85%EB%A5%98.png" width="500">

- 레코드 락은 인덱스의 레코드를 잠그는 락으로, 만약 인덱스가 없다면 InnoDB에서 내부적으로 만드는 clustered index를 이용해 잠근다.
- 갭락
  - 갭락은 다른 트랜잭션들이 gap에 insert하는 것을 막기 위한 용도로, 인덱스의 레코드와 레코드 사이, 첫번째 레코드의 앞과 마지막 레코드의 뒤를 잠글 수 있다.
  - 예를 들어 `SELECT name FROM users WHERE age BETWEEN 10 and 20 FOR UPDATE` 쿼리가 있다고 가정하자.
    - gap lock은 10과 20 사이에 15와 같은 데이터가 삽입되지 않도록 잠근다.
  - gap lock은 유니크 인덱스에서 특정 값을 찾는 쿼리에는 나타나지 않는다. 유니크 제약조건에 의해 중복으로 추가될 일이 없기 때문이다.
    - ex) `SELECT * FROM child WHERE id = 100`;
    - 쿼리에서 id가 유니크 인덱스가 아니라면 트랜잭션이 완료되기 전 다른 트랜잭션에서 id=100인 데이터를 추가할 수 있기에 gap lock을 건다.
  - gap lock은 격리 수준을 READ_COMMITED로 바꾸면 인덱스 검색과 스캔에서 비활성화 되며 오직 외래키와 중복키 체크에만 사용된다.
- 넥스트 키 락
  - 넥스트 키락은 레코드 락과 갭락을 합친 개념이다.
  - MySQL의 기본 격리 수준은 REPEATABLE_READ로 phantom read를 막기 위해 넥스트 키 락을 사용한다.
