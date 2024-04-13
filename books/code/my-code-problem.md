# 내 코드가 그렇게 이상한가요?

- [책 링크](https://product.kyobobook.co.kr/detail/S000202521361)

## 잘못된 구조의 문제 깨닫기

- 좋은 구조로 개선하기 위해서는 **`일단 나쁜 구조의 폐해를 인지해야 한다`**.
- 그 후 개선할 수 있는 좋은 구조를 배우면 나쁜 구조와 좋은 구조의 차이를 파악해 설계할 수 있다.

### 의미를 알 수 없는 이름

- 기술 중심이나 일련 번호`(001 -> 002 -> ...)`로 네이밍을 하면 코드에서 어떠한 의도도 읽어낼 수 없다.
- **`의도와 목적을 드러내는 이름을 사용해야 한다`**.

### 이해하기 어렵게 만드는 조건 분기 중첩

- if 조건문이 중첩될 수록 코드의 가독성이 나빠진다.

### 수 많은 악마를 만들어 내는 데이터 클래스

- 데이터만 가지고 있는 클래스를 데이터 클래스라고 부른다.
- 문제는 데이터를 계산하는 로직이 다른 클래스에 퍼져 있다면 아래와 같은 문제가 발생한다.
  - 수정시, 수 많은 클래스를 수정해야 하고, 일부를 놓치면 버그가 발생한다.
  - 관련 코드들이 중복되어 있고 가독성이 저하된다.

### 악마 퇴치의 기본

- 악마들을 물리치기 위해 나쁜 구조의 폐해를 인지해야 한다.
  - **`나쁜 폐해를 인지하면`** 어떻게든 대처해야겠다 라고 생각하게 된다.

## 설계 첫 걸음

### 의도를 분명히 전달할 수 있는 이름 설계하기

- 위 코드보다는 아래 코드가 낫다

  ```java
    // bad
    int d = 0;
    d = p1 + p2;
    d = d - ((d1 +d2) / 2);
    if (d < 0) {
      d = 0;
    }

    // better
    int damageAmount = 0;
    damageAmount = playerArmPower + playerWeaponPower;  // A
    damageAmount = damageAmount - ((enemyBodyDefence + eenemyArmorDefence) / 2);
    if (damageAmount < 0) {
      damageAmount = 0;
    }
  ```

### 목적별로 변수를 따로 만들어 사용하기

- 위 코드는 이해하기 쉬워졌지만 `damageAmount` 변수에 **`재할당이 반복되고 있다`**.
  - 재할당은 **`변수의 용도가 바뀌는 문제를 일으키기 쉽고`**, 코드를 읽는 사람을 혼란스럽게 만든다.
- 주석 A 부분은 실제 플레이어 공격력의 총량이므로 아래와 같이 수정하면 어떤 값을 계산하는데 어떤 값을 사용하는지 관계를 파악하기 쉽다.

  ```java
    // best
    int totalPlayerAttackPower = playerArmPower + playerWeaponPower;
    int totalEnemyDefence = (enemyBodyDefence + eenemyArmorDefence) / 2);

    int damageAmount = totalPlayerAttackPower - totalEnemyDefence;
    if (damageAmount < 0) {
      damageAmount = 0;
    }
  ```

### 관련된 데이터와 로직을 클래스에 모으기

- 어떤 **`변수와 변수를 조작하는 로직이`** 이곳저곳 만들어지면 버그가 발생할 것이다.
- 이러한 문제를 해결하기 위해 데이터와 메서드를 모아놓은 클래스가 존재한다.

  ```java
    // 게임의 HP를 나타내는 HitPoint
    class HitPoint {
      private static final int MIN = 0;
      private static final int MAX = 999;
      final int value;

      HitPoint(final int value) {
        if (value < MIN) throw new IllegalArgumentException("불라불라");
        if (value > MAX) throw new IllegalArgumentException("불라불라");

        this.value = value;
      }

      HitPoint damage(final int damageAmount) {
        // 데미지를 받은 경우 계산 로직
      }

      HitPoint recover(final int recoveryAmount) {
        // 데미지 회복 계산 로직
      }
    }
  ```

- HitPoint 클래스는 HP와 밀접한 변수와 로직을 담고 있어서 이곳저곳에서 수정하지 않아도 된다.

## 클래스 설계: 모든 것과 연결되는 설계 기반

### 클래스 단위로 잘 동작하도록 설계하기

- 우리가 사용하는 전자 제품은 그 자체로 잘 동작하게 설계되어 있다.
- 클래스 설계도 마찬가지로 **`클래스는 클래스 하나로도 잘 동작할 수 있도록 설계해야 한다`**.

#### 클래스의 구성요소

- 잘 만들어진 클래스는 **`다음 두 가지로 구성된다`**.
  - 인스턴스 변수
  - 인스턴스 변수에 **`잘못된 값이 할당되지 않게 막고, 정상적으로 조작하는 메서드`**
- 데이터 클래스는 변수의 초기화가 조작하는 로직이 다른 클래스에 있으므로 혼자서는 아무것도 할 수 없는 클래스가 된다.

#### 모든 클래스가 갖추어야 하는 자기 방어 임무

- 다른 클래스를 사용해서 초기화와 유효성 검사를 해야 하는 클래스는 그 자체로 안전할 수 없는 클래스다.
- **`클래스 스스로 자기 방어 임무를 수행할 수 있어야`** 품질을 높이는데 도움이 된다.

### 성숙한 클래스로 성장시키는 설계 기법

- 아래의 `Money` 데이터 클래스를 성장시켜 보자.

  ```java
    class Money {
      int amount;
      Currency currency;
    }
  ```

#### 생성자로 확실하게 정상적인 값 설정하기

- 유효성 검사와 초기화 로직을 생성자 내부에 구현하자

  ```java
    Money(int amount, Currency currency) {
      if (amount < 0) throw new IllegalArgumentException("불라불라");
      if (currency == null) throw new IllegalArgumentException("불라불라");

      this.amount = amount;
      this.currency = currency;
    }
  ```

- 생성자에 유효성 검사 로직을 작성해두면 항상 안전하고 정상적인 객체만 존재하게 된다.

#### 계산 로직도 데이터를 가진 쪽에 구현하기

- 객체의 데이터를 조작하는 메서드는 객체 내부에 둔다.

  ```java
    class Money {
      void add(int other) {
        amount += other;
      }
    }
  ```

#### 불변 변수로 만들어서 예상하지 못한 동작 막기

- **`변수의 값이 바뀌면`**, 값이 언제 변경되는지 현재 값은 무엇인지 계속 신경써야 한다.
- 또 요구사항이 바뀌면서 예상치 못한 부수 효과가 쉽게 발생할 수 있다.
- 이를 위해 **`변수를 불변으로 만든다`**. 값을 한 번 할당하면 다시 바꿀 수 없는 변수를 불변 변수라고 부른다.
- 인스턴스 변수에 `final` 키워드를 붙이면 한 번만 할당할 수 있다.

  ```java
    class Money {
      final int amount;
      final Currency currency;

      Money(int amount, Currency currency) {
        this.amount = amount;
        this.currency = currency;
      }
    }
  ```

#### 변경하고 싶다면 새로운 인스턴스 만들기

- 인스턴스 변수의 값을 변경하는게 아니라, 변경된 값을 가진 인스턴스를 만들어서 사용하면 된다.

  ```java
    class Money {
      Mony add(int other) {
        int added = amount + other;
        return new Money(added, currency);
      }
    }
  ```

#### 메서드 매개변수와 지역 변수도 불변으로 만들기

- 메서드의 매개 변수도 값이 바뀔 수 있는데, 값이 중간에 바뀌면 값의 변화를 추적하기 힘들어진다.
- **`매개 변수와 지역 변수도`** final 키워드를 붙여서 변경될 수 없도록 하자

  ```java
    class Money {
      Mony add(final int other) {
        final int added = amount + other;
        return new Money(added, currency);
      }
    }
  ```

#### 엉뚱한 값을 전달하지 않도록 하기

- 부수 효과로는 **`잘못된 값의 전달도 포함이 된다`**.
- 엉뚱한 값이 전달되지 않도록 하려면 `Money` 자료형만 받도록 수정한다.

  ```java
    class Money {
      Mony add(final Mony other) {
        final int added = amount + other.amount;
        return new Money(added, currency);
      }
    }
  ```

#### 의미 없는 메서드 추가하지 않기

- 시스템 사양에 필요하지 않은 메서드를 선의로 추가했다면, 이후 누군가 사용시 버그가 될 수 있다.
- 시스템 사양에 꼭 필요한 메서드만 정의하자.

### 악마 퇴치 효과 검토하기

- 위 코드를 전부 모은 `Money` 클래스를 살펴보자

  ```java
    class Money {
      final int amount;
      final Currency currency;

      Money(int amount, Currency currency) {
        if (amount < 0) throw new IllegalArgumentException("불라불라");
        if (currency == null) throw new IllegalArgumentException("불라불라");

        this.amount = amount;
        this.currency = currency;
      }

      Mony add(final Mony other) {
        final int added = amount + other.amount;
        return new Money(added, currency);
      }
    }
  ```

- 검증해보면 기존 **`악마가 잘 퇴치된 것을 볼 수 있다`**.
  - 필요한 로직이 클래스 내부에 모여 중복 코드가 작성될 일이 줄어듬
  - 중복 코드가 없으므로 수정 시, 누락이 발생할 일이 줄어듬
  - 필요한 로직이 모두 클래스에 있어 가독성이 높아짐
  - 생성자에서 변수의 값을 확정하므로 초기화되지 않은 상태가 있을 수 없음
  - 불변 변수를 사용하여 부수 효과로부터 안전함

### 프로그램 구조에 도움을 주는 디자인 패턴

- 응집도가 높은 구조, 방어적인 프로그램의 구조를 개선하는 설계 방법을 **`디자인 패턴이라 부른다`**.
- 위 `Money` 클래스는 사실 **`완전 생성자와 값 객체라는 두 가지 디자인 패턴을`** 적용한 것이다.

#### 완전 생성자

- **`완전 생성자는`** 잘못된 상태로부터 클래스를 보호하기 위한 디자인 패턴이다.
- **`생성자 내부에 잘못된 값이 들어오지 않도록 검사하고`**, 인스턴스 변수를 모두 초기화해야 객체를 생성할 수 있게 한다.
  - 이렇게 설계하면 값이 모두 정상인 완전한 객체만 만들어지게 된다.

#### 값 객체 (Value Object)

- **`값 객체란`** 값을 클래스로 나타내는 디자인 패턴이다.
  - 애플리케이션에서 사용하는 `금액, 날짜, 주문 수, 전화번호` 등 다양한 값을 값 객체로 만들 수 있다.
- 값 객체의 장점은 **`각각의 값과 로직을 응집도가 높은 구조로 만들 수 있다`**.
  - 예를 들어 금액을 단순히 int 자료형으로 사용할 경우 금액 계산 로직이 여기저기 분산될 수 있다.
- 애플리케이션 내부에서 다루는 값과 개념들은 모두 값 객체로 만들 수 있다.
  - ex) `세외 제금 금액, 상품명, 전화번호, 배송지, 연령, 성별, 몸무게, 공격력, 아이템 가격, 아이템 이름 등`
