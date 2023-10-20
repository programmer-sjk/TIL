# NestJS와 TypeORM에서 사용되는 다양한 캐시

- NestJS와 TypeORM에서 어떤 캐시들을 어떤 용도로 사용하는지 정리한다.

## 글로벌 캐시

- 여러 서버에서 공용으로 사용할 수 있는 캐시를 의미한다. 여러 서버간 공유가 쉽고 네트워크 통신이 발생하므로 로컬 캐시보다는 상대적으로 느리다.

### NestJS

### TypeORM

- typeorm은 getMany, getOne, getRawMany, find*, count* 메소드 등에 cache를 적용할 수 있다.
- 기본적으로 ormconfig.json 파일에 따로 설정이 없다면 query-result-cache 테이블을 사용하여 쿼리와 그 결과를 저장한다. 반대로 ormconfig.json에 cache 설정을 추가하면 redis나 ioredis 같은 타입의 캐시도 사용할 수 있다.
- ormconfig.json 예시

  ```json
  {
    type: "mysql",
    host: "localhost",
    username: "test",
    ...
    cache: {
      type: "redis",
      options: {
        host: "localhost",
        port: 6379
      }
    }
  }
  ```

- Typeorm 예제 코드

```ts
async getAdminUsers() {
  const result = await this.createQueryBuilder('user')
    .where("user.isAdmin = :isAdmin", { isAdmin: true })
    .cache('cache:admin-users', Milliseconds.ONE_HOUR) // ormconfig 설정에 따라 cache에 저장
    .getMany();
}
```

## 로컬 캐시

- 서버마다 가지고 있는 로컬 캐시를 의미한다. 서버마다 캐시 상태가 다를 수 있으므로 주의해야 한다.
  - 서버 A에 캐시를 적용해도 서버 B에는 아직 캐시가 적용되지 않았을 수 있고
  - 서버 A에서 무효화(invalidate)를 해도 서버 B에서는 무효화되지 않을 수 있다.
- NestJS 예제 코드

```ts
// module sample code
import { CACHE_MANAGER, CacheModule } from "@nestjs/cache-manager";
import { Module } from "@nestjs/common";
import { InMemoryCacheService } from "./in-memory-cache.service";

@Module({
  // 아래와 같이 CacheModule.register에 별다른 옵션 없이 주면 로컬 캐시로 동작
  imports: [CacheModule.register({ ttl: 60, keyPrefix: "cache:" })],
  providers: [
    InMemoryCacheService,
    {
      provide: "InMemoryCacheToken",
      useExisting: CACHE_MANAGER,
    },
  ],
  exports: [InMemoryCacheService],
})
export class InMemoryCacheModule {}

// service sample code
import { Inject, Injectable } from "@nestjs/common";
import { Cache } from "cache-manager";

@Injectable()
export class InMemoryCacheService {
  constructor(@Inject("InMemoryCacheToken") private cacheManager: Cache) {}

  set<T>(k: string, v: T, ttl = 60) {
    return this.cacheManager.set<T>(k, v, { ttl });
  }

  get<T>(k: string) {
    return this.cacheManager.get<T>(k);
  }
}
```
