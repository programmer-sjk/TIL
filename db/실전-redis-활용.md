# 실전! Redis 활용

- [인프런 강의](https://www.inflearn.com/course/%EC%8B%A4%EC%A0%84-redis-%ED%99%9C%EC%9A%A9/dashboard)

## Redis 소개

- Redis란 Remote Dictionary Server의 약자

### Redis 특징

- Redis는 C언어로 작성된 오픈소스 인메모리 데이터 저장소
- 메모리는 디스크에 비해 매우 빠르기 때문에 RDB와는 다른 구조와 특성을 갖는다.
- Redis는 싱글 스레드로 동작하여 단순한 디자인을 채택하였다.
- RDB(Redis DataBase) + AOF(Append Only File)을 통해 영속성을 제공할 수 있다.
- pub/sub 패턴을 지원해 손쉽게 어플리케이션 개발 가능

### Redis 장점

- 모든 데이터를 메모리에 저장하기 때문에 매우 빠른 읽기/쓰기 속도 보장
- Redis가 지원하는 다양한 Data Type을 활용해 다양한 기능 구현이 가능

### Redis 사용사례

- 캐싱 / Rate Limiter / 메시지 큐
- 실시간 분석 & 계산
  - 순위표 (랭킹)
  - 반경 탐색
  - 방문자 수 계산
- 실시간 채팅
  - Pub/Sub

### Redis 영속성

- Redis는 주로 캐시로 사용되지만 데이터 영속성을 위한 옵션 제공
- Redis DataBase
  - 특정 시간에 스냅샷을 생성하는 기술
  - 장애나 복제시 주로 사용됨
  - 새로운 스냅샷이 생성되기 전 데이터 유실될 수 있음
  - 스냅샷을 생성하는 동안 Redis 성능 저하가 발생
- Append Only File
  - Redis에 쓰기 작업을 모두 log에 저장
  - 데이터 유실 위험이 적지만 RDB 보다 느림

## Redis Data Type 알아보기

### Strings

- 문자열, 숫자 JSON 저장

```js
SET lecture inflearn  // lecture 키에 inflearn 저장
MSET price 100 age 30 // Multiple Set
MGET price age        // Multiple Get

INCR price       // 1 증가
INCRBY price 10  // 10 증가

SET inflearn-redis '{ price: 10000, language: ko }' // Json 저장 가능
SET inflearn-redis:ko:price 20000 // Redis에서는 콜론으로 의미를 구분
```

### Lists

- string을 Linked List로 저장
- Queue나 Stack을 쉽게 구현할 수 있음

```js
// queue로 동작
LPUSH queue job1 job2 job3
RPOP queue

// stack으로 동작
LPUSH stack job1 job2 job3
LPOP stack
```

### Sets

- 유니크한 문자열을 저장하는 정렬되지 않은 집합

```js
SADD user:1:fruits apple banana orange orange
SMEMBERS user:1:fruits
SCARD user:1:fruits // 카디널리티 확인
SISMEMBER user:1:fruits banana // 멤버 여부 확인

SINTER user:1:fruits user:2:fruits // user1,2가 공통으로 좋아하는 과일
SDIFF user:1:fruits user:2:fruits  // user1은 좋아하지만 user2는 좋아하지 않는 과일
SUNION user:1:fruits user:2:fruits // user1,2 합집합
```

### Hashes

- key-value 구조를 갖는 데이터 타입

```js
HSET lecture name inflearn price 100 language ko // 여러 key에 대한 value 추가
HGET lecture name
HMGET lecture name price language // Multiple key 조회
HINCRBY lecture price 100
```

### Sorted Set

- 유니크한 문자열을 score 기반으로 저장하는 정렬된 집합
- score가 동일하면 사전순으로 정렬됨

```js
ZADD points 10 teamA 10 teamB 100 teamC
ZRANGE points 0 -1 // 0부터 -1 범위는 전체 조회
ZRANGE points 0 -1 REV WITHSCORES // 역순으로 score와 함께 조회
ZRANK points teamA // teamA의 랭킹
```

### Geospatials

- 좌표를 저장하고 검색하는 데이터 타입
- 거리 계산, 범위 탐색 지원

```js
// 좌표 추가
GEOADD seoul:station 123.923123 37.556944 hong-dae
GEOADD seoul:station 127.027583 37.497927  gang-nam

GEODIST seoul:station hong-dae gang-nam KM // 거리 계산
```

## Redis 특수 명령어

### 데이터 만료(Expire)

- 기한이 만료된 데이터는 조회되지 않음
- 데이터가 만료되면 만료되었다는 표시만하고 백그라운드에서 주기적으로 삭제

```js
SET greeting hello
EXPIRE greeting 10 // 만료기한을 10s로 설정
TTL greeting // 만료까지 얼마나 남았는지
GET greeting
SETEX greeting 10 hello // 만료기한과 함께 저장
```

### SET NX/XX 옵션

- NX 옵션은 해당 Key가 존재하지 않을 경우에만 SET
- XX 옵션은 해당 Key가 존재하는 경우에만 SET
- SET이 동작하지 않은 경우 (nil) 반환

```java
SET greeting hello NX // key가 없었으니 성공
SET greeting bye NX // key가 있어서 (nil) 반환

SET greeting bye XX // key가 있어서 성공
SET greet hello XX // key가 없으니 (nil) 반환
```

### Transaction

- 다수의 명령을 하나의 트랜잭션으로 처리해서 원자성을 보장
- 중간에 에러가 발생하면 트랜잭션 내의 모든 작업이 Rollback

```js
MULTI // 트랜잭션 시작 명령어
INCR age
INCR age
INCR age
.
.
INCR age
DISCARD // 트랜잭션 롤백 명령어
EXEC // 트랜잭션 커밋 명령어
```