- **`값 객체 + 완전 생성자는`** 객체지향 설계에서 폭 넓게 사용되는 기법이다.
- 언어와 상관없이 중요한 것은 데이터와 로직을 한 곳에 모아 응집도를 높이는 것, 필요한 조작만 외부에 공개해 캡슐화 하는 것이다.

## 불변 활용하기: 안정적으로 동작하게 만들기

- **`상태를 변경할 수 없는 것을`** 불변이라고 부른다.
- 가능한 상태가 변경되지 않도록 설계해야 하는데, 불변은 최근 프로그래밍 스타일의 표준 트렌드이다.

### 재할당

- **`재할당은 변수의 의미가 바뀌며, 읽는 사람은 헷갈릴 수 밖에 없다`**.
- 변수 하나를 재활용하지 않고, 새로운 변수를 만들어 사용하면 재할당을 피할 수 있다.
- 재할당을 막는 방법은 변수와 매개 변수에 `final` 키워드를 붙여 불변으로 만드는 것이다.

### 가변으로 인해 발생하는 의도하지 않은 영향

- 공격력을 나타내는 `AttackPower`와 무기를 나타내는 `Weapon` 클래스가 있다.

  ```java
    class AttackPower {
      int value; // final이 없으므로 가변
      AttackPower(int value) {
        this.value = value;
      }
    }

    class Weapon {
      final AttackPower attackPower;
      Weapon(AttackPower attackPower) {
        this.attackPower = attackPower;
      }
    }
  ```

