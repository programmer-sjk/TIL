# MySQL InnoDB Buffer Pool

- MySQL의 InnoDB buffer pool을 정리하자.

## Buffer Pool 이란?

- 메인 메모리에 할당된 공간으로 테이블의 데이터와 인덱스 데이터가 접근되면 이를 cache 하는 공간이다.
- 자주 사용되는 데이터를 cache해 속도를 향상시키며, MySQL 전용 서버에서는 물리 메모리의 80%를 Buffer Pool에 할당하기도 한다.

## Buffer Pool 구조

- 빠른 읽기 성능을 위해 Buffer Pool은 페이지 단위로 나뉘며, cache 관리를 효율적으로 하기 위해 링크드 리스트로 구현되었다.
- LRU(Least Recently Used) 알고리즘을 사용하여 사용되지 않은 데이터는 제거된다.

## Buffer Pool LRU 알고리즘

- 새로운 페이지가 Buffer Pool에 추가되면 아래 그림과 같이 전체 링크드 리스트에서 중간 부분(Old SubList의 head)에 추가된다.
  - <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/innodb_buffer_pool.png" width="400">
- Buffer Pool 리스트는 크게 `New Sublist, Old Sublist`로 나뉘며, 이름에서 짐작할 수 있듯이 New는 최근 접근된 데이터가 유지되며 Old는 사용되지 않은 데이터의 리스트로 tail 쪽의 데이터는 최종적으로 제거될 수 있다.
- 그림에서 볼 수 있듯, `New Sublist`는 버퍼 풀의 5/8 정도를 사용하고, `Old Sublist`는 3/8을 사용한다.
- Buffer Pool의 LRU 알고리즘 동작에 대해 살펴보자.
  - InnoDB가 새로운 페이지를 Buffer Pool에 추가할 때는 `New Sublist`의 tail과 `Old Sublist`의 Head가 만나는 중간 지점에 추가한다. 정확히는 `Old SubList`의 head 이다.
  - `Old Sublist`에 있는 Page가 읽히면 New Sublist로 이동하게 된다. 만약 Page가 사용자 쿼리에 의해 읽혔다면 바로 New Sublist로 이동하지만 read ahead에 의해 읽힌다면 이동하지 않는다.
  - Buffer Pool에서 사용되지 않은 Page는 결국 `Old Sublist`의 tail에 위치하게 되며, 결국 제거된다.
