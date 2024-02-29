# Redis Sorted Set을 이용한 랭킹 정보 관리

- 최근 회사에서 영화 랭킹 정보를 **`Redis Sorted Set을`** 이용해 제공했다.
- Sorted Set이 무엇이고 왜 쓰는지 내용들을 정리해본다.

## 랭킹 정보를 제공한다면 어떤 방법들이 있을까?

- 만약 어떤 종류의 데이터를 **랭킹 순으로** 보여줘야 한다면, 기술적으로 어떻게 제공할 수 있을까?
- 가장 먼저 생각나는 방법은 **`RDB에 쿼리로 조회해서 제공하는 방법이**` 있다.
  - Seed 데이터를 합산해서 랭킹 순으로 정렬해서 보여준다.
  - 배치에서 랭킹 데이터를 만들어 저장하고, 랭킹 정보를 단순 조회해서 그대로 보여준다.
- RDB가 아닌 **`Redis를 이용하여 랭킹 정보를 쉽게 제공할 수 있는데,`** 이 때 활용 가능한 방법이 Sorted Set 이다.

## 왜 Redis를 쓸까?

- **`속도 때문이다.`**
- RDB를 사용하면 캐시되지 않았다고 가정할 경우 데이터를 디스크에서 불러오지만 Redis는 항상 메모리에서 가져온다.
- 실무에서 사용한다면 Redis의 랭킹 정보가 휘발될 수 있기 때문에 랭킹 정보 or 랭킹 정보를 생성하는 데이터가 DB에 있어야 하거나 Redis를 백업해야 한다.

## Redis Sorted Set 이란?

- Redis는 **`Sorted Set`** 이라는 자료 구조를 제공한다.
- Sorted Set은 **`Score를 기준으로 정렬된 유니크한 string을 관리하는 컬렉션이다`**.
  - 만약 Score 값이 동일하다면 사전순으로 정렬된다.
- 대부분의 Sorted Set 동작은 **`O(log N)`** 의 시간 복잡도를 가진다.

## Redis-cli를 이용한 Sorted Set 다루기

- Redis가 이미 설치되어 있다고 가정하고 local에서 `redis-cli` 명령어로 실습해본다.

### 추가

- Sorted Set의 추가는 `ZADD` 명령어로 추가한다.
  - 명령어: `ZADD SORTED_SET_NAME SCORE MEMBER`

```
ZADD mySortedSet 10 "1등"
ZADD mySortedSet 9 "2등"
ZADD mySortedSet 8 "3등"
ZADD mySortedSet 7 "4등"
ZADD mySortedSet 6 "5등"
```

### 수정

- 동일한 멤버에 대해 `ZADD` 명령어를 수행하면 SCORE 값이 업데이트 된다.

```
127.0.0.1:6379> ZREVRANGE mySortedSet 4 5 WITHSCORES
5등
6

127.0.0.1:6379> ZADD mySortedSet 3 "5등"
0

127.0.0.1:6379> ZREVRANGE mySortedSet 4 5 WITHSCORES
5등
3
```

### 조회

- `ZRANGE`로 Score가 낮은 순부터 조회할 수 있다.

```
127.0.0.1:6379> ZRANGE mySortedSet 0 5
5등
4등
3등
2등
1등
```

- `ZREVRANGE`를 이용해 점수가 높은 순부터 조회할 수 있다.
  - Redis 6.2 버전부터 해당 메소드는 [deprecated](https://redis.io/commands/zrevrange/) 된다.
  - Redis 6.2 버전부터 REV 인자가 기능을 대체하게 된다.

```
127.0.0.1:6379> ZREVRANGE mySortedSet 0 5
1등
2등
3등
4등
5등
```

### 삭제

- Sorted Set 내부에서 멤버 삭제는 `ZREM` 명령어를 이용한다.

```
127.0.0.1:6379> ZREM mySortedSet "5등"
1

127.0.0.1:6379> ZREVRANGE mySortedSet 0 5
1등
2등
3등
4등
```

- Sorted Set 자체 삭제는 del 명령어를 이용한다.

```
127.0.0.1:6379> DEL mySortedSet
1

127.0.0.1:6379> ZREVRANGE mySortedSet 0 5
```

## NestJS에서 Redis Sorted Set 다루기

- NestJS 프레임워크에서 Redis의 Sorted Set을 사용해보자.
- 예시에서 사용한 전체 코드는 [여기서](https://github.com/programmer-sjk/nestjs-redis) 확인할 수 있다.
- nest-cli를 활용해 NestJS 프로젝트를 생성하고 Redis 동작에 필요한 모듈을 설치한다.
  - `yarn add @nestjs-modules/ioredis ioredis`
- Redis와 연동하는 모듈을 생성한다.

  ```ts
  import { Module } from '@nestjs/common';
  import { RedisModule as IORedisModule } from '@nestjs-modules/ioredis';
  import { RedisService } from './redis.service';

  @Module({
    imports: [
      IORedisModule.forRoot({
        type: 'single',
        url: 'localhost',
      }),
    ],
    providers: [RedisService],
    exports: [RedisService],
  })
  export class RedisModule {}
  ```

- 다른 모듈에서 호출 할 RedisService를 아래와 같이 작성한다.

  - 랭킹을 추가할 때 호출할 `zadd` 명령어와 조회할 때 호출할 `zrange, zrevrange`를 선언한다.

  ```ts
  import { InjectRedis } from '@nestjs-modules/ioredis';
  import { Injectable } from '@nestjs/common';
  import { Redis } from 'ioredis';

  @Injectable()
  export class RedisService {
    constructor(@InjectRedis() private redis: Redis) {}

    async zadd(key: string, score: number, member: string) {
      return this.redis.zadd(key, score, member);
    }

    async zrange(key: string, max: number) {
      return this.redis.zrange(key, 0, max);
    }

    async zrevrange(key: string, max: number) {
      return this.redis.zrevrange(key, 0, max);
    }
  }
  ```

- person 서비스에서 redis 서비스를 호출한다.

  - `zrange` 명령어는 기본적으로 **`Score가 낮은 순으로 응답한다`**.
  - 만약 Sorted Set에 저장되는 Score가 높을 수록 랭킹이 높아진다면 `zrevrange`를 호출해야 한다.
  - 아래 예시에서는 Score가 곧 랭킹이므로 `zrange`를 호출했다.

  ```ts
  import { Injectable } from '@nestjs/common';
  import { RedisService } from '../redis/redis.service';

  @Injectable()
  export class PersonService {
    constructor(private readonly redisService: RedisService) {}

    async getPersonRanking() {
      return this.redisService.zrange('mySortedSet', 10);
    }

    async addPersonRanking() {
      const persons = [
        { ranking: 1, name: '마동석' },
        { ranking: 2, name: '손석구' },
        { ranking: 3, name: '아이유' },
        { ranking: 4, name: '유재석' },
        { ranking: 5, name: '이광수' },
      ];

      for (const person of persons) {
        await this.redisService.zadd(
          'mySortedSet',
          person.ranking,
          person.name
        );
      }
      return true;
    }
  }
  ```

- Person Controller를 생성한 뒤, postman을 이용해 호출하면 아래와 같은 결과를 확인할 수 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/redis-sorted-set-result.png" width="600">