- 처음 코드를 짤 때는 모든 무기의 공격력이 고정이라서 아래와 같이 코드를 작성했다.

  ```java
    AttackPower attackPower = new AttackPower(20);
    Weapon weaponA = new Weapon(attackPower);
    Weapon weaponB = new Weapon(attackPower);
  ```

- 이후 무기 각각의 공격력을 강화할 수 있도록 조건을 변경하자란 이야기가 나왔다.
- 그런데 어떤 무기의 공격력을 변경하면 다른 무기의 공격력도 변경되는 버그가 발생했다.
- 아래 코드와 같이 `AttackPower` 인스턴스를 재사용했기 때문이다.

  ```java
    AttackPower attackPower = new AttackPower(20);
    Weapon weaponA = new Weapon(attackPower);
    Weapon weaponB = new Weapon(attackPower);

    weaponA.attackPower.value = 25; // weaponB도 같이 수정됨
  ```

- 이처럼 **`가변 변수는 예상하지 못한 동작을 일으킨다`**. 아래와 같이 재사용하지 못하도록 수정해야 한다.

  ```java
    AttackPower attackPowerA = new AttackPower(20);
    AttackPower attackPowerB = new AttackPower(20);
    Weapon weaponA = new Weapon(attackPowerA);
    Weapon weaponB = new Weapon(attackPowerB);
  ```

- 함수가 매개변수를 받고 값을 리턴하는 것 외에 **`상태를 변경하는 것을 부수 효과라고 부른다`**.
  - 인스턴스 변수 / 전역 변수 / 매개 변수 변경, 파일을 읽고 쓰는 I/O가 모두 해당된다.

#### 함수의 영향 범위 한정하기

- **`부수 효과가 있는 함수는 영향 범위를 예측하기 힘들기 때문에`** 함수를 아래와 같이 설계하는게 좋다.
  - 데이터는 매개 변수로 받는다.
  - 상태를 변경하지 않고 값은 함수의 리턴으로 돌려준다.
- 즉 **`매개변수로 상태를 받고, 상태를 변경하지 않고 값을 리턴하는 함수가 이상적이다`**.

#### 불변으로 만들어서 예상치 못한 동작 막기

- 위 코드에서 value가 가변이기 때문에 부수 효과를 발생할 여지를 남겼었다.
  - 주의해서 코드를 작성할 테니 가변이어도 돼 라는 생각은 스스로를 너무 맹신하는 것이다.
- **`기능 변경 때에 의도하지 않게 부수 효과가 있는 함수가 만들어져서`** 오동작을 일으킬 가능성은 항상 존재한다.
- 따라서 부수 효과의 여지 자체를 없앨 수 있게 변수에 `final`을 붙여 불변으로 만들자.

  ```java
    class AttackPower {
      final int value; // 불변으로 만듬
      AttackPower(final int value) {
        this.value = value;
      }

      // 공격력 강화
      AttackPower reinForce(final AttackPower increment) {
        return new AttackPower(this.value + increment.value);
      }
    }
  ```

