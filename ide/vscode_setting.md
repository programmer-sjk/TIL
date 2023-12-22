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

## 단축키

- 사용되지 않는 import 제거: shift + option + O
