# IN vs Exists 쿼리

## IN

- IN 절에는 값이나 서브 쿼리가 올 수 있다.
- IN 절의 실행 순서는 서브 쿼리가 먼저 실행되고 main 쿼리가 실행된다.
  - 아래는 100 미만의 id를 가진 유저 중 리뷰를 작성한 유저를 찾는 쿼리로 서브 쿼리가 먼저 실행된다.
  - `select * from user where id in (select  userId from review where userId < 100);`
- IN 절에 해당하는 **`데이터가 적다면 성능이 빠른편이다`**. 반대로 **`몇백 ~ 몇천의 데이터가 IN 절에 들어가면 성능이 느려진다`**.

## Exists

- Exists에는 서브 쿼리만 올 수 있으며 외래키가 설정된 테이블간에 데이터를 찾는데 유리하다.
- 쿼리의 실행 순서는 main 쿼리가 먼저 실행된 후 서브 쿼리가 실행된다.
  - `select * from user u where exists (select 1 from review r where userId < 100 AND r.userId = u.id);`
- Exists는 IN 절과 달리 **`select 구문을 통해 데이터를 로드하지 않고 true/false만 반환하기에 일반적으로 IN 절에 비해 성능이 좋다`**.
- 다만 데이터가 적을 때는 IN절과 큰 차이가 없으며 오히려 IN이 나을때도 있다. 하지만 서브쿼리에 매칭되는 데이터가 많아질수록 exists를 고려해야 한다.
