# 격리수준과 SELECT 쿼리 성능

## 개요

- MySQL의 InnoDB 스토리지 엔진은 `REPEATABLE` 격리수준을 기본 값으로 사용한다.
- 최근 특정 기능에 대해 `READ_COMMITTED` 격리 수준으로 **변경하라는 리뷰**를 받았는데 이유는 아래와 같았다.
  - SELECT 쿼리에서 특정 값에 해당하는 엔티티를 조회만 하기 때문에 굳이 `phantom read`를 막을 필요가 없기 때문이라는 말이었다.
- 이런 리뷰를 받은 것도 좋은데, 이 내용에 대해 알아 볼 생각에 심장이 두근거린다. 여기서 이 내용에 대해 정리해보자.

## REPEATABLE 격리수준

- `REPEATABLE` 격리수준을 쉽게 이야기하면 **한 트랜잭션 내에서 select 결과가 같음**을 보장한다.
- `READ_COMMITTED` 격리수준은 A 트랜잭션이 끝나기 전에, B 트랜잭션이 값을 변경하고 COMMIT 하면 A 트랜잭션이 SELECT 할 때 변경된 값이 조회된다. 즉 한 트랜잭션 내에서 SELECT 결과가 다를 수 있다.
- 일반적으로 `REPEATABLE` 격리수준에는 **SELECT 조회 결과 수**가 달라질 수 있는 Phantom Read 현상이 발생하는데 MySQL의 InnodB 스토리지 엔진은 next key lock(record lock + gap lock)을 이용해 Phantom Read 현상을 막는다.

## Phantom Read

- SELECT 개수가 달라질 수 있는 `Phantom Read` 현상이 언제 발생하는지 아래 그림을 통해 알아보자.
  - 트랜잭션 1에서 남자를 조회해서 2개의 row를 얻었다.
  - 트랜잭션 1에서 잠시 다른일을 하는동안 트랜잭션 2에서 새로운 남자를 추가한다.
  - 트랜잭션 1에서 다시 남자를 조회하면 3개의 row를 조회하게 된다.
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/phantom_read.png" width="600">

## MySQL은 Phantom Read를 어떻게 막을까

- 우선 MySQL의 InnoDB 스토리지 엔진의 잠금을 이해해야 한다.
- InnoDB는 하나의 레코드를 잠그는 레코드락, 레코드 사이의 간격을 잠그는 갭락, 이 둘을 합쳐서 넥스트 키락이라는 잠금을 사용한다.
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/innodb_lock_%EC%A2%85%EB%A5%98.png" width="400">

- 테이블을 검색할 때 발견된(조회된) 레코드에 대해 레코드와 레코드 앞의 갭에 대해 락을 건다. 따라서 조회된 레코드를 수정할 수 없고, 레코드 앞의 갭에 락을 걸었기 때문에 레코드 앞에 데이터가 새로 추가될 수 없다.
- 또한 검색된 마지막 레코드 뒤에 갭락을 걸어서, 마지막 레코드의 다음 데이터를 추가할 수 없도록 락을 건다.

```sql
SELECT * FROM child WHERE id > 100
```

- A 트랜잭션에서 위 쿼리를 통해 102, 105가 나왔다고 가정하자. 이 과정에서 이미 102, 105 레코드와 주변 갭은 락이 걸린 상태이다.
- A 트랜잭션은 끝나지 않았고 이 떄 B 트랜잭션이 101을 추가한다면, 102 앞의 갭은 락이 걸려 추가할 수 없어 대기하게 된다. 마지막 레코드인 105 이후의 106도 마찬가지의 개념으로 대기하게 된다.

## 다시 격리수준과 SELECT 성능 예상

- 다시 본 주제로 돌아오자. 주제에 해당되는 격리수준과 `Phantom Read`에 대해서는 위에서 설명했다.
- 그렇다면 특정 값을 조회하는 쿼리에서 격리수준을 `REPEATABLE -> READ_COMMITED`로 바꾸면 어떤 성능의 이득을 예상할 수 있을까?
- `REPEATABLE은` Phantom Read를 막기 위해 Next Key 잠금을 수행하고, `READ_COMMITED는` Phantom Read가 발생하기 때문에 Next Key 잠금을 사용하지 않을 것을 짐작할 수 있다.
- 옵티마이저가 여기에 관여하지 않는다고 가정하면 두 격리수준에서 실제 차이는 Next Key 잠금을 사용하느냐 사용하지 않느냐다.
- 이 상태에서 쿼리 하나의 성능은 큰 차이가 없을 것이라 생각이 들었다. 그렇다면 10만건 이상 반복해서 조회한다면 성능에 영향이 날지 아래에서 정리하겠다.

## READ_COMMITED로 바꾼 후 성능은?

- 테스트를 위해 30개의 다양한 컬럼(integer, varchar, enum, timestamp)을 가진 50만건의 랜덤 데이터를 만들었다.
- 격리수준을 각각 READ_COMMITTED, REPEATABLE로 설정한 상태에서 1만건 ~ 20만건 까지 랜덤 id를 조회해서 SELECT시 성능을 비교한다.
- 동일한 id로 조회하게 되면 캐시나 버퍼 풀로 성능의 이득이 있을 수 있어 id는 랜덤하게 조회한다.
- 테스트 한 코드는 아래와 같은 형태이다.

```js
  // controller
  @Get()
  async test(): Promise<void> {
    // 아래 getMovieIds는 아래 쿼리의 결과를 반환한다.
    // SELECT id FROM movie ORDER BY RAND() LIMIT 10000 ~ 250000
    const movieIds = await this.isolationService.getMovieIds();

    console.time('test');
    for (const id of movieIds) {
      const movie = await this.isolationService.test(id);
    }
    console.timeEnd('test');
  }

  // service
  // 격리수준은 아래 Transactional을 유지하느냐 주석처리 하느냐로 성능을 비교한다.
  @Transactional({ isolationLevel: IsolationLevel.READ_COMMITTED })
  async test(id: number) {
    return this.isolationRepository.findOne(id);
  }
```

- 테스트 결과는 아래와 같았다.
  - Limit이 1만건 일 경우
    - REPEATABLE(default): 14.6s
    - READ_COMMITED: 7.4s
  - Limit이 5만건 일 경우
    - REPEATABLE(default): 1분 21초
    - READ_COMMITED: 35.4s
  - Limit이 10만건 일 경우
    - REPEATABLE(default): 2분 27초
    - READ_COMMITED: 2분 17초
  - Limit이 20만건 일 경우
    - REPEATABLE(default): 5분 33초
    - READ_COMMITED: 1분 56초
  - Limit이 25만건 일 경우
    - REPEATABLE(default): 56초
    - READ_COMMITED: 1분 44초

## 결론

- 위 테스트 결과를 완전히 신뢰할 수 없는게 조회되는 데이터 개수가 많아질수록 캐시에 영향을 받는다.
  - 25만건까지 수행 후, 다시 5만건을 조회하면 두 격리수준 모두 10s 대에서 조회가 완료된다.
- 하지만 애초에 확인하고 싶었던 `READ_COMMITTED로` 격리수준을 변경했을 때 `Phantom Read` 처리를 하지 않아
  **성능상 이점을 얻을 수 있다는 점은 확인**할 수 있었다.
- 조회만 하는 Service의 함수가 자주 호출되는게 예상된다면 격리수준을 `READ_COMMITTED` 으로 바꾸는게 좋을 것 같다.
