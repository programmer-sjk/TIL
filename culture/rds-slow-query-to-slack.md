# RDS Slow Query를 Slack 전송

- RDS의 Slow Query들을 Slack으로 전송
  - aws console을 직접 들어가거나 datadog으로 확인 가능하지만 수월한 모니터링을 위해 Slack 알림 전송
- 회사에서 serverless를 이용해 lambda를 관리하는 별도의 Repository가 존재
  - aws console에서 구독 필터를 설정하는게 아닌 코드에서 cloudwatch logs 구독 정보를 코딩
  - 따라서 이 문서에는 cloudwatch logs와 lambda 연동을 제외하고 진행

## 전체 흐름도

- 전반적인 흐름은 아래와 같다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/rds-slow-query/structure.png" width="600">