- value가 불변이므로 공격력을 변경하려면 `reinForce` 메서드를 사용해야 한다.
  - **`AttackPower 인스턴스를 새로 생성했기 때문에`** 변경전과 변경후는 서로 영향을 주지 않는다.

### 불변과 가변은 어떻게 다루어야 할까

#### 기본적으로 불변

- **`변수를 불변으로 만들면 아래와 같은 장점이 있다`**.
  - 변수의 의미가 변하지 않으므로 혼란을 줄이고 결과를 예측하기 쉽다.
  - 코드의 영향 범위가 한정적이라, 유지 보수가 편리해진다.
- 따라서 기본적으로 불변으로 설계하는 것이 좋고, 최근 언어는 불변이 디폴트가 되도록 만들어지고 있다.

#### 가변으로 설계하는 경우

- 불변이면 값을 변경할 때 인스턴스를 새로 생성해야 한다.
- 대량의 데이터를 처리하는 경우 크기가 큰 인스턴스를 새로 생성하면서 성능에 문제가 된다면 가변을 사용해도 좋다.
- 또한 반복문 내부에서만 사용되는 지역 변수는 가변으로 해도 괜찮다.

#### 코드 외부와 데이터 교환은 국소화하기

- 파일을 읽고 쓰는, DB 쿼리 I/O는 코드 외부의 상태에 의존한다.
- 특별한 이유 없이 외부 상태에 의존하는 코드를 작성하면 동작 예측이 힘들어지므로 문제가 발생할 가능성이 높아진다.
- 최근에는 **`코드 외부와 데이터 교환을 국소화하도록`** 레포지터리 패턴을 많이 사용한다.
  - 레포지터리 패턴은 특정 클래스에 DB 관련 로직을 격리해서 어플리케이션 로직과 섞이지 않도록 한다.

## 응집도

- **`응집도란`** 모듈 내부에 있는 데이터와 로직이 얼마나 강한지 나타내는 지표이다.
- 일반적으로 응집도가 높은 구조는 변경하기 쉽고 바람직한 구조이다.

### static 메서드 오용

- 아래 주문을 관리하는 클래스가 있다.

  ```java
    class OrderManager {
      static int add(int amount1, int amount2) {
        return amount1 + amount2;
      }
    }
  ```

- `static` 메서드를 정의하면 **`클래스의 인스턴스를 생성하지 않고도`** add 메서드를 호출할 수 있다.

  ```java
    // moneyData1, moneyData2는 데이터 클래스
    moneyData1.amount = OrderManager.add(moneyData1.amount1, moneyData2.amount);
  ```

- 이 구조의 문제는 무엇일까? **`데이터는 MoneyData에 있고 데이터를 조작하는 로직은 OrderManager에 있는게 문제다`**.
- `static` 메서드는 인스턴스 변수를 사용할 수 없고, 데이터와 로직 사이에 괴리가 생긴다.

#### 인스턴스 메서드인 척 하는 static 메서드 주의

- `static` 키워드가 없더라도 **`인스턴수 변수를 사용하지 않고 매개변수만 활용하는 메서드도 응집도를 낮춘다`**.

  ```java
    class PaymentManager {
      private int discountRate; // 할인률

      int add(int amount1, int amount2) {
        return amount1 + amount2;
      }
    }
  ```

### 초기화 로직 분산

- 클래스를 잘 설계해도 **`초기화 로직이 분산되어 응집도가 낮은 구조가 될 수 있다`**.

  ```java
    class GiftPoint {
      private static final int MIN_POINT = 0;
      final int value;

      GiftPoint(final int point) {
        // 예외처리
        value = point;
      }

      GiftPoint add(final GiftPoint other) {
        return new GiftPoint(value + other.value);
      }

      GiftPoint consume(final ConsumptionPoint point) {
        return new GiftPoint(value - point.value);
      }
    }
  ```

- 기프트 포인트와 관련된 데이터와 로직이 응집되어 보이지만 아래 코드를 보자

  ```java
    // 일반 회원
    GiftPoint standardMembershipPoint = new GiftPoint(3000);
    // 프리미엄 회원
    GiftPoint PremiumMembershipPoint = new GiftPoint(10000);
  ```

- **`생성자를 public으로 만들면`** 관련된 로직이 분산되어 유지보수하기 힘들어진다.
  - 예를 들어 회원 가입 포인트를 변경하고 싶을 때 소스 코드 전체를 확인해야 한다.

#### private 생성자 + 팩토리 메서드를 사용해 목적에 따라 초기화하기

- 위의 초기화 로직의 분산을 막으려면 **`생성자를 private로 만들고 목적에 따라 팩토리 메서드를 만든다`**.

  ```java
      class GiftPoint {
        private static final int MIN_POINT = 0;
        private static final int STANDARD_MEMBERSHIP_POINT = 3_000;
        private static final int PREMIUM_MEMBERSHIP_POINT = 10_000;
        final int value;

        GiftPoint(final int point) {
          // 예외처리
          value = point;
        }

        static GiftPoint forStandardMembership() {
          return new GiftPoint(STANDARD_MEMBERSHIP_POINT)
        }

        static GiftPoint forPremiumMembership() {
          return new GiftPoint(PREMIUM_MEMBERSHIP_POINT)
        }
      }
  ```

