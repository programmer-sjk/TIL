# Redis를 활용해 동시성 문제 해결하기

- 해당 문서에서는 Redis와 동시성에 대해 아래 3 가지 방법을 중점적으로 작성한다.
  - 동시성 문제를 재현한다.
  - Redis의 set nx 명령어를 이용해 동시성 문제를 해결한다.
  - Redis의 Red Lock을 이용해 동시성 문제를 해결한다.
- 문서에 나오는 전체 코드는 [여기서](https://github.com/programmer-sjk/nestjs-redis) 확인할 수 있다.
