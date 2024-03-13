# MySQL 쿼리 튜닝

- MySQL **`쿼리를 개선할 수 있는 방법을`** 정리한다.
- 기본적으로 MySQL 쿼리와 인덱스 등을 알고 있다는 가정하에 중요한 부분만 정리한다.

## 복합 인덱스

- 복합 인덱스는 **`명시된 컬럼 순으로 정렬되어 있어`** 쿼리하는 컬럼 순서에 따라 성능의 차이가 발생
- 복합 인덱스를 설계할 때 **`카디널리티가 높은 컬럼을`** 선행 컬럼으로 둘 수록 좋다.
  - 반면 **`카디널리티가 낮아도`** 여러 쿼리에서 특정 컬럼을 단독으로 사용한다면 선행으로 고려할 수 있다.
  - 반면 **`조인할 때 자주 사용되는`** 컬럼이라면 복합 인덱스에서 선행으로 고려할 수 있다.

## 커버링 인덱스

- 쿼리를 처리할 때 인덱스만 이용해서 쿼리를 처리할 수 있는 경우 **`커버링 인덱스라고`** 부른다.
- 커버링 인덱스를 사용하면 실행 계획의 `Extra`에서 `Using index`가 표기되어야 한다.

## Order By

- `order by` 컬럼에 인덱스가 있으면 이미 정렬된 데이터를 사용하므로 **`추가적으로 정렬할 필요가 없어진다`**.
- 인덱스가 없다면 `file sort` 방식으로 임시 테이블을 생성해서 정렬을 수행한다.
  - 실행 계획의 Extra 컬럼에 `Using filesort`가 있다면 인덱스를 타지 않는 것을 의미
  - **`인덱스가 없는데 LIMIT이 사용된 경우`** 전체 데이터를 정렬한 후에 LIMIT을 적용하므로 **효율이 안 좋다**.
- `order by` 실행 시 정렬 결과 순서가 동일해야 한다면 order by 절에 고유한 값을 가진 컬럼을 포함해야 한다.

  ```sql
    // category 값이 같다면 같은 category를 가진 row에 한해 정렬 순서가 실행마다 달라질 수 있음
    SELECT * FROM ratings ORDER BY category;

    // 이럴 때는 id까지 같이 정렬해야 순서가 보장됨
    SELECT * FROM ratings ORDER BY category, id;
  ```

## offset

- **`LIMIT 절에 offset을 사용하면`** 데이터를 가져온 후 offset 이전의 데이터는 버림
- 마지막으로 처리 된 레코드의 키를 사용해서 offset을 사용하지 않는다.

  ```sql
    // 처음 조회
    SELECT * FROM my_table WHERE status = 'pending' ORDER BY id LIMIT 10000;

    // 다음 조회
    SELECT * FROM my_table WHERE id > 15800 AND status = 'pending' ORDER BY id LIMIT 10000;
  ```
