# Jetbrains IDE 3년 쓰다가 VScode를 쓰고 느끼는 점

![logo](/images/ide/1_logo.png)

Jetbrains가 제공하는 `webstorm, IntelliJ`를 3년 가까이 쓰다가 다시 `Visual Studio Code`로 넘어왔다. 그 동안 몰랐던 Jetbrains IDE에 감사함을 느끼며 VScode를 1주일 사용한 관점에서 어떤 점을 느꼈는지 정리하도록 하겠다.

내가 알고 있는게 전부가 아니며 추후 발견하거나 잘못 알았던 사실들은 지속해서 업데이트 할 생각이다. 이번 포스팅은 **매우 주관적인 견해**임을 명시하겠다.

## VSCode의 단점들

### Multi 파일 경로 리팩토링시 import 문제

여러 파일을 드래그해서 다른 디렉토리로 옮긴다고 가정해보자. 각 파일마다 import 하는 경로가 있을텐데, JetBrains IDE인 `Webstorm, IntelliJ` 을 쓰면 정상적으로 경로들이 리팩토링 된다.

하지만 VSCode는 그 중 **하나의 파일만 리팩토링**이 된다. 즉 여러 파일을 하나의 디렉토리로 옮기기 위해선 파일을 하나씩 이동 시키면 문제없이(?) 리팩토링을 할 수 있다.

이 문제를 나만 겪느냐? 이미 관련된 글들이 있었고 이를 [해결하기 위한 PR](https://github.com/microsoft/vscode/pull/105111)도 존재한다.

문제는 2020년에 만들어진 이 **PR이 아직도 open 상태**라는 것이다. 2022년에 이 기능을 담당하던 개발자가 좀 더 분석이 필요하다고 남긴 comment도 인상적이다.

![pr](/images/ide/2_pr.png)

### 리팩토링 기능

Jetbrains IDE를 사용하는 정말 큰 요소라고 생각한다. VScode에서 아래와 같은 코드를 만났다고 가정해보자.

```javascript
function doSomething(data) {
  if (typeof data == "number" || typeof data == "boolean") {
    return false;
  }
  .
  .
}
```

**if 조건문에서 구현이 드러나므로** 이를 private한 함수로 추상화해서 무엇을 하는지 설명하도록 코드를 리팩토링 하기 위해 아래와 같이 코드를 작성한다. (최종 코드가 아니고 리팩토링 하기위한 중간 단계의 코드이다.)

```javascript
function doSomething(data) {
  if (typeof data == "number" || typeof data == "boolean") {
    return false;
  }

  if (isValidType(data)) {
    return false;
  }
  .
  .
}
```

`isValidType` 함수는 아직 존재하지 않는 함수로, `Webstorm을` 쓴다면 커서만 위에 두고 단축키를 써서 외부에 함수로 정의할 수 있다. 반대로 VScode는 커서를 두고 리팩토링 단축키인 `command + .` 버튼을 클릭하면 아래와 같이 아무런 행동을 할 수 없다.
![refactor1](/images/ide/3_refactoring.png)

만약 마우스 드래그나 키보드로 `isValidType(data)` 부분을 지정하고 `command + .` 버튼을 누르면 리팩토링 기능을 쓸 수 있다.

![refactor2](/images/ide/4_refactoring.png)

정리하자면 리팩토링 할 때 **영역을 잘 잡아야(드래그 해야) 리팩토링 기능이 제공**된다. 이건 하나의 예시일 뿐 몇 가지 더 해보니 화가 난다.

### 코드 copy시, 자동 import 경로 지원안됨

아래와 같은 코드가 있을 때 `fs.readFileSync('bla')` 부분을 내가 만들고 있는 파일에 paste(붙여넣기) 한다고 가정해보자.

```typescript
import fs from 'fs';
import http from 'http';

http.createServer((req, res) => {
    const data = fs.readFileSync('bla');
    res.write(data);
    res.end()
})
```

VScode는 해당 부분을 `copy & paste` 한 뒤에 fs의 import 경로를 수동으로 추가해야 한다. ```import fs from 'fs';```

반면에 Jetbrains IDE는 코드를 붙여넣기 할 때 자동으로 import가 추가된다. 실무에서 한 번도 import가 잘못돼서 문제가 된 적은 없으며, 만약 동일한 이름을 가질 경우 선택할 수 있는 UI가 지원된다.

이 기능은 현재 [기능개발 요청상태](https://github.com/microsoft/TypeScript/issues/50187)이며 별게 아니라고 생각할 수 있지만(나도 쓸때는 몰랐지) 정작 VScode에서 지원이 안 되니 너~무 불편하다.

### JetBrains IDE 내장 기능 vs VScode Extension

JetBrains IDE는 **대부분의 기능을 내장**하고 설정으로 이뤄지는 반면에 VSCode는 많은 기능들이 `Extension` 설치를 통해 제공한다. `Webstorm(Jetbrains node IDE)`에서 기본으로 제공되는 기능들을 VSCode에서는 어떻게 할 수 있는지 검색하고 그 기능을 지원하는 Extension을 설치하는 과정이 나에게는 꽤 번거로웠다.

![extension](/images/ide/5_extension.png)

### 호출되지 않는 함수

아래는 `controller와 service` 코드인데 **getHi** 함수는 어디에서도 호출되지 않는다. JetBrains 계열의 IDE는 이 경우 함수가 **회색빛으로 표현되어 호출되지 않는 함수인 걸 인지**할 수 있는데 VSCode는 직접 호출되고 있는지 클릭해서 확인해야 한다.

![display](/images/ide/6_display.png)

## VScode의 장점 없을까?

우선 JetBrains IDE 보다는 훨씬 가볍다. 그리고 공짜다. 그리고.. 음..

## 결론

VScode 1주일 썼는데 이정도가 나왔다.
Jetbrains IDE와 VScode 하나를 선택할 수 있다면 무조건 Jetbrains IDE 쓰자.
