# 메시지 시스템들의 전송 방식

- 카프카를 공부하다 보면 적어도 한번, 정확히 한 번과 같은 키워드를 접하게 된다.
- 메시지 전송 방식의 `적어도 한 번 전송(at-least-once)`, `최대 한 번 전송(at-most-once)`, `중복 없는 전송`, `정확히 한 번 전송`을 알아보자.

## 적어도 한 번 전송 (at-least-once)

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/at-least-once.png" width="400">

- 위 그림은 **`적어도 한 번 전송`** 과정을 그림으로 나타낸 것으로 순서대로 살펴보자
  - 프로듀서가 브로커의 특정 토픽으로 메시지 A를 전송한다.
  - 브로커는 잘 받았다는 의미로 ACK를 프로듀서에게 응답한다.
  - 브로커의 ACK를 받은 프로듀서는 다음 메시지 B를 브로커에게 전송한다.
  - 브로커는 메시지를 받았지만 네트워크 오류나 브로커 장애로 인해 **결국 ACK가 프로듀서에게 전달되지 않는다.**
  - 메시지 B의 ACK를 받지 못한 프로듀서는 브로커가 메시지 B를 받지 못했다고 판단해 메시지 B를 재전송한다.
- 프로듀서가 메시지 B의 ACK를 받지 못한 시점에, **프로듀서 입장에선 브로커가 데이터를 받고 ACK를 전송하지 못한 건지 데이터를 받지 못해서 ACK를 전송하지 못한 건지 알 수가 없다.** 이때 적어도 한 번 전송 방식에 따라 메시지 B를 다시 전송한다. 만약 브로커가 데이터를 받지 못했다면 처음 메시지를 저장하고, 브로커가 데이터를 받은 상태라면 중복 저장이 된다.
- 이렇게 **일부 메시지 중복이 발생하더라도 최소한 하나의 메시지는 반드시 보장한다는 것이 적어도 한 번 전송 방식이며 카프카는 기본적으로 이 방식으로 동작**한다.

## 최대 한 번 전송 (at-most-once)

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/at-most-once.png" width="400">

- 위 그림은 **`최대 한 번 전송`** 과정을 그림으로 나타낸 것으로 순서대로 살펴보자
  - 프로듀서가 브로커의 특정 토픽으로 메시지 A를 전송한다.
  - 브로커는 잘 받았다는 의미로 ACK를 프로듀서에게 응답한다.
  - 브로커의 ACK를 받은 프로듀서는 다음 메시지 B를 브로커에게 전송한다.
  - 브로커는 메시지를 받았지만 네트워크 오류나 브로커 장애로 인해 **결국 ACK가 프로듀서에게 전달되지 않는다.**
  - 프로듀서는 브로커가 받았다고 가정하고 메시지 C를 전송한다.
- **`최대 한 번 전송은 **ACK를 받지 못해도 재 전송하지 않는다`**. 위 그림에서 사실 ACK를 받는 부분은 의미가 없지만 적어도 한 번 전송과 비교하기 위해 넣어두었다. 정리하면 **`최대 한 번 전송은`\*\* **메시지 손실을 감안하더라도 중복 전송은 하지 않는 경우**이다.
- 일부 메시지가 손실되더라도 높은 처리량을 필요로 하는 대량의 로그 수집이나 IoT 같은 환경에서 사용된다.

## 중복 없는 전송 (멱등성 프로듀서)

- 카프카의 0.11 버전부터 프로듀서가 메시지를 중복 없이 브로커로 전송할 수 있는 기능이 추가되었다.

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/no-duplicate.png" width="400">

- 위 그림은 **`중복 없는 전송`** 과정을 그림으로 나타낸 것으로 순서대로 살펴보자
  - 프로듀서가 브로커의 특정 토픽으로 메시지 A를 전송한다. 이때 **PID(Producer ID)와 메시지의 시퀀스 번호** 0을 같이 전달한다.
  - 브로커는 메시지를 저장하고 PID와 시퀀스 번호를 기록해둔다. 그리고 ACK를 프로듀서에게 응답한다.
  - 브로커의 ACK를 받은 프로듀서는 다음 메시지 B를 브로커에게 전송한다. PID는 동일하지만 시퀀스 번호는 증가한 1 값을 보낸다.
  - 브로커는 메시지를 저장하고 PID와 시퀀스 번호를 기록해둔다. 하지만 네트워크 오류나 브로커 장애로 인해 **결국 ACK가 프로듀서에게 전달되지 않는다.**
  - 메시지 B의 ACK를 받지 못한 프로듀서는 브로커가 메시지 B를 받지 못했다고 판단해 메시지 B를 재전송한다.
