# 타입으로 견고하게 다형성으로 유연하게

- [책 링크](https://product.kyobobook.co.kr/detail/S000210397750)

## 타입 검사 흝어보기

### 타입 검사의 정의와 필요성

- **`타입 검사는 무엇이며 왜 필요할까?`** 타입검사는 불편하지만 불편을 감수하면서도 사용할 가치가 있다.
- 우선 버그에 대한 이야기를 해야 한다. 버그는 개발자에게 가장 큰 적이면서도 피할 수 없는 존재다.
- 버그의 가장 흔한 원인은 타입 오류다. 타입은 프로그램에 존재하는 값을 분류한 것이다.
- **`의도한 타입의 값이 아닌 다른 타입이 들어와 실행된다면 기능이 중단되고 오류 메시지를 출력한다`**.
- 타입 오류로 인한 버그를 프로그램 실행 중에 찾기보다 **`자동으로 타입 오류를 판단해주는 타입 검사기를 사용한다`**.
  - 타입 검사기는 어디서 타입 오류가 발생하는지 알려준다.
  - 코드를 보고 타입 오류가 있다면 코드를 고치면 된다.
  - 코드를 보고 타입 오류가 없다면 타입 검사기가 그 사실을 올바르게 알 수 있도록 코드를 살짝 바꾸면 된다.
- **`타입 검사는 프로그램에서 버그를 자동으로 찾아 준다는 가치를 지닌다`**.
  - 버그는 아주 많고 사람의 힘만으로 모든 버그를 찾기란 매우 어렵다.
- 그래서 사람들은 타입 검사기를 사용한다.

### 정적 타입 언어

- 여러 언어가 타입 검사기를 제공한다. 자바, C, C++, 타입 스크립트, Go, 코틀린 등이 그 예다.
- 이런 언어들은 **`정적 타입 언어라고 부른다`**. 정적의 의미는 프로그램을 **`실행하기 전에를 뜻한다`**.
  - 정적 타입 언어는 프로그램을 실행하기 전 타입이 올바르게 사용되었는지 확인하는 언어다.
  - 프로그램이 타입 검사를 통과하면 실행 중 타입 오류가 일어나지 않는다는 보장과 함께 실행된다.
- 한편 **`타입 검사기를 제공하지 않는 언어도 있다`**. 자바스크립트, 파이썬, 루비, 리스프가 대표적이다.
  - 프로그램이 그냥 실행될 수 있지만, 실행 중에 타입 오류가 발생할 수도 있다.
  - 이런 언어들을 **`동적 타입 언어라고 부른다`**. 동적은 프로그램 **`실행 중에를 뜻한다`**.
  - 타입이 잘못 사용되었어도 그 사실을 실행중에야 파악할 수 있는 언어인 것이다.

### 타입 검사의 원리

- 타입 검사기가 작동하는 원리는 현실 세계에서 자동차 검사 과정과 비슷한다.
  - 자동차를 검사할 때는 작은 부품에서 큰 부품으로 가면서 검사를 한다.
  - 타이어부터 두 타이어를 연결한 차축과 같은 방향으로 검사를 한다.
- **`PrintInt(5 + 7)`** 검사 방식을 살펴보자.
  - 우선 기본 부품인 `5,7`이 정수인지 살펴본다.
  - 큰 부품인 `PrintInt와 5 + 7`을 살펴본다.
  - 함수의 인자 타입과 인자로 전달되는 `5 + 7` 타입은 정수이므로 문제 없이 통과된다.
- 어떤 언어들은 정수와 문자열의 덧셈을 허용한다
  - 그런 언어에서는 `(5 + "7")`이 타입 검사를 통과한다.
  - 따라서 핵심은 덧셈에는 특정 타입을 요구하며 타입 검사 과정에서 덧셈을 하는 부품이 그 요구를 만족하는지 확인한다.

#### 함수

- 함수의 타입은 매개변수 타입과 결과 타입으로 구성된다.
- 타입 검사기는 함수를 호출하는 쪽에서 전달하는 인자가 매개 변수 타입과 일치하는지 확인한다.
- 함수의 경우 함수 몸통(body)도 검사 대상이다.
  - 몸통 내부에서 타입 오류를 발생시키는지, return 코드가 실제 반환 타입과 일치하는지 확인한다.
- 어떤 값도 반환하지 않는 함수(void)의 경우 return에 아무 값도 주어지지 않았는지 확인한다.

### 타입 검사 결과의 활용

- 여태 알아본 **`정적 타입 언어의 장점은 아래 두 가지이다`**.
  - 타입 오류를 모두 찾을 수 있다는 것
  - 타입 검사기의 오류 메시지로 코드를 올바르게 고치기 쉽다는 점이다.
  - 이 외에도 **`타입 검사로 얻을 수 있는 이점은 두 가지 더 있다`**.
- 코드 편집기
  - 타입 검사기는 프로그램 각 부품의 타입을 알아내 자동완성 기능을 정확히 추측할 수 있도록 돕는다.
  - 예를 들어 문자열에 `contains` 메서드가 있고 `collect` 메서드가 없다고 가정하자.
  - 개발자가 변수.co 까지 작성하면 변수가 String일 경우 `collect`는 바로 선택지에서 제외할 수 있다.
- 프로그램 성능
  - 같은 프로그램을 작성해도 정적 타입 언어의 경우 타입 검사 덕분에 실행 중에 할일을 줄일 수 있다.
  - 예를 들어 자바스크립트에서 변수의 자료형이 number가 아닌 경우를 처리하는 if문을 예로 들 수 있다.
  - **`언어의 성능을 최대한 높이고 싶다면 타입 검사기의 도움을 받아야 한다`**.
    - 성능을 중요한 목표로 두는 `C, C++, Rust`가 모두 정적 타입 언어인 것은 다 이유가 있다.