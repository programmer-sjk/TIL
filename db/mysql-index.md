# MySQL Index

## 커버링 인덱스란

- 쿼리를 실행하는데 인덱스로 모든 처리가 가능한 인덱스를 **커버링 인덱스**라고 부른다.
  - `ex) EXPLAIN SELECT id FROM report WHERE id = 1;`
- 커버링 인덱스를 사용할 경우 실행 계획에서 다음과 같이 **`using index가`** 표기된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/covering-index.png" width="400">

### clustered index vs non-clustered index

- innoDB는 **`clustered index라고`** 불리는 특별한 인덱스가 있는데 PK가 있다면 PK가 clustered index가 된다.
  - **만약 PK를 지정하지 않았다면** Not NULL인 유니크 인덱스를 찾아 clustered index로 지정한다.
  - **PK와 유니크 인덱스 둘 다 없다면** 6 byte의 hidden 키를 생성해 clustered index로 사용한다.
- clustered index가 중요한 이유는 모든 `non-clustered index(secondary index)`가 clustered index를 통해 실제 데이터를 찾는다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/non-clustered-index.png" width="800">

- users 테이블에 PK인 id 컬럼과 phone에 단일 인덱스가 걸려있다고 가정하면 **모든 non-clustered 인덱스는 clustered index를 알고 있기 때문에** 아래 쿼리는 커버링 인덱스를 타게 된다.
  - `ex) EXPLAIN SELECT id, phone FROM report WHERE id = 1;`

### Multiple-Column Index에도 적용될까?

- 아래와 같이 테이블을 생성하자. PK는 id 컬럼이고 `email, name` 으로 다중 컬럼 인덱스를 생성했다.

  ```sql
    CREATE TABLE `multi_index_test` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `email` varchar(255) NOT NULL,
      `password` varchar(255) NOT NULL,
      `name` varchar(255) NOT NULL,
      PRIMARY KEY (`id`),
      KEY `email_name` (`email`,`name`) # email + name 으로 다중 컬럼 인덱스 생성
    )

    # 데이터 추가
    insert into multi_index_test(`email`, `password`, `name`) values('a@test.com', 'password1', 'seo1');
    insert into multi_index_test(`email`, `password`, `name`) values('b@test.com', 'password2', 'seo2');
    insert into multi_index_test(`email`, `password`, `name`) values('c@test.com', 'password3', 'seo3');
  ```

- 아래 쿼리의 실행 계획을 확인하면 마찬가지로 `Using Index`가 표기되어 커버링 인덱스를 확인할 수 있다.
  - `EXPLAIN SELECT id, email, name FROM multi_index_test WHERE email = 'a@test.com' and name = 'seo1';`
- 또한 아래 쿼리는 다중 컬럼 인덱스에서 순서에 맞게 쿼리하지 않았기 때문(email이 인덱스 순서 상 앞에 배치되기 때문)에 커버링 인덱스가 표시되지 않는 것을 확인할 수 있다.
  - `EXPLAIN SELECT id, name FROM multi_index_test WHERE name = 'seo1';`
