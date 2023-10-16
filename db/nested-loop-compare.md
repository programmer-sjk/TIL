# Nested Loop Join과 Block Nested Loop Join

- Nested Loop Join과 Block Nested Loop Join 차이를 알아보자.
- MySQL 8.0.20 버전 부터 Block Nested Loop Join 대신 사용되는 Hash join을 알아본다.

## Nested Loop Join

- MySQL 서버에서 사용되는 대부분의 조인은 네스티드 루프 조인의 형태로, 일반적으론 조인의 연결 조건이 되는 컬럼에 모두 인덱스가 사용된다.

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/nested-loop.png" width="600">

- 위 방식을 sudo 코드로 나타내면 아래와 같다.

```js
for (row1 in tableA) {
  for (row2 in tableB) {
    if (conditionMatched) return row1, row2;
  }
}
```

## Block Nested Loop Join

- 위 Nest Loop Join에서 조인되는 컬럼에 인덱스가 걸려있다면 성능상에 문제가 없다. 하지만 결국 인덱스 없이 실행된다면 비용이 아주 비싸게 되며 이때 사용되는게 Block Nested Loop Join이다.

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/block-nested-loop.png" width="600">

- 위 방식을 sudo 코드로 나타내면 아래와 같다.

```js
for (row1 in tableA) {
  addJoinBuffer(row1);
  if (isBufferFull()) {
    for (row2 in tableB) {
      for (bufferItem in joinBuffer) {
        if (conditionMatched) return bufferItem, row2;
      }
    }

    flushBuffer();
  }
}
```

- 만약 드라이빙 테이블에 1000건의 데이터가 있고 Nested Loop Join 방식이라면 드리븐 테이블은 1000번 스캔해야 한다. 하지만 Join Buffer에 100건의 데이터를 넣는다고 가정하면 드리븐 테이블을 10번 스캔으로 종료할 수 있다.
- 인덱스가 없는 경우 Nested Loop Join 방식보다 나은 것이지 기본적으로 조인 쿼리가 Block Nested Loop Join으로 실행된다면 조인이 인덱스 없이 동작한다는 의미로 튜닝의 대상이 된다.
- Block Nested Loop 조인으로 동작하는지 여부는 실행계획에서 Using join buffer로 확인할 수 있다.
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/explain-bnl.png" width="800">

## hash 조인

- MySQL 8.0.20 버전부터 block nested loop는 없어지고 기본적으로 hash join을 사용한다.
- hash join은 인덱스가 있을때도 사용될 수 있고 인덱스가 없을 때도 block nested loop 대신 사용된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/hash-join.png" width="800">

- Hash 조인의 동작 방식은 우선 Build 단계에서 하나의 테이블의 조인되는 컬럼을 해시 함수로 돌리고 인 메모리에 캐시에 넣는다. 그리고 Probe 단계에서 다른 테이블의 조인되는 컬럼을 해시 함수로 수행 후 캐시에 저장된 값이 일치하는 경우 결과에 포함하는 방식으로 동작한다.
- 만약 Build 단계에서 데이터가 버퍼 메모리보다 큰 경우 디스크에 쓰기 시작한다.
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/hash-join2.png" width="800">

- 디스크에 데이터를 다 적재하면 각 청크 파일마다 비교를 통해 매칭되는 결과를 클라이언트에게 반환한다.
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/hash-join3.png" width="800">

### hash 조인 성능

- 인덱스가 없을 경우 Hash Join VS Nested Loop Join
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/hash-join-performance.png" width="800">

- 인덱스가 있을 경우 Hash Join VS Nested Loop Join
  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/hash-join-performance2.png" width="800">
