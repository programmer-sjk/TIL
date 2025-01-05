# MySQL Index

## 인덱스란 무엇인가?

- DB에서 인덱스란 지정한 컬럼을 메모리에 유지해 디스크의 데이터를 빠르게 접근하기 위한 기법이다.
- 인덱스를 생성하면 데이터를 추가, 수정, 삭제시 인덱스에도 반영해야 하기 때문에 **`쓰기 성능이 조금 느려지는 대신 조회 성능은 빠르게 높일 수 있다`**.

## B-Tree

- MySQL은 인덱스로 **`B-Tree(Balanced Tree)를`** 사용한다.
- B-Tree는 최상위에 루트 노드, 중간 노드들을 브랜치 노드, 최하위 노드를 리프 노드라고 부른다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/btree.png" width="800">

- B-Tree는 N개의 노드를 가질 수 있으며 정렬되어 있다. 좌우 자식간에 균형이 맞지 않으면 비 효율적인데 **`B-Tree는 노드가 추가/삭제 될 때 자동으로 균형을 잡아준다`**.

## 커버링 인덱스란

- 쿼리를 실행하는데 인덱스로 모든 처리가 가능한 인덱스를 **커버링 인덱스**라고 부른다.
  - `ex) EXPLAIN SELECT id FROM report WHERE id = 1;`
- 커버링 인덱스를 사용할 경우 실행 계획에서 다음과 같이 **`using index가`** 표기된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/covering-index.png" width="400">

## clustered index (클러스터링 인덱스)

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

## 복합 인덱스 카디널리티 기준

- 복합 인덱스를 생성할 때 **`카디널리티가 낮은 것 -> 높은 것 vs 높은 것 -> 낮은 것`** 중 어떤 경우가 더 빠를까?
- 일반적으로 where 절에서 조건이 **`equal(=)일`** 경우에는 카디널리티가 다른 컬럼 순서에 대해 성능의 차이가 없다.
  - `ex) SELECT * FROM users where name = 'seo' AND age = 33`;
- 만약 범위라면 달라지는데 상황마다 달라질 순 있지만 경험적으로 **`카디널리티가 높은 것 -> 낮은 순으로`** 만드는게 효율이 좋다.
