# MySQL InnoDB

- MySQL 서버는 **`사람의 머리 역할을 하는 MySQL 엔진과`** **`손발 역할을 담당하는 스토리지 엔진으로 구분할 수 있다`**.
  - MySQL 엔진은 요청된 SQL을 분석하거나 최적화하는 등 DBMS의 두뇌에 해당하는 처리를 한다.
  - 스토리지 엔진은 실제 데이터를 디스크에 저장하거나 데이터를 읽어오는 역할을 맡는다.
- 이번 문서에서는 **`스토리지 엔진에 속하는 InnoDB의 특성을 살펴본다`**.

## 프라이머리 키에 대한 클러스터링

- InnoDB의 모든 테이블은 **`기본적으로 PK를 기준으로 클러스터링 되어 저장된다`**. 즉 PK 순서대로 디스크에 저장된다는 의미이며, 모든 세컨더리 인덱스는 리프 노드에 레코드의 주소 대신 PK 값을 가리킨다. PK를 활용한 조회나 범위 검색이 매우 빠르기 때문에 다른 보조 인덱스에 비해 PK가 선택될 확률이 높다.

## 외래 키 지원

- 외래 키 기능은 InnoDB 스토리지 엔진 레벨에서 지원하는 기능으로 **`서버 운영의 불편함 때문에 실무에서 생성하지 않는 경우도 많다`**.
- 외래 키는 부모와 자식 테이블 모두 해당 컬럼에 인덱스를 생성하고, 변경 시 부모 테이블이나 자식 테이블에 데이터가 있는지 체크하는 작업이 필요해 잠금이 여러 테이블로 전파되어 데드락이 발생할 때가 많으므로 외래 키의 존재에 주의한 것이 좋다.
- 또한 수동으로 데이터를 적재하거나 스키마 변경 등의 관리 작업이 실패할 수 있다. 외래키가 복잡하게 얽혀있다면 테이블 관계를 명확히 파악해서 순서대로 작업해야 하지만, 긴급하게 조치를 해야 하는데 이런 문제가 발생하면 더 조급해질 수 있다.

## MVCC (Multi Version Concurreny Control)

- **`MVCC의 가장 큰 목적은 잠금을 사용하지 않는 일관된 읽기를 제공하는데 있다`**. InnoDB는 **`언두 로그를 이용해 이 기능을 구현한다`**.
- **`멀티 버전은 한 레코드에 대해 여러 버전이 동시에 관리된다는 의미다`**. 아래 그림은 처음 INSERT를 한 경우이다. 버퍼풀과 디스크에 동일한 데이터가 추가된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/mvcc-step1.png" width="400">

- 아래는 area가 경기로 업데이트 될 때의 상황을 보여준다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/mvcc-step2.png" width="400">

- UPDATE 문장이 실행되면 버퍼 풀은 즉시 내용이 수정되며, 디스크는 InnoDB의 Write 쓰레드에 의해 업데이트 될 수도 있고 아닐수도 있다. 다만 InnoDB가 ACID를 보장하기 때문에 일반적으로 InnoDB 버퍼 풀과 디스크는 동일한 상태라고 가정해도 무관하다.
- 위 상태에서 아직 `commit, rollback`이 완료된 상태가 아니라고 가정하면 **`조회했을 때 어떤 데이터를 조회할까? 이에 대한 답은 DB 격리 수준에 따라 다르다`**. READ_UNCOMMITED 격리 수준은 버퍼 풀의 데이터를 조회하고 READ_COMMITED 이상의 격리 수준은 아직 커밋되지 않았기 때문에 InnoDB 버퍼 풀이나 디스크가 아닌 변경되기 이전 내용을 보관하는 언두 영역의 데이터를 반환한다. **`이러한 과정을 DBMS는 MVCC라고 한다`**.
- COMMIT을 완료하면 변경된 상태를 영구적인 데이터로 만들어 버린다. 반대로 Rollback을 실행하면 InnoDB는 언두 영역에 있는 백업된 데이터를 InnoDB 버퍼 풀로 복구하고, 언두 영역의 내용을 삭제해버린다. **`커밋이 된다고 언두 영역의 데이터가 바로 삭제되는 것은 아니다`**. 이 언두 영역의 데이터를 필요로 하는 트랜잭션이 더는 없을 때 삭제된다.

## 잠금 없는 일관된 읽기 (Non-Locking Consistent Read)

- **`InnoDB 엔진은 MVCC를 이용해 잠금을 걸지 않고 읽기 작업을 수행한다`**. 격리 수준이 SERIALIZABLE이 아닌 이상 **`순수한 SELECT 작업은 다른 트랜잭션의 변경 작업과 관계없이 잠금을 대기하지 않고 바로 실행된다`**.
- 어떤 레코드를 업데이트 하고 아직 `commit, rollback`이 되지 않은 상태에서 순수한 **`SELECT 문은 격리 수준에 따라 버퍼 풀을 보거나 언두 영역의 데이터를 가져간다. 이를 잠금 없는 일관된 읽기라고 표현한다`**.
- 트랜잭션이 오랫동안 길어지면 언두 로그를 빨리 삭제하지 못하기 때문에 **`트랜잭션이 시작됐다면 가능한 빨리 롤백이나 커밋을 통해 트랜잭션을 완료하는게 좋다`**.