- 생성자를 `private`로 만들면 클래스 내부에서만 인스턴스를 생성할 수 있다.
- `static` 메서드에선 생성자를 호출한다. **`팩토리 메서드는 목적에 따라 만들어 두는 것이 일반적이다`**.
- 이렇게 만들면 신규 가입 포인트와 관련된 로직이 `GiftPoint` 클래스에 응집된다.
- 포인트와 관련된 사양에 변경이 있는 경우 `GiftPoint` 클래스만 변경하면 되고, 다른 클래스의 로직을 찾지 않아도 된다.

### 범용 처리 클래스 (Common/Util)

- `static`으로 실무에서 빈번하게 사용되는 클래스는 `Common, Util` 이라는 이름이 붙게 된다.
- 결국 `static` 메서드이기 때문에 응집도를 낮추게 된다.
- 꼭 필요한 경우가 아니라면 객체지향 설계의 기본으로 돌아가 필요한 클래스를 설계하자

#### 횡단 관심사

- 로그 출력과 오류 확인은 어플리케이션의 모든 동작에 필요한 기능이다.
- 이처럼 다양한 상황에서 넓게 사용되는 기능을 **`횡단 관심사라 부른다`**.
  - ex) `로그 출력, 오류 확인, 디버깅, 예외처리, 캐시 등등`
- 이런 기능은 범용 코드로 만들고 인스턴스화 할 필요가 없으니 static 메서드로 만들어도 좋다.

### 결과를 리턴하는데 매개변수 사용하지 않기

- 매개 변수를 잘못 다루면 응집도가 낮아지게 된다.

  ```java
    class Actor {
      void shift(Location location, int shiftX, int shifty) {
        location.x += shiftX;
        location.y += shiftY;
      }
    }
  ```

- 위 코드는 매개변수를 전달받아 이를 변경하고 있다.
  - 데이터는 `Location`, 조작 로직은 `Actor`로 응집도가 낮은 구조이다.
- 이처럼 매개변수를 리턴하지 말고 데이터와 데이터 조작 조직을 같은 클래스에 배치하자.

  ```java
    class Location {
      final int x;
      final int y;

      Location(final int x, final int y) {
        this.x = x;
        this.y = y;
      }

      Location Shift(final int shiftX, final int shiftY) {
        return new Location(x + shiftX, y + shiftY);
      }
    }
  ```

### 매개변수가 너무 많은 경우

- **`매개변수가 너무 많은 메서드는 응집도가 낮아지기 쉽다`**.
- 메서드에 매개변수를 전달한다는 것은 매개변수를 사용해 어떤 기능을 수행하고 싶다는 의미이다.
- 그래서 매개변수가 많다는 것은 **`많은 기능을 처리하고 싶다는 의미가 된다`**.
- 이런 경우 로직이 복잡해지거나 중복 코드가 생길 가능성이 높아진다.

#### 기본 자료형에 대한 집착

- 아래 `discountedPrice` 메서드는 매개변수와 리턴 값에 모두 기본 자료형만 쓰고 있다.

  ```java
    class Common {
      int discountedPrice(int regularPrice, float discountRate) {
        if (regularPrice < 0) throw new IllegalArgumentException("불라불라");
        if (discountRate < 0) throw new IllegalArgumentException("불라불라");
      }
    }
  ```

- 일반적인 구현 스타일이라고 생각할 수 있지만 다른 클래스에서도 유효성 검사 코드가 중복될 수 있다.
- **`기본 자료형만으로 동작하는 코드를 작성할 수 있다`**. 하지만 관련된 데이터와 로직을 집약하긴 힘들다.
- 데이터는 계산하거나 데이터에 따라 제어 흐름을 전환할 때 사용된다.
  - **`기본 자료형 만으로만 구현하려고 하면`** 계산과 제어 로직이 모두 분산되어 응집도가 낮은 구조가 된다.
- 아래 코드처럼 할인 요금, 정가 할인율을 하나하나의 클래스로 발전시켜 보자.

  ```java
    class RegularPrice {
      final int amount;

      RegularPrice(final int amount) {
        if (amoun < 0) // 예외처리
        this.amount = amount;
      }
    }

    class DiscountPrice {
      final int amount ;

      DiscountPrice(
        final RegularPrice regularPrice,
        final DiscountRate discountRate
      ) {
        // 로직을 사용해서 계산
      }
    }
  ```

- 위와 같이 하면 관련있는 데이터와 로직을 각각의 클래스에 응집할 수 있다.
- 매개변수가 많으면 데이터 하나하나를 매개 변수로 다루지 말고, 그 데이터를 인스턴스 변수로 갖는 클래스를 만들어 활용해보자.

### 메서드 체인

- 아래 코드는 멤버의 갑옷을 변경하는 메서드이다.

  ```java
    // 갑옷 입기
    void equipArmor(int memberId, Armor newArmor) {
      if (party.members[memberId].equipments.canChange) {
        party.members[memberId].equipments.armor = newArmor;
      }
    }
  ```

- 위처럼 점`(.)`으로 여러 메서드를 연결해 리턴 값의 요소에 접근하는 방법을 **`메서드 체인이라고`** 부른다.
- 이 방법도 **`응집도를 낮출 수 있어 좋지 않은 작성 방법이다`**.
  - armor에 할당하는 코드를 어디에서나 작성할 수 있다.
  - 비슷한 코드가 여러 곳에 중복 작성될 가능성이 있다.
  - 접근하는 요소의 사양이 변경되면, 해당 요소에 접근하는 모든 코드를 확인하고 수정해야 한다.
