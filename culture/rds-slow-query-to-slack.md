# RDS Slow Query를 Slack 전송

- RDS의 Slow Query들을 Slack으로 전송
  - aws console을 직접 들어가거나 datadog으로 확인 가능하지만 수월한 모니터링을 위해 Slack 알림 전송
- 회사에서 serverless를 이용해 lambda를 관리하는 별도의 Repository가 존재
  - aws console에서 구독 필터를 설정하는게 아닌 코드에서 cloudwatch logs 구독 정보를 코딩
  - 따라서 이 문서에는 cloudwatch logs와 lambda 연동을 제외하고 진행

## 전체 흐름도

- 전반적인 흐름은 아래와 같다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/rds-slow-query/structure.png" width="600">

## 데이터 확인

- cloudwatch logs와 lambda가 연동된 상태에서 sleep(10) 쿼리를 실행하면 아래와 같이 인코딩 + 압축된 데이터가 lambda로 전달된다.

  ![](rds-slow-query/raw.png)

- 원본 데이터를 얻기 위해 디코딩 + 압축을 풀어보자.

  ```ts
  const zlib = require('zlib');
  const payload = Buffer.from('H4sIAAAAAAAA/21Q72~~~~3M3sCAAA=', 'base64');
  const plain = JSON.parse(zlib.unzipSync(payload).toString());

  console.log(plain);
  ```

- 파싱한 데이터는 아래와 같다.

  ```json
  {
    "messageType": "DATA_MESSAGE",
    "owner": "097284161819",
    "logGroup": "testLogGroup",
    "logStream": "testLogStream",
    "subscriptionFilters": ["testSubscriptionFilters"],
    "logEvents": [
      {
        "logEvents": [
          {
            "id": "eventId",
            "timestamp": 1715597834471,
            "message": "# Time: 2024-05-13T10:57:14.471456Z\n# User@Host: 계정[계정] @  [127.0.0.1]  Id: 581102\n# Query_time: 10.002404  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1\nuse HOST;\nSET timestamp=1715597824;\nselect sleep(10);"
          }
        ]
      }
    ]
  }
  ```

- 데이터 구조를 확인했으니 본격적으로 파싱해서 Slack에 전달하는 로직을 작성해보자.

## Lambda 코드 작성
