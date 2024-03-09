# kafka가 File System을 쓰는데도 빠른 이유

- 디스크 I/O의 비용이 상대적으로 높다는 것은 다들 알 것이다. 카프카는 데이터를
  브로커의 로컬 디스크에 저장하게 되는데, 그럼에도 카프카가 빠른 이유가 무엇인지 정리해본다.

## 최신 OS와 순차 I/O

- 디스크 I/O는 어떻게 사용하는지에 따라 느릴수도 있고 빠를 수도 있다. 아래 그림처럼 **디스크에 순차적으로 데이터에 접근하는 속도**는 디스크 랜덤 엑세스에 비해 150,000배 빠르고 **메모리에 랜덤 엑세스 하는 것보다도 빠르다**.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/disk_io_performance.png" width="500">

- 최신 OS들은 read-ahead와 write-behind 같은 기술을 제공해 순차 읽기/쓰기 작업이 더 빠르게 수행되도록 지원한다. 카프카는 데이터를 메시지 큐 방식으로 저장하는데, 이는 **순차 I/O 혜택을 볼 수 있어 빠른 성능을 제공**한다.

## Page Cache

- 디스크 검색을 줄이고 처리량을 높이기 위해 최신 운영체제들은 **Page Cache(디스크 캐시)** 를 위해 메인 메모리를 더 공격적으로 사용하게 되었다. 모든 디스크 읽기/쓰기는 페이지 캐시를 거치게 되며, 사용자나 어플리케이션에 의해 관리되는게 아니라 OS에 의해 관리되므로 사용자 영역과 커널 영역의 중복 저장 없이 2배 정도의 캐시를 저장할 수 있으며 어플리케이션 재 시작시에도 빠르게 캐시의 혜택을 볼 수 있다.

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/page-cache.png" width="500">

- 또 Kafka는 JVM 위에서 동작하는데 Heap 메모리에 객체를 저장하는 비용은 매우 비싸고, 힙 메모리가 커질수록 GC가 느려진다는 단점이 있다. 결론적으로 순차 읽기/쓰기 혜택을 볼 수 있는 File **System과 page cache**를 사용하는게 메모리 캐시나 다른 구조를 사용하는 것보다 좋은 성능을 낼 수 있었다.

## Zero Copy

- 일반적으로 **네트워크를 통해 데이터를 전달하는 절차**는 아래와 같다.
  - OS는 디스크로부터 데이터를 읽어 커널 영역의 page cache에 저장한다.
  - 어플리케이션은 page cache의 데이터를 사용자 영역으로 읽어온다.
  - 어플리케이션은 커널 영역에 있는 socket buffer로 데이터를 쓴다.
  - OS는 socket buffer에 있는 데이터를 NIC buffer로 복사하고 네트워크를 통해 전송한다.
- 위 과정에서 불필요한 복사와 system call이 발생하는데 **OS가 제공하는 sendFile 함수는** 커널 영역의 Page Cache에서 NIC Buffer로 직접 복사가 가능해 **효율적으로 데이터를 전송**할 수 있다.
- Kafka는 **`Zero Copy`** 기술을 이용해 메시지가 생성/소비될 때 불 필요한 복사와 system call을 줄여 효과적으로 데이터를 전송한다. 한 건의 경우 성능에 큰 영향은 없겠지만 수 십, 수 백만 이상의 대량의 데이터가 카프카를 통해 전달될 경우 성능의 차이를 실감할 수 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/zero-copy.png" width="1000">

## 레퍼런스

- [LinkedIn](https://www.linkedin.com/pulse/why-kafka-so-fast-aman-gupta)
- [Confluent 문서](https://docs.confluent.io/kafka/design/file-system-constant-time.html)