- **`데메테르의 법칙이 있다`**. 사용하는 객체 내부를 알아서는 안 된다는 법칙이다.

#### 묻지말고 명령하기

- SW 설계에선 **`묻지말고, 명령하기(Tell, Don't Ask)라는`** 유명한 격언이 있다.
- 다른 객체의 내부 상태를 판단하거나 제어하려 하지 말고, 메서드로 명령해서 객체가 알아서 판단하고 제어하도록 설계하란 의미다.
- 인스턴스 변수를 `private`로 변경해 외부에서 접근할 수 없게 하고, 외부에선 메서드로 명령해 인스턴스 변수를 제어해야 한다.

  ```java
    class Equipments {
      private boolean canChange;
      private Equipment armor;
      private Equipment head;

      // 갑옷 입기
      void equipArmor(final Equipment newArmor) {
        if (canChange) {
          armor = newArmor;
        }
      }

      // 전체 장비 해제
      void deactivateAll() {
        head = Equipment.EMPTY;
        armor = Equipment.EMPTY;
      }
    }
  ```

- 위와 같이 하면 방어구의 탈착과 관련된 로직이 `Equipments`에 응집된다.
- 방어구와 관련된 요구사항이 변경되었을 때 `Equipments`만 보면 된다. 코드 이곳저곳을 찾을 필요가 없다.

## 조건 분기

- **`if문을 중첩하면 가독성이 크게 떨어지고`** 어디부터 어디까지 if문의 블록인지 이해하기 힘들게 된다.

  ```java
    // 게임에서 마법을 쓰는 경우
    if (member.hitPoint > 0) { // 살아 있는가
      if (member.canAct()) { // 움직일 수 있는가
        if(magic.point <= member.magicPoint) {
          // 마법 시전
        }
      }
    }

    // 중첩과 코드가 복잡한 경우
    if (조건) {
      // 수십 ~ 수백 줄의 코드
      if (조건) {
        // 수십 ~ 수백 줄의 코드
        if (조건) {
        // 수십 ~ 수백 줄의 코드
          if (조건) {
          // 수십 ~ 수백 줄의 코드
          }
        }
      }
    }
  ```

### 조기 리턴으로 중첩 제거하기

- 조기 리턴으로 중첩을 제거하고 가독성이 좋아진 코드

  ```java
    if (member.hitPoint < 0) return;
    if (!member.canAct()) return;
    if (member.magicPoint < magic.point) return;

    // 마법 수행
  ```

- 조기 리턴의 다른 장점은 **`조건 로직과 실행 로직을 분리할 수 있다는 점이다`**.
- 마법을 쓸 수 없는 조건은 앞 부분에 조기 리턴으로 모았고, 마법 발동시 실행 로직은 뒤로 모았다.

### 가독성을 낮추는 else 구문도 조기 리턴으로 해결하기

- `else` 구문도 가독성을 나쁘게 만드는 원인 중 하나이다.
- HP 비율에 따라 건강 상태를 리턴하는 로직을 생각해보자.

  ```java
    float hitPointRate = member.hitPoint / member.maxHitPoint;
    HealthCondition currentHealthCondition;

    if (hitPointRate == 0) {
      currentHealthCondition = HealthCondition.dead;
    } else if (hitPointRate < 0.3) {
      currentHealthCondition = HealthCondition.danger;
    } else if (hitPointRate < 0.5) {
      currentHealthCondition = HealthCondition.caution;
    } else {
      currentHealthCondition = HealthCondition.fine;
    }
  ```

- 조기 리턴을 사용해 아래와 같이 수정해볼 수 있다.

  ```java
    float hitPointRate = member.hitPoint / member.maxHitPoint;

    if (hitPointRate == 0) return HealthCondition.dead;
    if (hitPointRate < 0.3) return HealthCondition.danger;
    if (hitPointRate < 0.5) HealthCondition.caution;

    return HealthCondition.fine;
  ```

- 위 코드는 단순히 가독성이 좋아진 것 외에도 요구사항을 그대로 표현한 형태가 되었다는 의미도 있다.

### switch 조건문 중복

- switch 조건이 어떤 문제를 일으킬 수 있는지 게임을 예로 들어보자.
  - 어떤 게임에서 마법은 다음과 같은 요구사항을 갖는다. `(마법 이름, 매직포인트 소비량(MP), 공격력)`
  - 개발 초기에 마법의 종류는 파이어, 라이트닝 밖에 없었다.
- enum을 사용해 마법 종류를 정의하고 switch 로직을 작성했다.

  ```java
    enum MagicType {
      fire,
      lighting
    }

    class MagicManager {
      String getName(MagicType magicType) {
        String name = '';

        switch (magicType) {
          case fire:
            name = '파이어';
            break;
          case lighting:
            name = '라이트닝'
            break;
        }

        return name;
      }
    }
  ```

#### 같은 형태의 switch 조건문이 여러개 사용되기 시작

- 마법의 종류에 따라 처리가 달라지는 부분은 마법 이름뿐만이 아니다.
- 매직 포인트와 공격력도 모두 마법에 따라 달라진다.

  ```java
    // 매직 포인트
    int costMagicPoint(MagicType magicType, Member member) {
      int magicPoint = 0;

      switch (magicType) {
        case fire:
          magicPoint = 2;
          break;
        case lighting:
          magicPoint = 5;
          break;
      }

      return magicPoint;
    }

    // 공격력도 이런 switch 문을 통해 계산
    int attackPower(MagicType magicType, Member member) {
      int attackPower = 0;
      switch (magicType) {
        ...
      }
      return attackPower;
    }
  ```

