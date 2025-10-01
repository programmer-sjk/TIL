# vscode 설정 & 단축키

## import시, single quote 하도록

- 필요한 이유
  - 더블 쿼트로 할 경우 lint에서 지속적인 빨간 줄로 표시
- 방법
  - `"typescript.preferences.quoteStyle": "single"`

## 사용하지 않는 import 제거

- F1을 누르고 setting.json 파일 오픈

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/ide/setting-json.png" width="500">

- 아래 설정 추가

  ```json
    "editor.codeActionsOnSave": {
      "source.fixAll": true,
      "source.fixAll.eslint": true
    },
  ```

## import시 상대 경로로 바꾸고 싶을 때

- import module 검색 -> shortest 에서 relative로 수정

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/ide/import-module.png" width="500">

## 파일의 마지막에 빈줄 추가

- 필요한 이유
  - lint에서 자꾸 표시
- 방법
  - Insert Final Newline 검색해서 활성화

## lint 관련

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/ide/8_lint_space.png" width="500">
- space 때문에 빨간 줄이 뜬다면 아래 설정 검색해서 enable
  - `trimTrailingWhitespace`

## 저장 시 prettier 적용

- 설정에서 default formatter를 검색해서 기본 값을 Prettier로 변경한다.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/ide/save_prettier.png" width="600">

- 설정에서 Prettier: Single Quote를 활성화 한다.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/ide/enable-single-quote.png" width="600">

## 생성자 cmd + 클릭으로 class로 이동하려는데 선택지가 뜨는 경우

- 예를 들어 ProductResponse 구현부로 들어가려는데 class 이름으로 이동할지 생성자로 이동할지 뜨는 경우가 있다.
- 이때는 사용자의 settings.json을 열고 아래 코드를 추가한다.
- `"editor.gotoLocation.multipleDefinitions": "goto"`

## 단축키

- 사용되지 않는 import 제거: shift + option + O
