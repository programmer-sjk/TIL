# 데브원영 아파치 카프카 for beginner

- [인프런 무료 강의](https://www.inflearn.com/course/%EC%95%84%ED%8C%8C%EC%B9%98-%EC%B9%B4%ED%94%84%EC%B9%B4-%EC%9E%85%EB%AC%B8)

## 아파치 카프카 개요

- 카프카는 source / target 시스템 간 결합도를 낮추기 위해 나왔다.
  - source application은 데이터를 카프카에게 보내고 target application은 카프카에서 데이터를 가져오면 된다.
- 카프카에는 각종 데이터를 담는 Topic이 있는데 쉽게 말해 큐라고 보면 된다.
  - Topic에 데이터를 넣는 역할은 producer가 하고, 데이터를 가져오는 역할은 consumer가 담당한다.
- 카프카는 fault tolerant, 즉 고가용성을 제공하고 서버에 장애가 생겨도 데이터를 손실없이 복구할 수 있다.
  - 또 낮은 지연과 높은 처리량을 통해 대용량의 데이터 처리가 가능하다.
