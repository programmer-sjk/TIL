# NoSQL

- NoSQL 특징과 SQL과 차이에 대해 정리한다.

## NoSQL 특징

- NoSQL은 non sql과 Not Only SQL 이라는 두 가지 개념을 내포한다.
- NoSQL DB는 대표적으로 key-value, 문서형, graph DB가 존재한다.
- **`SQL과 달리 스케일 아웃이 간편하여 확장성이 매우 좋다`**.
- schemaless 특징으로 데이터 구조가 규격화 되어 있지 않으며 추가되는 데이터 구조를 별도의 오퍼레이션 없이 저장할 수 있다.
- 대량의 데이터 쓰기/읽기 작업에 적합하여 조인이나 관계, 빈번한 데이터의 수정에는 비효율적이다.

## 언제 SQL을 쓰는 DB가 적절한가

- ACID 속성으로 데이터의 무결성이 중요할 때
- 고정된 schema로 규격화된 데이터를 저장할 때
- 테이블간 관계를 맺어 조인하는 데이터가 많을 때

### ACID란?

- SQL에서는 데이터의 무결성을 보장하기 위해 트랜잭션을 사용한다.
- **`트랜잭션은 ACID의 성질을 가진다`**.
  - `Atomicity (원자성)`: 트랜잭션 작업이 전부 성공하거나 전부 실패하는 성질
  - `Consistency (일관성)`: 트랜잭션 작업 후 데이터 타입이나 제약이 바뀌지 않고 일관성 있게 유지하는 성질
  - `Isolation (격리성)`: 병렬의 트랜잭션 작업들이 서로 영향받지 않고 독립적으로 반영
  - `Durability (지속성)`: 트랜잭션 결과가 영원히 반영되는 성질, 즉 commit 되었다면 그 결과를 유지해야 함

## 언제 NoSQL을 쓰는 DB가 적절한가

- schemaless로 데이터 구조가 빈번히 바뀌는 데이터를 저장할 경우
- 대용량의 쓰기/조회 데이터인 경우
- 저장되는 데이터의 확장성이 중요한 경우

## 왜 NoSQL은 데이터 수정에 비효율적일까

- 사용자와 주문의 객체를 예로 들어보자
- SQL을 사용한다면 주문 테이블은 userId만 참고하고 있고, user의 데이터는 모두 user 테이블에만 존재하며 업데이트가 용이하다.
- **`NoSQL을 사용한다면 주문 document에 user의 어떤 데이터를 포함할 수 있다`**. (예를 들어 조회의 용의성을 위해 name 속성을 넣는다던지)
  - 이러면 주문 collection에 동일 사용자에 대해 중복된 데이터가 저장될 수 있다.
  - 이때 **`이 중복된 데이터에 대해 업데이트가 발생하는데`**, 여러 컬렉션에 대해 수행이 필요한 경우도 고려하면 업데이트가 빈번한 구조를 NoSQL과 맞지 않는다.