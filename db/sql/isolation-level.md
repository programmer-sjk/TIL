# DB 격리 수준

- 격리 수준이란 동시에 여러 트랜잭션에서 쿼리가 수행될때, 격리 수준에 따라 어떤 데이터를 보여줄지 결정하는 것이다.
- 격리 수준은 낮은 것부터 높은 순으로 다음 4가지가 있다.
  - READ_UNCOMMITED
  - READ_COMMITED
  - REPEATABLE_READ
  - SERIALIZABLE

## READ_UNCOMMITED

- 다른 트랜잭션에서 커밋하지 않은 데이터도 보여주는 격리 수준이다.

  ```sql
    -- 트랜잭션 A, B 공통적으로 READ UNCOMMITTED 사용
    SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;

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
