# MySQL Query Cache와 사용되지 않는 이유

## Query Cache

- SELECT 쿼리가 DB에 전달되면 DB는 쿼리와 쿼리 결과를 저장하고, 동일한 쿼리가 전달되면 cache에 접근해 결과를 빠르게 응답한다.
- 데이터가 자주 변경되지 않는 테이블이 있고, 동일한 쿼리를 많이 받는 환경에서 매우 유용하다.
- 테이블이 수정되면, 테이블과 관련된 캐쉬된 쿼리들은 제거된다.
- MySQL 5.7.20에서 Deprecated 되었고 8.0 버전부터 제거된 기능이다.

## 왜 사용하지 않을까?

- 동일한 쿼리와 결과에 대해 캐시한 결과를 제공하는 것은 얼핏 보면 꽤 유용해 보이는 방법이다. 이 기능이 왜 사용되지 않을까?

### Query Cache 옵션 확인 쿼리

- Query Cache 사용하는지 확인
  - `SHOW VARIABLES LIKE 'query_cache_type';`
  - 0 or OFF
    - Query Cache를 사용하지 않으며 결과를 캐시하지도 않는다.
  - 1 or ON
    - SELECT SQL_NO_CACHE 명령어를 제외하곤 캐시한다.
  - 2 or DEMAND
    - SELECT SQL_CACHE로 시작하는 명령어만 캐시한다
- 캐시 상태 변수 확인
  - `SHOW STATUS LIKE 'Qcache%';`

### Query Cache 비활성화

- 서버 실행 시점에 비활성화 시키려면 `query_cache_size` 시스템 변수를 0으로 설정한다
- 특정 쿼리를 비활성화 시키려면 `SQL_NO_CACHE` 키워드를 사용한다.
  - `SELECT SQL_NO_CACHE id, name FROM customer;`


## 레퍼런스

- https://dev.mysql.com/blog-archive/mysql-8-0-retiring-support-for-the-query-cache/
