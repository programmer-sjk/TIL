# 클린 코드

- [책 링크](https://product.kyobobook.co.kr/detail/S000001032980)

## 깨끗한 코드

- 80년대 후반 어떤 앱이 출시되었으나 제품 출시 주기가 점점 길어지기 시작했다. 이전 버전에 있었던 버그가 현재 버전에도 그대로 있고 프로그램이 죽는 횟수도 늘어났다. 결국 회사는 얼마 못가 망했다. 20년이 지나 그 회사의 초창기 직원을 만나 들었는데 짐작한대고 출시에 바빠 코드를 마구 짜고 결국은 감당이 불가능한 수준에 이르렀다. **`회사가 망한 원인은 바로 나쁜 코드 탓이었다`**. 우리 모두는 자신이 짠 쓰레기 코드를 쳐다보며 나중에 손보겠다고 생각한 경험이 있다. 우린 르블랑의 법칙을 법칙을 몰랐다. **`나중은 결코 오지 않는다`**.
- 나쁜 코드는 개발 속도를 떨어뜨린다. 프로젝트 초반에는 번개처럼 나가다가 1-2년 만에 굼벵이처럼 기어가는 팀도 많다. 코드를 고칠때마다 엉뚱한 곳에 문제가 생기고 시간이 지나면서 쓰레기 더미는 점점 높아지고 깊어진다. 생산성이 떨어지면 관리층은 나름대로 복구를 시도하기 위해 새로운 인력을 투입한다. 하지만 새 인력은 시스템 설계에 대한 조예가 깊지 않다. **`새 인력은 생산성을 향상시켜야 한다는 압박으로 나쁜 코드를 더 많이 양산한다`**. 결과적으로 생산성은 더욱 떨어져 거의 0이 된다.
- **`코드는 왜 나쁜 코드가 되었을까?`** 우리는 온갖 이유를 들이댄다. 요구사항이 변했고 일정이 촉박해 제대로 할 시간이 없었다고 한탄한다. 멍청한 관리자와 조급한 고객 때문이라고 말하지만 잘못은 전적으로 프로그래머에게 있다. 관리자가 일정과 요구사항을 밀어붙이는 이유는 그것이 그들의 책임이기 때문이다. **`좋은 코드를 사수하는 일은 바로 프로그래머들의 책임이다`**. 나쁜 코드의 위험을 이해하지 못하는 관리자 말을 그대로 따르는 행동은 전문가답지 못하다.
- 진짜 전문가는 기한을 맞추려면 나쁜 코드를 생산해야 한다고 생각하지 않는다. 나쁜 코드를 양산하면 기한을 맞추지 못한다. **`기한을 맞추는 유일한 방법은 언제나 코드를 최대한 깨끗하게 유지하는 습관이다`**. 그렇다면 깨끗한 코드를 어떻게 작성할까?

### 깨끗한 코드란?

- 프로그래머 수 만큼 정의도 다양할 것이다. 저자는 유명하고 노련한 프로그래머들에게 깨끗한 코드에 대한 의견을 물었다.
- **`비야네 스트롭스트룹 (C++ 창시자)`**
  - 나는 우아하고 효율적인 코드를 좋아한다. 논리가 간단해야 버그가 숨어들지 못한다. 의존성을 최대한 줄여야 유지보수가 쉬워진다. 오류는 명백한 전략에 의거해 철저히 처리한다. 성능을 최적으로 유지해야 사람들이 코드를 망치려는 유혹에 빠지지 않는다. 깨끗한 코드는 한 가지를 제대로 한다.
  - 비야네에 따르면 깨끗한 코드는 보기에 즐거운 코드다. 또한 유혹이란 단어를 사용했는데, **`나쁜 코드는 나쁜 코드를 유혹한다. 마지막으로 깨끗한 코드란 한 가지를 잘 한다고 단언한다`**. 수 많은 SW 설계 원칙이 이 간단한 교훈으로 귀결된다는 사실은 우연이 아니다. 나쁜 코드는 너무 많은 일을 하려다 목적이 흐려지지만, 깨끗한 코드는 한 가지에만 집중한다.
- **`그래디 부치`**
  - 깨끗한 코드는 단순하고 직접적이다. 또한 잘 쓴 문장처럼 읽힌다. 깨끗한 코드는 설계자의 의도를 숨기지 않는다. 오히려 명쾌한 추상화와 단순한 제어문으로 가득하다.
- **`데이브 토마스`**
  - 깨끗한 코드는 작성자가 아닌 사람도 읽기 쉽고 고치기 쉽다. 테스트 케이스가 존재하며 의미있는 이름이 붙는다.
- **`마이클 페더스`**
  - 깨끗한 코드는 누군가 주의깊게 짰다는 느낌을 준다. 고치려고 살펴봐도 딱히 손 댈 곳이 없어 제자리로 돌아온다. 그리고는 누군가 남겨준 코드, 누군가 주의깊게 짜놓은 작품에 감사를 느낀다.
- **`론 제프리스`**
  - 모든 테스트를 통과하고, 중복이 없고, 시스템 내 모든 설계 아이디어를 표현하고, 클래스와 메서드 함수 등을 최대한 줄인다. 중복 줄이기, 표현력 높이기, 추상화 고려하기. 내게는 이 세가지가 깨끗한 코드를 만드는 비결이다.
- **`워드 커닝햄`**
  - 코드를 읽으면서 짐작했던 기능을 각 루틴이 그대로 수행한다면 깨끗한 코드라 불러도 되겠다. 코드가 그 문제를 풀기 위한 언어처럼 보인다면 아름다운 코드라 불러도 되겠다.

### 깨끗한 코드의 중요성

- 코드를 읽는 시간 vs 코드를 작성하는 시간 비율은 10대 1을 훌쩍 넘는다. 새 코드를 작성하며 우리는 기존 코드를 읽는다. 비율이 이렇게 높으므로 읽기 쉬운 코드가 매우 중요하다. 이 논리에서 빠져나갈 방법은 없다. 주변 코드를 읽기 쉬우면 새 코드를 짜기도 쉽다. 주변 코드를 읽기가 어려우면 새 코드를 짜기도 어렵다. **`급하고 서둘러서 끝내려면, 읽기 쉽게 만들면 된다`**.
- 잘 짠 코드가 전부는 아니다. 시간이 지나도 언제나 깨끗하게 유지해야 한다. 미국 보이스카우트가 따르는 간단한 규칙이 우리에게도 유용하다. 캠프장은 처음 왔을 때 보다 더 깨끗하게 해놓고 떠나라. **`한꺼번에 많은 시간과 노력을 투자해 코드를 정리할 필요는 없다`**. 변수 이름 하나를 개선하고 조금 긴 함수를 분할하고 약간의 중복을 제거하면 충분하다. 시간이 지날수록 코드가 좋아지는 프로젝트에서 작업한다니! 지속적인 개선이야 말로 전문가 정신의 본질이 아니던가?

## 의미 있는 이름

- 이름만 잘 지어도 여러모로 편하다. 이 장에서는 이름을 잘 짓는 간단한 규칙을 소개하도록 하겠다.

### 의도를 분명히 밝혀라

- 의도가 분명한 이름이 정말로 중요하다는 사실을 거듭 강조한다. 아래 코드에서 d는 아무 의미도 드러나지 않는다. 특정하려는 값과 단위를 표현하는 이름이 중요하다.

  ```java
  int d; // 경과 시간

  // 좋은 예시
  int elapsedTimeInDays;
  int daysSinceCreation;
  int daysSinceModification;
  int fileAgeInDays;
  ```

### 그릇된 정보를 피하라

- 여러 계정을 그룹으로 묶을 때 실제 List가 아니라면 `accountList`라고 명명하지 않는다. 프로그래머에게 List는 특수한 의미다. 실제 List가 아니라면 프로그래머에게 그릇된 정보를 제공하는 셈이니 `accountGroup, accounts`라 명명한다.
- 유사한 개념은 유사한 표기법을 사용한다. 이것도 정보다. 일관성이 떨어지는 표기법은 그릇된 정보다.

### 의미 있게 구분하라

- 아래 코드는 저자의 의도가 전혀 드러나지 않는다. 인수 이름으로 `source와 destination`을 사용하면 코드 읽기가 훨씬 쉬워진다.

  ```java
    public static void copyChars(char a1[], char a2[]) {
      for (int i = 0; i < a1.length; i++) {
        a2[i] = a1[i];
      }
    }
  ```

- 고객 급여 이력을 찾으려면 아래에서 어느 클래스를 뒤져야 할까?

  ```java
    getActiveAccount();
    getActiveAccounts();
    getActiveAccountInfo();
  ```

- 명확한 관례가 없다면 moneyAmount는 money와 구분이 안된다. customerInfo는 customer와 theMessage는 message와 구분이 안 된다. 읽는 사람이 차이를 알도록 이름을 지어라.

### 클래스 이름

- 클래스 이름은 명사나 명사구가 적합하다. `Customer, WikiPage, Account` 등이 좋은 예다. `Manager, Processor, Data, Info` 등과 같은 추상적인 단어는 피하고 동사는 사용하지 않는다.

### 메서드 이름

- 메서드 이름은 동사나 동사구가 적합하다. `postPayment, deletePage, save` 등이 좋은 예다.
- 생성자를 중복 정의할 때는 정적 팩토리 메서드를 사용한다. 메서드는 인수를 설명하는 이름을 사용한다.

  ```java
    Complex fulCrumPoint = Complex.FromRealNumber(23.0); // 이 코드가 아래 코드보다 좋다.
    Complex fulCrumPoint = new Complex(23.0);
  ```

- 생성자 사용을 제한하려면 해당 생성자를 private로 선언한다.

### 한 개념에 한 단어를 사용하라

- 클래스마다 메서드 이름에 `fetch, get, retrieve`로 제각각 부르면 혼란스럽다. 마찬가지로 동일한 코드 기반에 `controller, manager, driver`를 섞어 쓰면 혼란스럽다. `DeviceManager와 ProtocolController`는 근본적으로 어떻게 다른가? 왜 둘다 Controller가 아닌가?

### 불 필요한 맥락을 없애라

- 고급 휘발유 충전소(Gas Station Deluxe)라는 어플리케이션을 만든다고 가정하자. 모든 클래스 이름을 GSD로 시작하겠다는 생각은 전혀 바람직하지 못하다.
- **`일반적으로 짧은 이름이 긴 이름보다 좋다. 단, 의미가 분명한 경우에 한해서다`**. 이름에 불필요한 맥락을 추가하지 않도록 주의한다.

## 함수

- 함수를 잘 만드는 방법을 소개한다.

### 작게 만들어라

- **`함수를 만드는 첫째 규칙은 작게다`**. 저자는 지난 40년 동안 온갖 크기로 함수를 구현해봤다. 지금까지 경험을 바탕으로 그리고 오랜 시행착오를 바탕으로 작은 함수가 좋다고 확신한다.
- **`다르게 말하면 if문 while문 등에 들어가는 블록은 한 줄이어야 한다는 의미이다`**. 대개 거기서 함수를 호출한다. 그러면 바깥을 감싸는 함수가 작아질 뿐더러 블록안에서 호출하는 함수 이름을 적절히 짓는다면 코드를 이해하기도 쉬워진다. 이 말은 중첩 구조가 생길만큼 함수가 커저서는 안 된다는 의미이다.

### 한 가지만 해라

- **`함수는 한 가지를 잘 해야 한다`**. 그 한 가지만을 해야 한다. 여기서 문제라면 그 한가지가 무엇인지 알기 어렵다는 점이다. **`지정된 함수 이름 아래에서 추상화 수준이 하나인 단계만 수행한다면 그 함수는 한 가지 작업만 한다`**. 어쨌거나 우리가 함수를 만드는 이유는 큰 개념을 다음 추상화 수준에서 여러 단계로 나눠 수행하기 위해서이다.
- 함수가 한 가지 작업만 하려면 함수 내 모든 문장의 추상화 수준이 동일해야 한다. 한 함수 내에 추상화 수준이 섞이면 코드를 읽는 사람이 헷갈린다. 특정 표현이 근본 개념인지 세부 사항인지 구분하기 어려운 탓이다.
- 코드는 위에서 아래로 이야기처럼 읽혀야 한다. 한 함수 다음에는 추상화 수준이 한 단계 낮은 함수가 온다. 즉 위에서 아래로 프로그램을 읽으면 함수 추상화 수준이 한 번에 한 단계씩 낮아진다. 저자는 이를 내려가기 규칙이라 부른다.

### Switch 문

- switch 문은 작게 만들기 어렵다. case 분기가 두개인 switch 문도 저자 취향에는 너무 길며, 한 가지 작업만 하는 switch 문도 만들기 어렵다. **`본질적으로 switch 문은 N가지를 처리한다`**. 불행하게도 switch 문을 완전히 피할 방법은 없지만 switch 문을 저차원 클래스에 숨기고 다형성을 이용할 수 있다.

  ```java
    public Money calculatePay(Employee e) {
      switch (e.type) {
        case COMMISSIONED:
          return calculateCommissionedPay(e);
        case HOURLY:
          return calculateHourlyPay(e);
        case SALARIED:
          return calculateSalariedPay(e);
        default:
          throw new InvalidEmployeeType(e.type);
      }
    }
  ```

- 위 함수에는 몇 가지 문제가 있다. **`먼저 함수가 길다. 또한 한 가지 작업만 수행하지 않는다. 코드를 변경할 이유가 여럿이기 때문에 SRP(Single Responsibility Principle)를 위반한다`**. 이 문제를 해결한 코드가 아래 코드다.

  ```java
    public abstract class Employee {
      public abstract boolean isPayDay();
      public abstract Money calculatePay();
    }

    public interface EmployeeFactory {
      public Employee makeEmployee(EmployRecord r);
    }

    public class EmployeeFactoryImpl implements EmployeeFactory {
      public Employee makeEmployee(EmployRecord r) {
        switch (r.type) {
          case COMMISSIONED:
            return new CommissionedEmployee();
          case HOURLY:
            return new HourlyEmployee();
          case SALARIED:
            return new SalariedEmployee();
          default:
            throw new InvalidEmployeeType(r.type);
        }
      }
    }
  ```

- 저자는 switch 문을 **`단 한번만 참아준다. 다형적 객체를 생성하는 코드 안에서다`**. 이렇게 상속 관계로 숨긴 후에는 절대로 다른 코드에 노출하지 않는다.

### 서술적인 이름을 사용하라

- 함수가 작고 단순할수록 서술적인 이름을 고르기 쉬워진다. 이름이 길어도 괜찮다. **`길고 서술적인 이름이 짧고 어려운 이름보다 좋다`**. 길고 서술적인 이름이 길고 서술적인 주석보다 좋다.

### 함수 인수

- 함수의 인자가 3개 이상은 가능한 피하는 편이 좋다.
- 플래그 인수는 추하다.왜냐면 함수가 여러가지를 처리한다고 대놓고 공표하기 때문이다. render 라는 함수에 플래그 변수를 넘기기 보다는 `renderForSuite()`와 `renderForSingleTest()`라는 함수로 나눠야 마땅하다.

### 부수 효과를 일으키지 마라

- 부수 효과는 거짓말이다. 함수에서 한 가지를 하겠다고 약속해놓고선 남몰래 다른 짓을 하는 것이다.

  ```java
    public boolean checkPassword(String userName, String password) {
      User user = userRepository.findByName(userName);
      if (user) {
        String codedPhrase = user.getPhraseEncodedByPassword();
        String phrase = cryptographer.decrypt(codedPhrase, password);
        if ("Valid Password".equlas(phrase)) {
          Session.initialize();
          return true;
        }
      }

      return false;
    }
  ```

- 위에서 **`부수 효과는 Session.initialize() 호출이다`**. 함수의 이름은 암호를 확인하는데 이름만 봐서는 세션을 초기화하는 사실이 드러나지 않는다. 만약 꼭 이런 코드가 필요하다면 함수 이름에 분명히 명시한다. `checkPasswordAndInitializeSession` 이라는 이름이 훨씬 좋다. 물론 함수가 한 가지만 한다는 규칙을 위반하긴 하지만.

### 명령과 조회를 분리하라

- 함수는 객체 상태를 변경하거나 조회하거나 둘 중 하나만 해야 한다. 다음 함수를 살펴보자.

  ```java
    public boolean set(String attribute, String value);
  ```

- 이 함수는 값을 value로 설정하고 성공하면 true를 반환한다. 따라서 다음과 같은 괴상한 코드가 나온다.

  ```java
    if (set("username", "unclebob")) ...
  ```

- 이 코드를 읽는 입장에선 `username 속성이 unclebob으로 설정되어 있다면`으로 읽힌다. username이 unclebob으로 설정하는데 성공하면 이라고 읽히지 않는다. 이에 대한 **`해결책은 명령과 조회를 분리해 혼란을 애초에 뿌리뽑는 방법이다`**.

### 오류 코드보다 예외를 사용하라

- 명령 함수에서 오류 코드를 반환하는 방식은 명령/조회 분리 규칙을 미묘하게 위반한다. 오류 코드 대신 예외를 사용하면 오류를 받아 비교하는 처리가 없어져 코드가 깔끔해진다.

  ```java
    // 오류 코드를 처리하는 패턴
    if (deletePage(page) === E_OK) {
      if (registry.deleteReference(page.name) === E_OK) {
        if (configKeys.deleteKey(page.name.makeKey()) === E_OK) {
          ...
        }
      }
    }

    // 예외를 뱉는 패턴
    try {
      deletePage(page);
      registry.deleteReference(page.name);
      configKeys.deleteKey(page.name.makeKey())
    } catch (Exception e) {
      logger.log(e.getMessage());
    }
  ```

### 함수를 어떻게 짜죠?

- SW를 개발하는 행위는 여느 글짓기와 비슷하다. 초안을 서투르게 작성하고 문장을 고치고 문단을 정리한다. **`저자가 함수를 만들 때도 마찬가지이다`**. 처음에는 길고 복잡하다. 기능을 완성하면 코드를 빠짐없이 테스트 하는 단위 테스트 케이스를 만든다. 그런 다음 저자는 코드를 다듬고 함수를 만들고 이름을 바꾸고 중복을 제거한다. 때로는 전체 클래스를 쪼개기도 한다. **`이 와중에도 코드는 항상 단위 테스트를 통과한다`**.

## 주석

- 잘 달린 주석은 그 어떤 정보보다 유용하지만 경솔하고 근거 없는 주석은 코드를 이해하기 어렵게 만든다. 오래되고 조잡한 주석은 거짓된 정보를 전달해 해악을 미친다. 우리는 코드로 의도를 표현하지 못해, 실패를 만회하기 위해 주석을 사용한다. 주석은 언제나 실패를 의미한다. **`저자가 주석을 무시하는 이유는 자주 거짓말을 하기 때문이다. 주석은 오래될수록 코드에서 멀어진다`**.
- **`코드에 주석을 추가하는 일반적인 이유는 코드 품질이 나쁘기 때문이다`**. 코드를 짜고 보니 엉망이고 알아보기 어렵다. 저자는 이때 주석을 달기 보다는 코드를 정리하라고 말하고 싶다. 표현력이 풍부하고 깔끔하며 주석이 거의 없는 코드가 복잡하고 어수선하며 주석이 많이 달린 코드보다 훨씬 좋다.
- **`어떤 주석은 필요하거나 유익하다`**. 글자 값을 한다고 생각하는 주석 몇 가지를 소개한다.
  - 법적인 주석: 때로는 법적인 이유로 특정 주석을 넣으라고 명시할 때가 있다. 소스 첫 머리에 들어가는 저작권과 소유권 정보는 타당하다.
  - 중요성을 강조하는 주석: 자칫 대수롭지 않다고 여겨질 뭔가의 중요성을 강조하기 위해 사용한다.
  - TODO 주석: 앞으로 할 일을 TODO 주석으로 남겨두면 편하다.
  - 공개 API에서 Javadocs: 설명이 잘 된 공개 API는 유용하고 만족스럽다. 공개 API를 구현한다면 반드시 훌륭한 Javadocs를 작성한다.
- **`대다수의 주석은 나쁜 주석에 속한다`**. 몇 가지 사례를 살펴보자.
  - 같은 이야기를 중복하는 주석
  - 의무적으로 다는 주석: 모든 함수에 Javadocs를 달거나 모든 변수에 주석을 달아야 하는 규칙은 어리석기 그지없다.
  - 있으나 마나 한 주석: 생성자에 기본 생성자라고 주석을 다는 경우로 정보를 전달하지 못하는 주석이다.
  - 너무 많은 정보: 주석에 흥미없는 역사나 정보를 장황하게 늘어놓지 않는다.
  - 비공개 코드에서 Javadocs: 시스템내부에 속한 클래스나 함수에 Javadocs를 생성할 필요는 없다.

## 형식 맞추기

- 프로그래머라면 형식을 깔끔하게 맞춰 코드를 짜야 한다. 코드 형식을 맞추기 위한 간단한 규칙을 정하고 그 규칙을 착실히 따라야 한다. 필요하다면 규칙을 자동으로 적용하는 도구를 활용한다.
- 개념은 빈 행으로 분리하자. 모든 코드는 왼쪽에서 오른쪽으로 그리고 위에서 아래로 읽힌다. 일련의 행 묶음은 완결된 생각 하나를 표현한다. **`생각 사이는 빈 행을 넣어 분리해야 마땅하다. 빈 행은 새로운 개념을 시작한다는 시각적 단서다`**.

  ```java
    // 빈 행 적용
    package fitnesses.wikitext.widgets;

    import java.util.regex.*;

    public class BoldWidget extends ParentWidget {
      public static final String REGEXP = "'''.+?'''";

      public BoldWidget(ParentWidget parent, String text) {
        super(parent);
        Matcher match = pattern.matcher(text);
        match.find();
        addChildWidget(match.group(1));
      }
    }

    // 빈행이 없는 경우
    package fitnesses.wikitext.widgets;
    import java.util.regex.*;
    public class BoldWidget extends ParentWidget {
      public static final String REGEXP = "'''.+?'''";
      public BoldWidget(ParentWidget parent, String text) {
        super(parent);
        Matcher match = pattern.matcher(text);
        match.find();
        addChildWidget(match.group(1));
      }
    }
  ```

- 위에서 두번째 코드는 빈행을 빼버린 코드로 코드 가독성이 현저하게 떨어진다.
- 줄 바꿈이 개념을 분리한다면 **`세로 밀집도는 연관성을 의미한다. 즉 서로 밀접한 코드 행은 세로로 가까이 놓아야 한다는 뜻이다`**. 변수는 실제로 사용하는 위치에 최대한 가까이 선언한다. 한 함수가 다른 함수를 호출한다면 두 함수는 세로로 가까이 배치한다. 또한 가능하다면 호출되는 함수를 뒤에 배치한다. 그러면 프로그램이 자연스럽게 읽힌다.
- **`프로그래머라면 각자 선호하는 규칙이 있지만 팀에 속한다면 자신이 선호해야 할 규칙은 팀 규칙이다`**. 팀은 한 가지 규칙에 합의해야 하고 모든 팀원은 그 규칙을 따라야 한다. 그래야 SW가 일관적인 스타일을 보인다. 개개인이 따로국밥처럼 맘대로 짜는 코드는 피해야 한다. 좋은 SW는 읽기 쉬운 문서로 이뤄진다는 사실을 기억해야 한다. 스타일은 일관적이고 한 소스에서 봤던 형식이 다른 소스 파일에도 쓰이리라느 신뢰감을 독자에게 줘야 한다.

## 객체와 자료구조

- 변수를 private로 정의하는 이유는 남들이 변수에 의존하지 않게 만들고 싶어서다. 그렇다면 어째서 수 많은 개발자가 get, set 함수를 당연하게 공개해 비공개 변수를 외부에 노출할까?
- 아래 두 클래스는 모두 2차원점을 표현한다. 한 클래스는 구현을 외부로 노출하고 다른 클래스는 구현을 완전히 숨긴다.

  ```java
    public class Point {
      public double x;
      public double y;
    }

    public interface Point {
      double getX();
      double getY();
      void setCartesian(double x, double y);
      double getR();
    }
  ```

- 두 번째 예시는 구현을 숨김에도 불구하고 자료 구조를 명백하게 표현한다. 또 클래스 메서드가 접근 정책을 강제한다. 좌표를 읽을 때는 개별적으로 읽어야 하지만 좌표를 설정할 때는 두 값을 한꺼번에 설정해야 한다. 반면 첫번째 코드는 개별적으로 좌표 값을 읽고 설정하게 한다. 변수를 private로 선언하더라도 각 값마다 get, set 함수가 공개된다면 구현을 외부로 노출하는 셈이다.
- 변수 사이에 함수라는 계층을 넣는다고 구현이 저절로 감춰지지는 않는다. 구현을 감추려면 추상화가 필요하다. 그저 get, set 함수로 변수를 다룬다고 클래스가 되지는 않는다. 그보다는 추상 인터페이스를 제공해 사용자가 구현을 모른채 자료의 핵심을 조작할 수 있어야 진정한 클래스다. 자료를 세세하게 공개하기 보다는 추상적인 개념으로 표현하는 편이 좋다.

### 자료/객체 비대칭

- 위에서 설명한 두 Point는 객체와 자료구조를 보여준다. 객체는 추상화 뒤로 자료를 숨긴채 자료를 다루는 함수만 공개한다. 자료구조는 자료를 그대로 제공하며 별다른 함수를 제공하지 않는다. 아래 코드는 절차적인 도형 클래스다. 각 도형 클래스는 아무 메서드도 제공하지 않고 Geometry 클래스가 세 가지 도형 클래스를 다룬다.

  ```java
    public class Square {
      public Point topLeft;
      public double side;
    }

    public class Rectangle {
      public Point topLeft;
      public double height;
      public double width;
    }

    public class Circle {
      public Point center;
      public double radius;
    }

    public class Geometry {
      public final double PI = 3.14;

      public double area(Object shape) {
        if (shape instanceof Square) {
          Square s = (Square)shape;
          return s.side * s.side;
        } else if (shape instanceof Rectangle) {
          Rectangle r = (Rectangle)shape;
          return r.height * r.width;
        } else if (shape instanceof Circle) {
          Circle c = (Circle)shape;
          return PI * c.radius * c.radius;
        }
      }
    }
  ```

- 위 코드는 나름의 장단점이 있다. 만약 Geometry 클래스에 둘레 길이를 구하는 함수를 추가하고 싶다면 도형 클래스는 아무런 영향을 받지 않는다. 반대로 새 도형을 추가하고 싶다면 Geometry 클래스에 속한 함수를 고쳐야 한다.
- 이번에는 객체 지향적인 도형 클래스를 살펴보자. 여기서 area는 다형 메서드이다.

  ```java
    public class Square implements Shape {
      public Point topLeft;
      public double side;

      public double area() {
        return side * side;
      }
    }

    public class Rectangle implements Shape {
      public Point topLeft;
      public double height;
      public double width;

      public double area() {
        return height * width;
      }
    }

    public class Circle implements Shape {
      public Point center;
      public double radius;

      public double area() {
        return PI * radius * radius;
      }
    }
  ```

- 절차적인 코드는 기존 자료구조를 변경하지 않으면서 새 함수를 추가하기 쉽다. 반면 객체지향 코드는 기존 함수를 변경하지 않으면서 새 클래스를 추가하기 쉽다. 반대로 절차적인 코드는 새로운 자료구조를 추가하기 어렵다. 모든 함수를 고쳐야 한다. 객체지향 코드는 새로운 함수를 추가하기 어렵다. 모든 클래스를 고쳐야 한다.
- 정리하면 객체지향 코드에서 어려운 변경이 절차적인 코드에선 쉽고, 절차적인 코드에서 어려운 변경이 객체지향 코드에선 쉽다. 분별 있는 프로그래머라면 모든 것이 객체라는 생각이 미신임을 안다. 때로는 단순한 자료구조와 절차적인 코드가 적합한 상황도 있다.

### 디미터 법칙

- 디미터 법칙은 조작하는 객체의 속사정을 몰라야 한다는 법칙이다. 구체적으로 말하면 클래스 C의 메서드 f는 다음과 같은 객체의 메서드만 호출해야 한다고 주장한다. 다른 말로 낯선 사람은 경계하고 친구랑만 놀라는 법칙이다.
  - 클래스 C
  - f가 생성한 객체
  - f 인수로 넘어온 개체
  - C 인스턴스 변수에 저장된 객체
- 다음 코드는 디미터 법칙을 어기는 듯이 보인다.

  ```java
    String outputDir = ctxt.getOptions().getScratchDir().getAbsolutePath();
  ```

- 위와 같은 코드를 기차 충돌이라 부른다. 일반적으로 조잡하다 여겨지는 방식이므로 피하는 편이 좋으며 아래와 같이 나누는 편이 좋다.

  ```java
    Options opts = ctxt.getOptions();
    File scratchDir = opts.getScratchDir();
    String outputDir = scratchDir.getAbsolutePath();
  ```

- 위 코드가 디미터 법칙을 위반하는지 여부는 ctxt, Options, ScrachDir이 객체인지 자료구조인지에 달려있다. 객체라면 내부 구조를 숨겨야 하므로 디미터 법칙을 위반한다. 반면 자료구조라면 당연히 내부 구조를 노출하므로 디미터 법칙이 적용되지 않는다.
- 만약 객체라면 어떻게 해야 할까? ctxt가 객체라면 뭔가를 하라고 말해야지 속을 드러내라고 말하면 안 된다. outputDir이 어디에 쓰이는지 찾아보니 실제로 임시 파일을 생성하기 위한 목적이였다. 그렇다면 ctxt 객체에 임시 파일을 생성하라고 시키면 어떨까?

  ```java
    BufferedOutputStream bos = ctxt.createScratchFileStream(classFileName);
  ```

- ctxt는 내부 구조를 드러내지 않으며 모듈에서 여러 객체를 탐색할 필요가 없다. 따라서 디미터 법칙을 위반하지 않는다.
