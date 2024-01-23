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

  - From Scratch 버튼 클릭

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step1.png" width="400">

  - 생성하는 app 이름과 workspace를 지정한다.

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step2.png" width="400">

  - 왼쪽 Feature -> OAuth & Permissions 클릭

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step3.png" width="400">

  - 단순히 메시지 전송이라면 `char:write` 으로 충분하지만 Slack에서 사진과 이름까지 커스터마이징 할 수 있는 `chat:write:customize`를 클릭한다

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step4.png" width="400">

  - Basic Information 페이지에서 Install to Workspace 버튼을 클릭한다.

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-step5.png" width="400">

- 만약 아래와 같이 `앱에 설치할 봇 사용자가 없습니다` 메시지가 뜬다면

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-problem.png" width="400">

  - App Home 페이지에서 App Display Name 옆에 Edit 버튼을 클릭한다.

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-solve.png" width="400">

  - 아래처럼 원하는 Name과 username을 저장하고 Install to Workspace 버튼을 다시 클릭한다

    <img src="https://github.com/programmer-sjk/TIL/blob/main/images/culture/pr-reminder/slack-bot-solve2.png" width="400">
