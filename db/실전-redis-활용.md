# 실전! Redis 활용

- [인프런 강의](https://www.inflearn.com/course/%EC%8B%A4%EC%A0%84-redis-%ED%99%9C%EC%9A%A9/dashboard)

## Redis 소개

- Redis란 Remote Dictionary Server의 약자

### Redis 특징

- Redis는 C언어로 작성된 오픈소스 인메모리 데이터 저장소
- 메모리는 디스크에 비해 매우 빠르기 때문에 RDB와는 다른 구조와 특성을 갖는다.
- Redis는 싱글 스레드로 동작하여 단순한 디자인을 채택하였다.
- RDB(Redis DataBase) + AOF(Append Only File)을 통해 영속성을 제공할 수 있다.
- pub/sub 패턴을 지원해 손쉽게 어플리케이션 개발 가능

### Redis 장점

- 모든 데이터를 메모리에 저장하기 때문에 매우 빠른 읽기/쓰기 속도 보장
- Redis가 지원하는 다양한 Data Type을 활용해 다양한 기능 구현이 가능

### Redis 사용사례

- 캐싱 / Rate Limiter / 메시지 큐
- 실시간 분석 & 계산
  - 순위표 (랭킹)
  - 반경 탐색
  - 방문자 수 계산
- 실시간 채팅
  - Pub/Sub

### Redis 영속성

- Redis는 주로 캐시로 사용되지만 데이터 영속성을 위한 옵션 제공
- Redis DataBase
  - 특정 시간에 스냅샷을 생성하는 기술
  - 장애나 복제시 주로 사용됨
  - 새로운 스냅샷이 생성되기 전 데이터 유실될 수 있음
  - 스냅샷을 생성하는 동안 Redis 성능 저하가 발생
- Append Only File
  - Redis에 쓰기 작업을 모두 log에 저장
  - 데이터 유실 위험이 적지만 RDB 보다 느림
