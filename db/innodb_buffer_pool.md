# MySQL InnoDB Buffer Pool

- MySQL의 InnoDB buffer pool을 정리하자.

## Buffer Pool 이란?

- 메인 메모리에 할당된 공간으로 테이블의 데이터와 인덱스 데이터가 접근되면 이를 cache 하는 공간이다.
- 자주 사용되는 데이터를 cache해 속도를 향상시키며, MySQL 전용 서버에서는 물리 메모리의 80%를 Buffer Pool에 할당하기도 한다.

## Buffer Pool 구조

- 빠른 읽기 성능을 위해 Buffer Pool은 페이지 단위로 나뉘며, cache 관리를 효율적으로 하기 위해 링크드 리스트로 구현되었다.
- LRU(Least Recently Used) 알고리즘을 사용하여 사용되지 않은 데이터는 제거된다.

