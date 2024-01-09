# 실전! Redis 활용

- [인프런 강의](https://www.inflearn.com/course/%EC%8B%A4%EC%A0%84-redis-%ED%99%9C%EC%9A%A9/dashboard)

## Redis 소개

- Redis란 **Remote Dictionary Server**의 약자

### Redis 특징

- Redis는 C언어로 작성된 오픈소스 **`인 메모리 데이터 저장소`**
- 메모리는 **디스크에 비해 매우 빠르기 때문에** RDB와는 다른 구조와 특성을 갖는다.
- Redis는 싱글 스레드로 동작하여 단순한 디자인을 채택하였다.
- **`RDB(Redis DataBase) + AOF(Append Only File)을`** 통해 영속성을 제공할 수 있다.
- **pub/sub** 패턴을 지원해 손쉽게 어플리케이션 개발 가능

### Redis 장점

- 모든 데이터를 메모리에 저장하기 때문에 **`매우 빠른 읽기/쓰기 속도 보장`**
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
- **`Redis DataBase`**
  - 특정 시간에 스냅샷을 생성하는 기술
  - 장애나 복제시 주로 사용됨
  - 새로운 스냅샷이 생성되기 전 데이터 유실될 수 있음
  - 스냅샷을 생성하는 동안 Redis 성능 저하가 발생
- **`Append Only File`**
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

- **NX** 옵션은 해당 Key가 존재하지 않을 경우에만 SET
- **XX** 옵션은 해당 Key가 존재하는 경우에만 SET
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

## 데이터 타입 활용

### OTP (One-Time Password)

- 인증을 위해 사용하는 6자리 코드가 있다고 가정
- 휴대폰 번호에 맞게 cache key를 생성하고 3분의 유효시간을 설정
  - **`otp:010-1111-2222 928512 EX 180`**
- 사용자는 전달받은 OTP를 서버에게 전달
- 서버는 캐시 키에 해당하는 값과 OTP를 비교해서 일치하면 인증 다음의 로직을 실행

### 분산 락 (Distributed Lock)

- 분산 환경에서 동일한 자원에 접근 시, **동시성 문제 해결**
- 잔고가 100원인 통장에 2개의 요청이 동시 접근
  - 요청 1에서는 20을 차감, 요청 2에서는 20을 더함
  - Lock이 없는 경우 나중에 수행된 요청에 의해 결과가 덮어 씌워짐
- 공유 자원을 표현할 수 있는 키로 NX 옵션으로 OK 명령을 받은 경우에만 Lock을 획득하도록 함

### Rate Limiter

- 특정 IP/유저/어플리케이션에 따라 요청의 수를 제한하는 기술

### Fixed Window

- **`시간마다 허용량을 초기화`**
  - 1분마다 20개의 요청을 허용
  - 최악의 경우 59초에 20개의 요청이 들어오고 0초에 요청개수가 초기화되면 20개의 요청이 들어옴
- Redis 키에 **`ip:분(ex: 192.168.10.15:10)`** 을 키로 값을 조회
- 아직 요청수가 기준보다 낮다면 INCR 명령어로 접근 수를 1 증가시킨다.

### Sliding Window Rate Limiter

- 시간에 따라 Window를 이동시켜 **`동적으로 요청 수를 조절하는 기술`**
- Sorted Set에 Score 대신 시간을 초 단위로 변경한 유닉스 타임을 저장
  - 10번 요청이 들어오면 Sorted Set에 서로 다른 유닉스 타임으로 10개가 저장
- 요청이 들어오면 **`ZREMRANGEBYSCORE`** 명령어로 0부터 현재 유닉스 타임에서 60초를 뺀 값으로 1분이 지난 요청 삭제
- **`ZCARD`** 명령어로 카디널리티를 계산하면 현재 시간 기준으로 60초 전까지 요청 횟수를 계산할 수 있음

### 장바구니

- 사용자가 구매 의사가 있는 상품을 임시로 모아두는 공간으로 수시로 변경이 발생할 수 있고 실제 구매로 이어지지 않을 수 있다.
- 장바구니에 상품을 추가하면 **`SADD`** card item1과 같은 명령어 저장
- 장바구니를 조회하면 Redis에 **`SMEMBERS`** 명령어로 item을 조회해서 응답

### 온라인 상태

- 사용자의 현재 온라인 상태를 표현
- 실시간성을 완벽히 표현하지 않으며 수시로 변경되는 값
- 강의에서는 Redis의 **BitMap**을 활용
- 사용자 수가 많지 않다면 **`online-status:user:분`** 같은 키로도 표현 가능

## Redis 사용 시 주의사항

### O(n) 명령어

- Redis의 **`대부분 명령어는 O(1)의 시간 복잡도를 가져서 매우 빠름`**
- **`일부 명령어의 경우 O(n)의 시간 복잡도를`** 가지기 때문에 주의
- Redis는 싱글 스레드로 명령을 순차적으로 수행하기 때문에 오래 걸리는 O(n) 명령어 수행 시, 성능 저하
- 대표적인 명령어
  - **KEYS** (지정된 패턴과 일치하는 모든 Key 조회)
  - **SMEMBERS** (Set의 모든 멤버 반환, 데이터가 많을수록 오래 걸림)
  - **HGETALL** (Hash의 모든 필드 반환)
  - **SORT** (List, Set, ZSet의 item을 정렬해서 반환)

### Thundering Herd Problem

- **병렬 요청이 공유 자원에 접근할 때 급격한 과부하를 의미함**
  - 주로 대규모 트래픽을 다루는 서비스에서 발생
- 문제가 발생하는 예시
  - 1s 걸리는 통계 작업이 있어 이걸 Cache 해둠
  - 하루가 지나면서 (0시 0분 0초) 통계 캐시가 만료됨
  - 다수의 요청은 캐시가 만료되어 동시에 DB에 쿼리를 발생시킴
  - 1s 걸리는 통계 작업이 10,000번 발생하여 순식간에 어플리케이션 서버에 부하를 주게 된다.
- 이처럼 **`Thundering Herd Problem는`** 캐시 만료에 의해 발생할 수 있음
- 이를 방지하기 위해 배치에서 통계 캐시를 주기적으로 최신화 시켜주는 방식이 있음

### Cache Invalidation

- 원본 DB에는 잔액이 8만원 남았는데 Cache에는 여전히 100만원이 남아있는 경우 클라이언트에게 잘못된 데이터를 응답할 수 있음