- 같은 형태의 switch 조건문을 여러 번 사용하는 것은 좋지 않다.

#### 요구사항 변경 시 수정 누락

- 출시일이 다가와 정신없는데 새로운 마법 헬 파이어가 추가 되었다.
  - 담당자는 이전에 마법 종류 별로 switch 코드를 기억하고 getName 메서드에 case 구문을 추가했다.
  - 출시 후, 헬 파이어 공격력이 너무 약한 것을 발견했는데 확인해보니 attackPower 메서드에 case 구문을 추가하지 않은 것이다.
- 새로운 요구사항으로 마법을 사용하면 테크니컬 포인트를 소비하는 기능이 추가되었다.
  - 다른 팀에서 개발을 담당했는데 추가 된 마법 헬 파이어를 모르고 case에 넣지 않아 문제가 되었다.

#### 폭발적으로 늘어나는 switch 조건문 중복

- 이 예시에서 마법은 세 종류밖에 없었고 처리할 대상은 이름, 매직포인트 소비량, 공격력, 테크니컬 포인트 뿐이었다.
- 실무에서는 훨씬 많은 대상들이 있고, switch 조건문의 중복이 많아지면 주의 깊게 대응해도 실수가 발생할 수 밖에 없다.
- 결국 **`요구사항이 추가될 때 마다 case 구문이 누락될 것이고 버그가 만들어지게 된다`**.

#### 인터페이스로 switch 조건문 중복 해소하기

- **`인터페이스를 사용하면`** 분기 로직을 작성하지 않고도 분기 기능을 구현할 수 있다.
- 아래 면적을 구하는 `Circle, Rectangle` 클래스는 서로 다르다.

  ```java
    interface Shape {
      double area();
    }

    class Rectangle implements Shape {...}
    class Circle implements Shape {...}
  ```

- 하지만 인터페이스를 사용해 조건 분기를 작성하지 않고도 각각의 코드를 실행할 수 있다.

  ```java
    Shape circle = new Circle(10);
    circle.area();

    Shape rectangle = new Rectangle(20):
    rectangle.area();
  ```

#### 인터페이스를 switch 조건문 중복에 응용하기 (전략 패턴)

- 인터페이스의 큰 장점 중 하나는 다른 로직을 같은 방식으로 처리할 수 있다는 점이다.
- 앞서 마법 예시는 switch 조건문을 이용해 이름, 매직포인트 소비량, 공격력, 테크니컬 포인트 소비량을 다르게 처리했다.
- 이를 Magic 인터페이스로 구현해보자.

  ```java
    interface Magic {
      String name();
      int costMagicPoint();
      int attackPower();
      int costTechnicalPoint();
    }
  ```

- 이어서 마법 종류 별로 클래스로 만들어보자.

  ```java
    class Fire implements Magic {
      private final Member;

      Fire(final Member member) {
        this.member = member;
      }

      public String name() {
        return "파이어";
      }

      public int costMagicPoint() {
        return 2;
      }

      public int attackPower() {
        return 20;
      }

      public int costTechnicalPoint() {
        return 0;
      }
    }

    class Lightning implements Magic() {
      private final Member;

      Fire(final Member member) {
        this.member = member;
      }

      public String name() {
        return "라이트닝";
      }

      // 생략
    }

    class HellFire implements Magic() {
      // 생략
    }
  ```

- 이와 같이 구현하면 파이어, 라이트닝, 헬 파이어를 모두 Magic 자료형으로 활용할 수 있다.
- switch 문 대신 Map으로 변경해보자.

  ```java
    final Map<MagicType, Magic> magics = new HashMap<>();

    final Fire fire = new Fire(member);
    final Lightning lightning = new Lightning(lightning);
    final HellFire hellFire = new HellFire(hellFire);

    magics.put(MagicType.fire, fire);
    magics.put(MagicType.lightning, lightning);
    magics.put(MagicType.hellFire, hellFire);
  ```

- 이제 데미지를 계산하기 위해 Map에서 꺼내 attackPower를 호출한다.

  ```java
    void magicAttack(final MagicType magicType) {
      final Magic magic = magics.get(magicType);
      magic.attackPower();
    }
  ```

- 이제 매개 변수로 파이어, 라이트닝, 헬 파이어를 전달함에 따라 조건문처럼 각각의 처리를 하게 된다.
- 이름, MP, 공격력, 테크니컬 포인트를 모두 변경해보자.

  ```java
    void magicAttack(final MagicType magicType) {
      final Magic magic = magics.get(magicType);

      magic.name(); // 호출 후 이름 출력
      magic.attackPower(); // 공격력 계산
      magic.costMagicPoint(); // MP 계산해서 소비
      magic.costTechnicalPoint(); // TP 계산해서 소비
    }
  ```

