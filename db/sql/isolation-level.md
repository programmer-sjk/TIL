# DB 격리 수준

- 격리 수준이란 동시에 여러 트랜잭션에서 쿼리가 수행될때, 격리 수준에 따라 어떤 데이터를 보여줄지 결정하는 것이다.
- 격리 수준은 낮은 것부터 높은 순으로 다음 4가지가 있다.
  - READ_UNCOMMITED
  - READ_COMMITED
  - REPEATABLE_READ
  - SERIALIZABLE

## MySQL 터미널에서 격리 수준 확인 및 변경

- mysql -u root -p 명령어를 통해 mysql에 접속했다면 현재 세션에서 격리수준을 아래처럼 조회할 수 있다.

  ```sql
    SELECT @@transaction_ISOLATION;
  ```

- 현재 세션에서 격리 수준을 다음과 같이 수정할 수 있다.

  ```sql
    SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
    SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
    SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
    SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;
  ```

## READ_UNCOMMITED

- 다른 트랜잭션에서 커밋하지 않은 데이터도 보여주는 격리 수준이다.

  ```sql
    -- 트랜잭션 A
    begin;
    INSERT INTO users(name) values('하하');

    -- 트랜잭션 B
    begin;
    SELECT * FROM users;
  ```

- 위에서 트랜잭션 A가 새로운 데이터 추가 후 commit을 하지 않았음에도 트랜잭션 B에서 조회를 할 경우 추가된 데이터가 보여지게 된다.
- 트랜잭션이 완료되지 않았음에도 다른 트랜잭션에 보이는 현상을 Dirty Read라고 부른다.
- READ UNCOMMITED는 정합성에 문제가 있는 격리 수준으로 상용 DB에서는 보통 사용하지 않는다.

## READ_COMMITED

- 다른 트랜잭션에서 커밋된 데이터를 보여주는 격리 수준이다.

  ```sql
    -- 트랜잭션 A
    begin;
    INSERT INTO users(name) values('하하');

    -- 트랜잭션 B
    begin;
    SELECT * FROM users; -- 트랜잭션 A가 아직 커밋하지 않으므로 보이지 않음

    -- 트랜잭션 A
    commit;

    -- 트랜잭션 B
    SELECT * FROM users; -- 커밋되었으니 보여짐
  ```

- 만약 트랜잭션 A에서 데이터를 수정했다면 테이블의 데이터는 변경되고 원래 데이터는 언두 로그로 백업된다. 아직 커밋하지 않은 시점에서 다른 트랜잭션이 해당 데이터를 조회하게 되면 언두 로그의 데이터를 참조한다.
- READ_COMMITED 격리 수준에선 Non Repeatable Read(반복 읽기 불가능) 문제가 발생할 수 있다. 아래 예시를 보자.

  ```sql
    -- 트랜잭션 A
    begin;
    SELECT * FROM users where name = '광수'; -- 조회된 데이터 X

    -- 트랜잭션 B
    begin;
    UPDATE users set name = '광수' where name = '광ㅅ'; -- 이름이 잘못 기입되서 수정함

    -- 트랜잭션 A
    SELECT * FROM users where name = '광수'; -- 조회된 데이터 O
  ```

- 이처럼 한 트랜잭션에서 읽기를 반복할 경우 결과가 다를 수 있다는 부정합 문제를 Non Repeatable Read라고 한다.
- 일반적인 경우 문제가 되지 않을 수 있지만, 하나의 트랜잭션 내에서 동일한 데이터를 여러번 처리하는 로직이 민감한 금융과 관련된 문제라면 조심해야 한다.
- READ_COMMITED 격리 수준은 트랜잭션 내에서 실행되는 SELECT와 그냥 실행되는 SELECT가 차이가 없다는 특징이 있다.
