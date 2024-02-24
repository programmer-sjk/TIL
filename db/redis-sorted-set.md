# Redis Sorted Set을 이용한 랭킹 정보 관리

- 최근 회사에서 영화 랭킹 정보를 Redis Sorted Set을 이용해 제공했다.
- 관련 된 내용들을 정리해본다.

## 랭킹 정보를 제공한다면 어떤 방법들이 있을까?

- 만약 어떤 종류의 데이터를 랭킹 순으로 보여줘야 한다면, 기술적으로 어떻게 제공할 수 있을까?
- 가장 먼저 생각나는 방법은 RDB에 쿼리로 조회해서 제공하는 방법이 있다.
  - Seed 데이터를 합산해서 랭킹 순으로 정렬해서 보여준다.
  - 배치에서 랭킹 데이터를 만들고, 단순 조회해서 그대로 보여준다.
- RDB가 아닌 Redis를 이용하여 랭킹 정보를 쉽게 제공할 수 있는데, 이 때 활용 가능한 방법이 Sorted Set 이다.

## Redis Sorted Set 이란?

- Redis는 Sorted Set 이라는 자료 구조를 제공한다.
- Sorted Set은 Score를 기준으로 정렬된 유니크한 string을 관리하는 컬렉션이다.
  - 만약 Score 값이 동일하다면 사전순으로 정렬된다.
- 대부분의 Sorted Set 동작은 O(log N)의 시간 복잡도를 가진다.

## Redis-cli를 이용한 Sorted Set 다루기

## NestJS에서 Redis Sorted Set 다루기
