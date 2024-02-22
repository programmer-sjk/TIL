# Random Lunch 서포터

- HR 담당자분이 바쁠 경우 랜덤 런치가 진행되지 않음
- 조금이라도 수고를 줄이고자 자동화를 조금 지원
- 예시 코드는 [여기](https://github.com/programmer-sjk/random-lunch-support/blob/main/random-lunch.js)서 확인이 가능하다.

## 제안 방식

- HR 담당자에게 두 가지 방법을 제시

### 엑셀

- 엑셀에 전사 팀원의 닉네임과 오늘 출근여부를 기록
- 매 주 랜덤 런치를 하는 날, 오늘의 출근 여부만 갱신
- 프로그램을 실행시키면 출근 여부에 값이 있는 팀원들을 랜덤하게 배치
- 장/단점
  - **`장점`**: 기존에 랜덤 런치하는 프로세스랑 비슷하며 섞는 부분만 자동화
  - **`단점`**: 오늘 출근한 사람들을 HR 담당자가 파악해야 함

### Slack

- Slack에 Bot을 이용해 오늘 출근했냐는 질문을 전 팀원에게 보내고 Yes를 클릭한 팀원들을 저장
- 점심 시간 30분 전에 Yes를 클릭한 팀원들을 랜덤하게 배치
- 장/단점
  - **`장점`**: HR 담당자가 수동으로 하는 작업이 거의 줄어들음
  - **`단점`**: 급한 회의/티타임/업무 등으로 메시지를 놓친 팀원들을 추가적으로 고려하는 상황이 존재

### HR 담당자의 선택

- **`엑셀을 활용하는 방법을 선호`**
- 회사 규모가 25명 정도이고 재택이 활성화 되서 출근 인원을 세는게 어렵지 않음
- 현재 규모에선 오히려 Slack을 활용하는 방법이 추가 작업들이 존재할 것으로 예상

## 요구사항

- Mac에서 개발한 프로그램이 Windows에서 동작해야 한다. (HR 담당자 PC가 윈도우)
  - **`HR 담당자가`** NodeJS를 설치하거나 명령어를 외워서 실행하지 않고 **`Windows 프로그램을 실행한다`**.
- 엑셀 파일로부터 닉네임과 출근 여부를 읽어야 한다.
- 출근 여부에 데이터가 있는 경우 3-4명씩 팀이 되도록 랜덤하게 배치한다.
- Windows 클립보드에 복사해서 Slack에 `Ctrl+v`만 누르면 복사되어야 한다.
- 랜덤하게 섞인 멤버들을 보고 재 배치할 수 있어야 한다.

## 주요 개발 내용

### Mac에서 만든 Node JS파일을 Windows에서 실행시키기

- NodeJS 파일을 Windows 프로그램으로 변환할 수 있는 방법으로 3가지를 고민했다.
  - [pkg](https://github.com/vercel/pkg)
  - [node-webkit](https://nwjs.io/)
  - [electron](https://www.electronjs.org/)
- 위 3가지 방법을 비교하다가 `electron, node-webkit`은 제외했다.
  - `electron, node-webkit`은 Chromium이랑 NodeJS가 같이 빌드가 된다.
  - 만드려는 프로그램 기능은 간단한데, 빌드하면 기본적으로 **`100MB가 넘는 용량을 가지기 떄문에`** 제외했다.

### pkg

- [pkg](https://github.com/vercel/pkg)는 실행 가능한 프로그램을 만들어주는 패키지다.
- 만들어진 프로그램은 각 머신에서 별도의 NodeJS 설치 없어도 실행이 가능하다.
- 아래와 같이 설치하고 windows에서 동작하는 실행 프로그램을 생성한다.
  - `npm install pkg --dev`
  - `pkg -t node16-win-x64 random-lunch.js`
- 내 경우 js 파일의 라인 수가 140줄 정도 되었는데 **`윈도우 실행 프로그램의 크기는 45MB였다`**.

### 주요 로직

- `random-lunch.js` 파일에 작성된 주요 로직은 아래와 같다.
  - js와 같은 경로에 있는 `random_lunch.xlsx` 파일을 읽어 멤버와 참석여부 데이터를 가져온다.
  - `이름과 참석 여부`, 둘 중 하나라도 데이터가 없으면 필터링한다.
  - 참석하는 멤버들을 3,4명을 기준으로 섞어서 조를 배치하고 화면에 보여준다.
  - r 키를 누르면 랜덤 런치 멤버를 재 배치하고 보여주고, 그 외의 키를 입력하면 `클립보드에 결과를 복사한다`.
- 인사 담당자는 멤버를 확인하고 아무 키나 입력해 프로그램을 종료, Slack 채널에 그대로 복사해서 공유한다.

## 실행 결과

<img src="https://github.com/programmer-sjk/random-lunch-support/blob/main/images/result.png" width="550">

## Trouble Shooting

### ESM 문제

- 프로그램이 윈도우에서 실행 하자마자 종료되어 확인해보니 빌드시 아래와 같이 `warning` 메시지를 볼 수 있었다.

```txt
pkg -t node16-win-x64 random-lunch.js
> pkg@5.8.1
> Warning Babel parse has failed: import.meta may appear only with 'sourceType: "module"' (5:45)
> Warning Babel parse has failed: import.meta may appear only with 'sourceType: "module"' (6:45)
```

- 검색 후 [해당 이슈](https://github.com/vercel/pkg/issues/1291)를 보고 **`import 대신 require를 사용했다`**.
- 설치한 라이브러리 중 ESM만 지원하는 라이브러리는 버전을 낮춰 require로 동작하는 버전으로 수정했다.

### Spawn ENOENT 에러

- 윈도우에서 프로그램을 실행하니 아래와 같이 **`SPAWN ENOENT`** 에러를 보게 되었다.

  <img src="https://github.com/programmer-sjk/random-lunch-support/blob/main/images/spawn-enoent-error.png" width="650">

- [이 링크](https://github.com/vercel/pkg/issues/342)를 보고 runtime에 필요한 의존성 exe 파일을 다른 경로에 copy하려 했으나 이상하게 0btye의 껍데기 파일만 존재했다.
- 문제가 되었던건 `clipboardy` 라이브러리 한 개 뿐이라 `copy-paste` 모듈로 대체했다.
