# AWS Aurora RDS vs AWS RDS

- Aurora RDS와 RDS의 차이점을 비교해본다.

## Aurora RDS

- Aurora RDS는 기존 RDS를 성능과 간편성, 가용성 관점에서 AWS에 의해 한 번 디자인된 RDBMS
- Aurora와 RDS의 가장 큰 차이점은 스토리지이다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/db/aurora-storage.png" width="400">

- 위 그림에서 알 수 있는 몇 가지 특징을 정리해보자.
  - 한 대의 aurora 서버를 운영하더라도 3개의 AZ 영역에 총 6개의 스토리지에 데이터가 저장된다.
  - Aurora는 Shared 스토리지를 사용하며, DB 인스턴스와 분리되어 있다.
  - 마스터 DB에 새로운 데이터가 들어오면 6개의 스토리지에 바로 데이터를 저장한다.
