# 격리수준과 select 쿼리 성능

## 개요

- MySQL의 InnoDB 스토리지 엔진은 REPEATABLE 격리수준을 기본 값으로 사용한다.
- 최근 특정 기능에 대해 READ_COMMITTED 격리 수준으로 쓰라는 리뷰를 받았는데 이유는 아래와 같았다.
  - SELECT 쿼리에서 특정 값에 해당하는 엔티티를 조회만 하기 때문에 굳이 phantom read를 막을 필요가 없기 때문이라는 말이었다.
- 이런 리뷰를 받은 것도 좋은데, 이 내용에 대해 알아 볼 생각에 심장이 두근거린다. 여기서 이 내용에 대해 정리해보자.

## REPEATABLE 격리수준

- REPEATABLE 격리수준을 쉽게 이야기하면 한 트랜잭션 내에서 select 결과가 같음을 보장한다.
- READ_COMMITTED 격리수준은 A 트랜잭션이 끝나기 전에, B 트랜잭션이 값을 변경하고 COMMIT 하면 A 트랜잭션이 SELECT 할 때 변경된 값이 조회된다. 즉 한 트랜잭션 내에서 SELECT 결과가 다를 수 있다.
- 일반적으로 REPEATABLE 격리수준에는 SELECT 개수가 달라질 수 있는 Phantom Read 현상이 발생하는데 MySQL의 InnodB 스토리지 엔진은 next key lock(record lock + gap lock)을 이용해 Phantom Read 현상을 막는다.

## Phantom Read

- SELECT 개수가 달라질 수 있는 Phantom Read 현상이 언제 발생하는지 아래 그림을 통해 알아보자.

## MySQL은 Phantom Read를 어떻게 막을까

- it will be added

## READ_COMMITED로 바꾼 후 성능은?

- it will be added

## 결론

- it will be added
