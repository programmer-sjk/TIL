# 격리수준과 select 쿼리 성능

## 개요

- MySQL의 InnoDB 스토리지 엔진은 REPEATABLE 격리수준을 기본 값으로 사용한다.
- 최근 특정 기능에 대해 READ_COMMITTED 격리 수준으로 쓰라는 리뷰를 받았는데 이유는 아래와 같았다.
  - SELECT 쿼리에서 특정 값에 해당하는 엔티티를 조회만 하기 때문에 굳이 phantom read를 막을 필요가 없기 때문이라는 말이었다.
- 이런 리뷰를 받은 것도 좋은데, 이 내용에 대해 알아 볼 생각에 심장이 두근거린다. 여기서 이 내용에 대해 정리해보자.

## REPEATABLE 격리수준

- it will be added

## Phantom Read

- it will be added

## MySQL은 Phantom Read를 어떻게 막을까

- it will be added

## READ_COMMITED로 바꾼 후 성능은?

- it will be added

## 결론

- it will be added
