# Query Cache와 Buffer Pool Cache 차이

## Query Cache

- SELECT 쿼리가 DB에 전달되면 DB는 쿼리와 쿼리 결과를 저장하고, 동일한 쿼리가 전달되면 cache에 접근해 결과를 빠르게 응답한다.
- 데이터가 자주 변경되지 않는 테이블이 있고, 동일한 쿼리를 많이 받는 환경에서 매우 유용하다.

### 실시간으로 쓰기 작업(추가/수정/삭제)된 데이터는 어떻게 되는가?

- 테이블이 수정되면, 테이블과 관련된 캐쉬된 쿼리들은 제거된다.

### Query Cache 옵션 확인 쿼리

- Query Cache 사용하는지 확인
  - `SHOW VARIABLES LIKE 'have_query_cache';`
- 캐시 상태 변수 확인
  - `SHOW STATUS LIKE 'Qcache%';`

### Query Cache 비활성화

- 서버 실행 시점에 비활성화 시키려면 `query_cache_size` 시스템 변수를 0으로 설정한다
- 특정 쿼리를 비활성화 시키려면 `SQL_NO_CACHE` 키워드를 사용한다.
  - `SELECT SQL_NO_CACHE id, name FROM customer;`


