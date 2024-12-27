# DB 격리 수준

- 격리 수준이란 동시에 여러 트랜잭션에서 쿼리가 수행될때, 격리 수준에 따라 어떤 데이터를 보여줄지 결정하는 것이다.
- 격리 수준은 낮은 것부터 높은 순으로 다음 4가지가 있다.
  - READ_UNCOMMITED
  - READ_COMMITED
  - REPEATABLE_READ
  - SERIALIZABLE