- **위 과정은 적어도 한 번 전송과 동일**하다. 하지만 프로듀서가 메시지 B를 재전송할 경우의 동작은 다르다. 브로커는 전달받은 데이터로부터 **PID와 시퀀스 번호를 확인하여 메시지 B가 이미 저장된 것을 확인하고 중복 저장하지 않고 ACK 응답만 보낸다.** 결국 메시지의 중복 저장을 피할 수 있게 된다.
- 메시지 중복을 피하기 위해 사용되는 PID와 시퀀스 번호는 브로커의 메모리에 유지되고, 리플리케이션 로그에도 저장된다. 따라서 예기치 못한 브로커의 장애로 리더가 변경되더라도 새로운 리더가 PID와 시퀀스 번호를 정확히 알 수 있어 중복 없는 메시지 전송이 가능하다.
- 컨플루언트 블로그 글에 의하면 중복 없는 전송의 경우 오버헤드로 인해 20% 정도의 성능이 감소했다고 하는데 이는 그렇게 높은 편은 아니다. 따라서 성능이 민감하지 않은 상황에서 중복 없는 메시지 전송이 필요하다면 이 방식을 적용해야 한다.
- 중복 없는 전송을 위해 프로듀서 설정 값의 일부를 변경해야 한다.
  - **`enable.idempotence:`** true로 설정
  - **`max.in.flight.requests.per.connection:`** ACK를 받지 않은 상태에서 하기본 값은 5이며, 5 이하로 설정해야 한다.
  - **`acks:`** all로 설정해야 한다.
  - **`retries:`** ACK를 못 받은 경우 재시도 해야 하므로 0보다 큰 값으로 설정
- 위 설정을 반영하여 **중복없는 전송을 제공하는 프로듀서를 멱등성 프로듀서**라고 부른다. 만약 멱등성 프로듀서로 동작하는 어플리케이션이 종료되고 재시작되면 PID가 달라진다. PID가 달라지면 브로커 입장에선 다른 프로듀서 어플리케이션에서 데이터를 보냈다고 판단하기 때문에 장애가 발생하지 않을 경우에만 중복없는 전송이 보장되는 것을 고려해야 한다.

## 정확히 한 번 전송 (exactly-once)

- **`정확히 한 번 전송은`** 브로커의 장애나 프로듀서의 재 전송에도 메시지의 유실이나 중복 없이 **데이터를 한 번 전송함을 의미**하며 **멱등성과 트랜잭션**을 통해 **`정확히 한 번 전송을`** 제공할 수 있다. 멱등성은 위에서 정리한 중복없는 전송과 동일하며 트랜잭션이란 여러 레코드를 원자(atomic) 단위로 처리하여 전체 성공하거나 전체 실패하는 것을 보장한다.
- 위에서 설명한 중복 없는 전송을 이해했다면 프로듀서가 데이터를 파티션에 중복없이 저장할 수 있다는 것을 이해했을 것이다. 그렇다면 컨슈머가 데이터를 가져갈 때 어떤 문제가 있을까?

### consumer 데이터 유실/중복

- 아래는 메시지를 유실할 수 있는 그림을 표현한다. 어플리케이션이 컨슘 후 offset을 commit 하고, 컨슘한 데이터로부터 비지니스 로직을 처리하는 도중 실패했을 떄를 보여준다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/consumer-data-loss.png" width="500">

- 아래는 메시지를 중복 처리할 수 있는 그림을 표현한다. 어플리케이션이 컨슘 후 비지니스 로직을 정상적으로 처리하고 offset commit에 실패할 경우 다시 컨슘하여 중복 처리할 수 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/consumer-duplicate.png" width="500">

### 데이터 파이프라인에서 consume -> process -> send 정확히 한 번 처리를 보장하는 방법