## 자동 데드락 감지

- InnoDB 엔진은 내부적으로 잠금이 교착 상태에 빠지지 않았는지 체크하기 위해 잠금 대기 목록을 그래프 형태로 관리한다. InnoDB 엔진은 **`데드락 감지 스레드를 가지고 있는데 이 녀석이 주기적으로 잠금 대기 그래프를 검사해 교착 상태에 빠진 트랜잭션들 중 하나를 강제 종료한다`**. 이때 어떤 트랜잭션을 먼저 강제 종료할지 기준은 언두 로그 레코드를 더 적게 가진 트랜잭션이다.
- 언두 로그 레코드를 적게 가졌다는 이야기는 롤백을 해도 언두 처리를 해야 할 내용이 적다는 것이며, 트랜잭션 강제 롤백으로 인한 MySQL 서버의 부하도 덜 유발하기 때문이다.
- InnoDB는 상위 레이어인 MySQL 엔진에서 관리되는 테이블 잠금을 볼 수가 없어 데드락 감지가 불확실 할 수도 있는데 **`innodb_table_locks`** 시스템 변수를 활성화하면 테이블 레벨의 잠금까지 감지할 수 있게 된다. **`별 다른 이유가 없다면 innodb_table_locks 시스템 변수를 활성화하자`**.
- 일반적인 서비스에서 데드락 감지 스레드가 데드락을 찾아내는 작업은 크게 부담되지 않는다. 하지만 **`동시성이 너무 높아지거나 각 트랜잭션이 가진 잠금의 개수가 많아지면 데드락 감지 스레드가 느려진다`**. 이를 해결하기 위해 MySQL은 **`innodb_deadlock_detect 시스템 변수를 제공하며, 이 값이 OFF일 경우 데드락 감지 스레드는 더 이상 동작하지 않는다`**. 데드락이 걸릴 경우 무한정 대기하게 되는 문제가 있으나, innodb_lock_wait_timeout 시스템 변수를 활성화하면 데드락이 일정 시간이 지나면 자동으로 실패하게 된다. 따라서 innodb_deadlock_detect를 OFF로 비활성화 하는 경우 innodb_lock_wait_timeout을 기본값인 50초보다 훨씬 낮은 시간으로 변경해서 사용할 것을 권장한다.
- 실제 구글은 **`아주 높은 빈도로 실행되는 서비스에서 매우 많은 트랜잭션이 동시에 실행되기에 데드락 감지 스레드가 성능을 저하시킨 다는 것을 알아냈다`**. 구글은 자체적으로 MySQL 서버의 코드를 변경해 데드락 감지 스레드를 비활성화할 수 있게 설정해서 사용했다. 이를 인지하고 오라클이 MySQL에 요청을 해서 추가된 환경변수가 innodb_deadlock_detect 이다.

## 언두 로그

- InnoDB 스토리지 엔진은 트랜잭션과 격리 수준을 보장하기 위해 DML(insert, update, delete)로 변경되기 이전 버전의 데이터를 별도로 백업한다. 이렇게 **`백업된 데이터를 언두 로그라고 부른다`**. 보통 아래의 경우 사용된다.
  - 트랜잭션이 롤백되면 변경된 데이터를 변경 전 데이터로 복구해야 하는데 이때 언두 로그에서 백업해서 복구한다.
  - 특정 커넥션에서 데이터를 변경 도중, 다른 커넥션에서 해당 데이터를 조회하면 트랜잭션 격리 수준에 맞게 언두 로그에 백업해둔 데이터를 읽기도 한다.

## 리두 로그

- 리두 로그는 트랜잭션의 ACID 중에서 D에 해당하는 영속성과 밀접하게 관련돼 있다. **`리두 로그는 HW, SW 등의 문제로 MySQL 서버가 비정상적으로 종료됐을 때 데이터 파일에 기록되지 못한 데이터를 잃지 않게 해 주는 안전장치다`**.
- MySQL 서버를 포함한 대부분 DBMS는 데이터 변경 내용을 로그로 먼저 기록한다. 데이터 쓰기 비용으로 인한 성능 저하를 막기 위해 **`쓰기 비용이 낮은 자료구조를 가진 리두 로그를 가지며`**, 비정상 종료가 발생하면 리두 로그의 내용을 이용해 데이터 파일을 서버가 종료되기 직전의 상태로 복구한다. ACID 만큼 성능도 중요하기 때문에 리두 로그를 버퍼링할 수 있는 로그 버퍼와 같은 자료구조도 가지고 있다.
- MySQL 서버가 비정상 종료되면 InnoDB 스토리지 엔진은 아래 두 가지 종류의 데이터를 가질 수도 있다.
  - 커밋됐지만 데이터 파일에 기록되지 않은 데이터
  - 롤백됐지만 데이터 파일에 이미 기록된 데이터
- 위에서 1번 경우 리두 로그에 저장된 데이터를 데이터 파일에 다시 복사하기만 하면 된다. 2번의 경우는 그 변경이 커밋됐는지, 롤백됐는지, 트랜잭션의 실행 중간 상태였는지 리두 로그를 통해 확인 후, 변경 전의 데이터를 가진 언두 영역의 내용을 가져와 데이터 파일에 복사한다.