- switch 조건문을 사용하지 않고도 마법별로 처리를 나누었다.
- 이처럼 인터페이스를 사용해 처리를 전환하는 설계를 전략 패턴이라고 부른다.
- 인터페이스를 활용한 전략 패턴은 그 외에도 장점이 있다.
  - switch 구문을 쓸 때 HellFire 처리를 깜빡 잊었다고 해보자.
  - 인터페이스를 사용하면 컴파일 조차 실패한다. 인터페이스의 메서드는 반드시 구현되어야 하기 때문이다.
  - 따라서 switch 구문처럼 구현하지 않는다는 실수 자체를 방지할 수 있다.

### 조건 분기 중복과 중첩

- 인터페이스는 switch 조건문의 중복 뿐 아니라 다중 중첩된 복잡한 분기를 제거할 수 있다.
- 아래는 온라인 쇼핑몰에서 우수 고객인지 판정하는 로직으로 다음 조건을 모두 만족하면 골드 회원으로 인정된다.

  ```java
    // 총 구매 금액 100만원 이상 & 한 달 구매 횟수 10회 이상 & 반품률 0.1 이하
    boolean isGoldCustomer(PurchaseHistory history) {
      if (history.totalAmount >= 1_000_000) {
        if (history.purchaseFrequencyPerMonth >= 10) {
          if (history.returnRate <= 0.001) {
              return ture;
          }
        }
      }

      return false;
    }
  ```

- 아래는 실버 회원의 판단 로직이다.

  ```java
    // 한 달 구매 횟수 10회 이상 & 반품률 0.1 이하
    boolean isSilverCustomer(PurchaseHistory history) {
      if (history.purchaseFrequencyPerMonth >= 10) {
        if (history.returnRate <= 0.001) {
          return ture;
        }
      }

      return false;
    }
  ```

- 만약 다이아나 브론즈 등급이 추가되고, 비슷한 조건들이 사용된다면 어떻게 해야 할까?

#### 정책 패턴으로 조건 집약하기

- 이러한 상황에서 유용하게 활용할 수 있는 패턴이 정책 패턴이다.
- 조건을 부품처럼 만들고 부품으로 만든 조건을 조합해서 사용한다.
- 우선 아래와 같이 인터페이스를 하나 만든다.

  ```java
    interface ExcellentCustomerRule {
      // 구매 조건을 만족해야 true
      boolean ok(final purchaseHistory history);
    }
  ```

- 골드 회원이 되려면 3개의 조건을 만족해야 한다. 각 조건을 인터페이스를 구현하여 만든다.

  ```java
    // 골드 회원 구매 금액 규칙
    class GoldCustomerPurchaseAmountRule implements ExcellentCustomerRule {
      public boolean ok(final PurchaseHistory history) {
        return history.totalAmount >= 1_000_000;
      }
    }

    // 구매 빈도 규칙
    class PurchaseFrequencyRule implements ExcellentCustomerRule {
      public boolean ok(final PurchaseHistory history) {
        return history.purchaseFrequencyPerMonth >= 10;
      }
    }

    // 반품률 규칙
    class ReturnRateRule implements ExcellentCustomerRule {
      public boolean ok(final PurchaseHistory history) {
        return history.returnRate <= 0.001;
      }
    }
  ```

- 이어서 정책 클래스를 만든다. add 메서드로 규칙을 넣고 complyWithAll 메서드에서 모든 규칙을 만족하는지 확인한다.

  ```java
    class ExcellentCustomerPolicy {
      private final Set<ExcellentCustomerRule> rules;
      ExcellentCustomerPolicy() {
        rules = new HashSet();
      }

      void add(final ExcellentCustomerRule rule) {
        rules.add(rule);
      }

      boolean complyWithAll(final PurchaseHistory history) {
        for (ExcellentCustomerRule each: rules) {
          if (!each.ok(history)) return false;
        }
        return true;
      }
    }
  ```

- 사용하는 쪽에선 골드 회원의 조건 3가지를 추가하고 판정한다.

  ```java
    ExcellentCustomerPolicy goldCustomerPolicy = new ExcellentCustomerPolicy();
    goldCustomerPolicy.add(new GoldCustomerPurchaseAmountRule());
    goldCustomerPolicy.add(new PurchaseFrequencyRule());
    goldCustomerPolicy.add(new ReturnRateRule());

    goldCustomerPolicy.complyWithAll(purchaseHistory); // 골드 회원 조건 검증
  ```

- if 조건문이 complyWithAll 메서드 내부에 하나만 있어 로직이 단순해졌다.
- 이런 경우 골드 회원과 무관한 로직을 삽입할 가능성이 있으니 확실하게 골드 회원을 판단하는 클래스를 만든다.

  ```java
    class GoldCustomerPolicy {
      private final ExcellentCustomerPolicy policy;

      GoldCustomerPolicy() {
        policy = new ExcellentCustomerPolicy();
        policy.add(new GoldCustomerPurchaseAmountRule());
        policy.add(new PurchaseFrequencyRule());
        policy.add(new ReturnRateRule());
      }

      boolean complyWithAll(final PurchaseHistory history) {
        return policy.complyWithAll(history);
      }
    }
  ```

- 실버 회원도 같은 방법으로 만들 수 있고 규칙이 재사용되고 있으므로 괜찮은 클래스 구조라 할 수 있다.

  ```java
    class SilverCustomerPolicy {
      private final ExcellentCustomerPolicy policy;

      SilverCustomerPolicy() {
        policy = new ExcellentCustomerPolicy();
        policy.add(new PurchaseFrequencyRule());
        policy.add(new ReturnRateRule());
      }
    }
  ```
