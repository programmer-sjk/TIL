# 코드 리뷰 리마인더 봇

## 리뷰 리마인더의 필요성

- MSA로 **코드가 여러 Repository에서 관리**되는 환경에서 모든 PR을 리뷰해야 하는 환경에 있다.
- 리뷰를 한 적이 없는 PR이라면 [review-requested](https://github.com/pulls/review-requested) 페이지에서 확인이 가능하지만 comment를 달았던 PR은 이 페이지에서 확인되지 않는다.
- 매일 출근해서 리뷰를 하는데 두 단계로 나뉘어진다.
  - [review-requested](https://github.com/pulls/review-requested) 페이지에서 나에게 요청은 왔으나 리뷰를 한 적 없는 PR을 리뷰한다.
  - github 각 Repository에서 pulls 페이지를 들어가 comment를 단 적이 있는 PR을 찾아 리뷰한다.
- 매일 하는 리뷰인데 이 과정들이 너무 귀찮아서 매일 아침 리뷰 목록을 알려주는 봇을 만들기로 결심했다.

## 요구사항

- MSA로 분리된 여러 Repository 정보에 접근할 수 있어야 한다.
- 내가 리뷰어로 할당 된 PR 중 승인한 PR을 제외한 모든 PR 목록을 얻어와야 한다.
- PR Label에 접근하여 Label을 수정할 수 있어야 한다.
  - 최대 언제까지 리뷰를 해달라는 `D-0`, `D-1`, `D-2`, `D-3` Label을 생성
  - 하루가 지나면 자동으로 `D-3 -> D-2 -> D-1 -> D-0` 으로 업데이트 하는 기능을 고민 중 이었음
- 리뷰를 도와주는 util 기능이기 때문에 안정성이나 성능이 중시되진 않는다.
  - 이것 말고도 해야 할게 많다
- 특정 Slack 채널에서 자신만 볼 수 있어야 한다.
  - 각자 일정에 따라 바쁜 날에는 리뷰를 적게 하고, 어떤 날에는 리뷰를 많이 할 수 있다.
  - 각자 PR 쌓인게 모두에게 공개되면 누구에게는 압박이 될 수도 있다는 생각이 들어 Slack 채널에서 자신만 볼 수 있어야 한다.

## 사용할 기술 선정

- 찾아보니 [PyGithub](https://github.com/PyGithub/PyGithub)가 제일 유명하고 사용하기 편해보였다.
  - Python으로 뚝딱뚝딱 만들 수 있겠다고 생각했지만, 사내 기술 스택은 JS이고 Python을 사용해 본 적 없는 분들도 있음
  - 내가 아니어도 편하게 유지보수가 가능하도록 JS 스택으로 개발하기로 결정
- `js github api`로 검색하며 아래 두 개의 라이브러리로 추렸다.
  - [Octokit](https://docs.github.com/en/rest/guides/scripting-with-the-rest-api-and-javascript?apiVersion=2022-11-28)
    - [github 주소](https://github.com/octokit/octokit.js)
  - [Github api](https://github.com/github-tools/github)
- Octokit은 Github REST API를 사용할 수 있는 SDK로 Github에 의해 관리된다.
- Github API는 Github REST API와 연동을 쉽게 해주는 라이브러리로 Node와 브라우저에서 사용이 가능하다.
  - Github API도 결국 내부적으로 REST API를 사용
- 둘 중 문서화가 좀 더 깔끔한 Octokit을 사용하기로 결정했다.

## Slack 메시지 전송

- 라이브러리는 사용 목적에 충분한 기능을 제공하는 `@slack/web-api` 모듈을 사용했다.
- 테스트를 위해 Slack workspace를 새로 만들고 app을 생성한다.

### 메시지를 보내기 위한 Slack 토큰 생성 과정

- From Scratch 버튼 클릭

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step1.png" width="400">

- 생성하는 app 이름과 workspace를 지정한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step2.png" width="400">

- 왼쪽 Feature -> OAuth & Permissions 클릭하고 Bot token을 생성한다.

  - Bot Token으로 생성시 전달된 메시지는 workspace에 설치된 app에 의해 전송된다.
  - User Token은 워크스페이스 멤버를 의미하며 전달된 메시지는 나에 의해 전송된다.

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step3.png" width="500">

  - Bot Token(위) vs User Token(아래)로 메시지 전송시 비교. 다른 사람들은 내가 보낸 메시지로 보여진다.
    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/bot-vs-user.png" width="400">

- 단순히 메시지 전송이라면 `char:write` 으로 충분하지만 Slack에서 사진과 이름까지 커스터마이징 할 수 있는 `chat:write:customize`를 클릭한다

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step4.png" width="400">

- Basic Information 페이지에서 Install to Workspace 버튼을 클릭한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step5.png" width="400">

- 만약 아래와 같이 `앱에 설치할 봇 사용자가 없습니다` 메시지가 뜬다면

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-problem.png" width="400">

- App Home 페이지에서 App Display Name 옆에 Edit 버튼을 클릭한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-solve.png" width="600">

- 아래처럼 원하는 Name과 username을 저장하고 Install to Workspace 버튼을 다시 클릭한다

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-solve2.png" width="400">

- 정상적으로 토큰을 발급하면 아래와 같이 Token을 확인할 수 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step6.png" width="600">

### Slack 토큰으로 메시지를 받기 위한 설정

- 만약 User Token을 사용한다면 채널에 참여하고 발급받은 토큰으로 메시지를 전달받으면 된다.
- 만약 Bot Token을 사용한다면 채널 세부정보 -> 통합 -> 앱 추가 버튼을 클릭한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-setting1.png" width="400">

- 생성한 app이 목록에 나타난다면 추가 버튼을 눌러준다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-setting2.png" width="400">

- 만약 나오지 않는다면 Slack 하단에 앱에서 세부정보를 클릭한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-setting3.png" width="400">

- 앱을 채널에 추가하기 버튼을 클릭하고 원하는 채널에 앱을 추가한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-setting4.png" width="400">

- 그 후 발급받은 토큰을 `@slack/web-api`에서 사용할 수 있는 함수에 넣어주고 메시지를 보내면 Slack에서 정상적으로 메시지를 확인할 수 있다.

  ```js
  import { WebClient } from "@slack/web-api";

  const web = new WebClient("slack에서 발급받은 토큰");
  const result = await web.chat.postMessage({
    text: "이 메시지가 Slack에 전달됩니다.",
    channel: "랜덤",
  });
  ```

- 메시지 전송 결과

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-message.png" width="400">

## Avatar와 username 커스터마이징

- Slack에서 노출되는 Avatar와 이름을 변경하고 싶다면 icon_emoji, username 인자를 활용하면 된다.

  ```js
  const web = new WebClient("slack에서 발급받은 토큰");
  const result = await web.chat.postMessage({
    text: "이 메시지가 Slack에 전달됩니다.",
    channel: "랜덤",
    icon_emoji: "cubimal_chick",
    username: "리뷰 비서",
  });
  ```

- 아래 화면처럼 메시지가 전송된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-message2.png" width="400">

## 각자의 PR 목록을 자신만 보고 싶다면?

- 각자의 PR 목록이 공개된 채널에 노출되어 개발팀 전체가 본다고 가정하자.
- 리뷰어 A,B에게 동일한 리뷰가 요청왔는데, 업무 부하에 따라 누군가는 남은 리뷰가 적고, 누군가는 남은 리뷰가 많을 수 있다.
  - 여기서 민감하신 분들은 리뷰를 적게 하는 것처럼 보일까봐 스트레스 일 수도 있다. (워낙 다양한 사람들이 일하고 있기 때문에)
- 만약 특정 채널에서 자신에게 남은 PR을 오직 자신만 볼 수 있게 하려면 어떡해야 할까?
- postEphemeral 함수를 사용하면 된다.

  ```js
  const web = new WebClient("slack에서 발급받은 토큰");
  const result = await web.chat.postEphemeral({
    text: "이 메시지가 Slack에 전달됩니다.",
    channel: "랜덤",
    user: "SLACK ID",
    icon_emoji: "cubimal_chick",
    username: "리뷰 비서",
  });
  ```

- 여기서 특정 유저의 Slack ID는 프로필을 클릭하고 멤버 ID 복사를 클릭한다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-user-id.png" width="400">

- 전달된 메시지는 아래와 같이 표기된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-message-only-me.png" width="400">