- **멱등성 프로듀서와 트랜잭션 프로듀서/컨슈머 기능**으로 **`A토픽->컨슈머->데이터 처리->프로듀서->B토픽`** 으로 이어지는 파이프라인에서 **정확히 한번**을 달성할 수 있게 되었는데 이를 위해 코드는 아래 흐름으로 진행된다.

```java
  // 카프카 프로듀서를 트랜잭션 프로듀서로 실행시키는 명령어 및 옵션.
  KafkaProducer producer = createKafkaProducer( "bootstrap.servers", "localhost:9092", "transactional.id", "my-transactional-id", "enable.idempotence", "true");

  // 시작과 동시에 트랜잭션 초기화를 진행
  producer.initTransactions();

  // 컨슈머의 격리수준을 read_committed로 실행하며 트랜잭션 커밋이 완료된 데이터만 읽을 수 있도록 한다.
  KafkaConsumer consumer = createKafkaConsumer( "bootstrap.servers", "localhost:9092", "group.id", "my-group-id", "isolation.level", "read_committed");

  consumer.subscribe(singleton("inputTopic"));

  while (true) {
    // 토픽으로부터 레코드를 읽는다.
    ConsumerRecords records = consumer.poll(Long.MAX_VALUE);

    // 메시지가 들어오면 트랜잭션을 시작한다.
    producer.beginTransaction();

    for (ConsumerRecord record : records) {
      // 들어온 레코드를 프로듀서에게 전송
      producer.send(producerRecord("outputTopic", record))
    }

    // 프로듀서가 consumer 코디네이터를 통해 __consumer_offsets에 offset을 증가시킨다.
    producer.sendOffsetsToTransaction(offsetMapFunction(), "my-group-id");

    // 트랜잭션 코디네이터가 commit 한다.
    producer.commitTransaction();
  }
```

- 위에서 한 가지 언급할 것은 **컨슈머 그룹의 오프셋을 증가시키는 책임이 프로듀서에 있다는 점**이다. 이 방법으로 데이터 소비 및 프로세스 처리, 오프셋 커밋이 진행되어 정확히 한 번 처리가 가능하다. 결과적으로 멱등성 + 트랜잭션으로 카프카는 정확히 한 번 전송을 보장하게 된다.
- kafka streams도 내부적으로 위와 동일한 방법으로 정확히 한 번 처리를 지원하며 이런 복잡한 내용 없이 **`processing.guarantee="exactly_once"`** 옵션 하나를 추가하는 방법으로 지원된다.

### consumer에서 정확히 한번 처리

- 토픽으로부터 **메시지를 받아 DB에 처리하는 어플리케이션**을 생각해보자. 여기에선 정확히 한 번 처리를 어떻게 해야 할까?
컨슈머는 레코드를 토픽에서 가져가서 처리한 뒤 중복처리를 막기 위해 offset 저장을 수행한다. 예를 들어 토픽의 레코드들을 DB에 저장한다고 가정할 경우 **데이터 insert와 commit offset을 하나의 트랜잭션으로 묶는 것은 불가능**하다. 서로 연동되는 서비스가 아니기 때문이다. 따라서 토픽->컨슈머 처리에서 **정확히 한번**을 달성하기 위해서는 **`멱등성(idempotence) 처리`** 를 하도록 가이드한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/message-queue/application-check.png" width="500">

## 정리하며

- 카프카를 사용할 때 **정확히 한 번을 달성하기 위해 각 구간별 사용해야 하는 기술**은 다음과 같다.
  - **`프로듀서 -> 토픽:`** 멱등성 프로듀서
  - **`토픽 -> 컨슈머 -> 프로듀서 -> 토픽:`** 트랜잭션 프로듀서 / 컨슈머
  - **`토픽 -> 컨슈머:`** 멱등성 처리

## 레퍼런스

- <https://www.confluent.io/blog/transactions-apache-kafka/>
- <https://blog.digitalis.io/read-process-write-with-kafka-transactions-29bc0a70febd>
- <https://www.confluent.io/kafka-summit-london18/dont-repeat-yourself-introducing-exactly-once-semantics-in-apache-kafka/>
