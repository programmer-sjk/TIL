# kafka가 File System을 쓰는데도 빠른 이유

## 최신 OS와 순차 I/O

- 일반적으로 디스크는 느리다라는 인식이 있는데, 어떻게 사용하는지에 따라 느릴수도 있고 빠를 수도 있다.
- 아래 그림처럼 디스크에 순차적으로 데이터에 접근하는 속도는 랜덤 엑세스에 비해 150,000배 빠르고 메모리에 랜덤 엑세스 하는 것보다 빠르다.

  - <img src="https://github.com/programmer-sjk/TIL/blob/main/images/devops/disk_io_performance.png" width="500">

- 순차 읽기/쓰기는 예측가능한 패턴으로 OS에서 크게 최적화가 가능하고, 최신 OS들은 read-ahead와 write-behind 같은 기술을 제공해 순차 읽기/쓰기 작업이 더 빠르게 수행되도록 지원한다.
- 카프카는 데이터를 메시지 큐 방식으로 저장하는데, 이는 순차 I/O 혜택을 볼 수 있어 빠른 성능을 제공한다.
