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
- `주석 A` 부분은 실제 플레이어 공격력의 총량이므로 아래와 같이 수정하면 어떤 값을 계산하는데 어떤 값을 사용하는지 관계를 파악하기 쉽다.

  ```java
    // best
    int totalPlayerAttackPower = playerArmPower + playerWeaponPower;
    int totalEnemyDefence = ((enemyBodyDefence + eenemyArmorDefence) / 2);

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

- 잘 만들어진 클래스는 **`다음 두 가지로`** 구성된다.
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
- **`매개 변수와 지역 변수도 final 키워드를 붙여서`** 변경될 수 없도록 하자

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

- **`응집도란`** 모듈 내부에 있는 데이터와 로직이 얼마나 응집되어 있는지를 나타내는 지표다.
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
- 이처럼 매개변수를 리턴하지 말고 데이터와 데이터 조작 로직을 같은 클래스에 배치하자.

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
- **`방어구와 관련된 요구사항이 변경되었을 때 Equipments만`** 보면 된다. 코드 이곳저곳을 찾을 필요가 없다.

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
  - 담당자는 이전에 마법 종류 별로 switch 코드를 기억하고 `getName` 메서드에 case 구문을 추가했다.
  - 출시 후, 헬 파이어 공격력이 너무 약한 것을 발견했는데 확인해보니 `attackPower` 메서드에 case 구문을 추가하지 않은 것이다.
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

- 인터페이스의 큰 장점 중 하나는 **`다른 로직을 같은 방식으로 처리할 수 있다는 점이다`**.
- 앞서 마법 예시는 switch 조건문을 이용해 이름, 매직포인트 소비량, 공격력, 테크니컬 포인트 소비량을 다르게 처리했다.
- 이를 `Magic` 인터페이스로 구현해보자.

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

- 이와 같이 구현하면 파이어, 라이트닝, 헬 파이어를 모두 `Magic` 자료형으로 활용할 수 있다.
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

- 이제 데미지를 계산하기 위해 Map에서 꺼내 `attackPower`를 호출한다.

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

- **`switch 조건문을 사용하지 않고도`** 마법별로 처리를 나누었다.
- 이처럼 **`인터페이스를 사용해 처리를 전환하는 설계를 전략 패턴이라고 부른다`**.
- 인터페이스를 활용한 전략 패턴은 그 외에도 장점이 있다.
  - switch 구문을 쓸 때 HellFire 처리를 깜빡 잊었다고 해보자.
  - **`인터페이스를 사용하면 컴파일 조차 실패한다`**. 인터페이스의 메서드는 반드시 구현되어야 하기 때문이다.
  - 따라서 switch 구문처럼 구현하지 않는다는 실수 자체를 방지할 수 있다.

### 조건 분기 중복과 중첩

- 인터페이스는 switch 조건문의 중복 뿐 아니라 **`다중 중첩된 복잡한 분기를 제거할 수 있다`**.
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

- 이러한 상황에서 유용하게 활용할 수 있는 패턴이 **`정책 패턴이다`**.
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

- 이어서 정책 클래스를 만든다. `add` 메서드로 규칙을 넣고 `complyWithAll` 메서드에서 모든 규칙을 만족하는지 확인한다.

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

- if 조건문이 `complyWithAll` 메서드 내부에 하나만 있어 로직이 단순해졌다.
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

### 자료형 확인에 조건 분기 사용하지 않기

- 인터페이스는 **`조건 분기를 제거할 때 활용할 수 있다`**.
- 하지만 인터페이스를 충분히 이해하지 못하고 사용하면 조건 분기가 줄어드지 않는 경우도 있다.
- 호텔 숙박 요금을 예로 일반 객실과 비싼 객실이 있다고 하자

  ```java
    interface HotelRates {
      Money fee(); // 요금
    }

    class RegularRates implements HotelRates {
      public Money fee() {
        return new Money(70_000);
      }
    }

    class PremiumRates implements HotelRates {
      public Money fee() {
        return new Money(120_000);
      }
    }
  ```

- 위 까지는 좋은데 성수기처럼 특정 기간에 숙박 요금을 높게 설정하는 경우 아래처럼 구현했다고 가정하자

  ```java
    Money busySeasonFee;
    if (hotelRates instanceof RegularRates) {
      busySeasonFee = hotelRates.fee().add(new Money(30_000));
    } else if (hotelRates instanceof PremiumRates) {
      busySeasonFee = hotelRates.fee().add(new Money(50_000));
    }
  ```

- 모처럼 인터페이스를 사용했는데 조건 분기가 그대로 있다.
  - 특정 기간에 적용되는 요금이 추가된다면 분기는 더 늘어나게 될 것이다.
- 성수기 요금도 인터페이스로 변경해보자

  ```java
    interface HotelRates {
      Money fee(); // 요금
      Money busySeasonFee(); // 성수기 요금
    }

    class RegularRates implements HotelRates {
      public Money fee() {
        return new Money(70_000);
      }

      public Money busySeasonFee() {
        return new fee().add(new Money(30_000));
      }
    }

    class PremiumRates implements HotelRates {
      public Money fee() {
        return new Money(120_000);
      }

      public Money busySeasonFee() {
        return new fee().add(new Money(50_000));
      }
    }
  ```

### 인터페이스 사용 능력이 중급으로 올라가는 첫 걸음

- **`인터페이스를 잘 사용하는지가 설계 능력의 전환점 중 하나이다`**.
- 저자의 개인 생각으로 설계 레벨에 따른 사고 방식의 차이를 아래와 같이 제시함
  - **`초보자`**: if, switch 조건 그냥 씀 + 분기마다 처리는 그냥 로직을 작성
  - **`중급자`**: 분기는 인터페이스 설계 사용 + 분기마다 로직은 클래스에게 위임
- 조건 분기를 써야 하는 상황에는 일단 인터페이스 설계를 떠올리자! 새겨두기만 해도 방식 자체가 달라질 것이다.

### 플래그 매개변수

- 메서드 내부에서 기능을 전환하는 `boolean` 자료형의 매개변수를 **`플래그 매개변수라고`** 부른다.

  ```java
    // 메서드 호출 로직
    damage(true, damageAmount);

    // 메서드
    void damage(boolean damageFlag, int damageAmount) {
      if (damageFlag) {
        // 물리 데미지
      } else {
        // 마법 데미지
      }
    }
  ```

- 플래그 매개변수를 받는 메서드는 **`어떤 일을 하는지 예측하기 힘들어 가독성이 떨어지고 개발 생산성이 저하된다`**.
- boolean 자료형 뿐 아니라 아래와 같이 int 자료형을 사용해 기능을 전환하는 경우에도 문제가 발생한다.

  ```java
    void execute(int processNumber) {
      if (processNumber == 0) {...}
      else if (processNumber == 1) {...}
      else if (processNumber == 2) {...}
    }
  ```

### 메서드 분리하기

- **`플래그 매개변수를 받는 메서드는`** 내부적으로 **`여러 기능을 수행하게 된다`**.
- **`메서드는 하나의 기능만 가지도록 설계하는 것이 좋기 때문에`** 기능별로 분리하고 메서드에 맞는 이름을 붙이면 가독성이 좋아진다.

  ```java
    void hitDamage(final int damageAmount) {
      // 물리 데미지 입은 로직
    }

    void magicDamage(final int damageAmount) {
      // 마법 데미지 입은 로직
    }
  ```

### 전환은 전략 패턴으로 구현하기

- 히트 데미지와 매직 데미지를 **`전환해야 하는 경우 boolean 자료형을 사용하면`** 플래그 매개변수로 되돌아 가게 된다.
- **`플래그 매개변수가 아니라 전략 패턴을 사용하자`**. 전환 대상은 히트 데미지와 매직 데미지이다.
- `Damage` 인터페이스를 구현하고 전환하고자 하는 로직을 각 클래스에게 맡긴다.

  ```java
    interface Damage {
      void execute(final int damageAmount);
    }

    class HitDamage implements Damage {
      public void execute(final int damageAmount) {
        ...
      }
    }

    class MagicDamage implements Damage {
      public void execute(final int damageAmount) {
        ...
      }
    }

    Enum DamageType {
      hit,
      magic
    }

    private final Map<DamageType, Damage> damages;

    void applyDamage(final DamageType damageType, final int damageAmount) {
      final Damage damage = damages.get(damageType);
      damage.execute(damageAmount);
    }

    // 호출 로직
    applyDamage(DamageType.magic, damageAmount);
  ```

- 가독성이 높아진 것 외에도 **`전략 패턴으로 설계하면`** 새로운 종류의 데미지가 추가되었을 때 쉽게 대응할 수 있다.

## 컬렉션

- 배열과 List 같은 **`컬렉션을 따라다니는 악마를`** 소개하고 퇴치 방법을 알아본다.

### 이미 존재하는 기능을 다시 구현하지 말자

- 소지품에 감옥 열쇠가 있는지 확인하는 아래 코드는 **`for문 내부에 반복문이 있어 가독성이 좋지 않다`**.

  ```java
    boolean hasPrisonKey = false;
    for (Item each: items) {
      if (each.name.equals('감옥 열쇠')) {
        hasPrisonKey = true;
        break;
      }
    }
  ```

- 같은 기능을 아래와 같이 구현할 수 있다.

  ```java
    boolean hasPrisonKey = items.stream().anyMatch(item -> item.name.equals('감옥 열쇠'))
  ```

- `anyMatch` 메서드를 알고 있으면 복잡한 로직을 직접 구현하지 않아도 된다.
- **`바퀴의 재발명`**
  - 이미 널리 사용되고 있는 기술과 해결법이 있는데 이를 몰라 비슷한 것을 만들어 내는 것을 **`바퀴의 재발명이라고 한다`**.
  - 참고로 존재하는 것보다 좋지 못한 결과물을 만들어내면 네모난 바퀴의 재발명이라 부른다.
  - 라이브러리가 어떻게 동작하는지, 구조를 학습하는 과정에선 이런 과정이 도움이 된다.

### 반복문 내부의 조건 분기 중첩

- **`반복문 내부에서 특정 조건을 만족시키는 요소에 대해서만`** 어떤 작업을 수행하고 싶은 경우가 있다.
- RPG 파티원 중 독에 의해 중독된 멤버들의 HP를 감소하는 로직이 아래와 같이 있다.

  ```java
    for (Member member: members) {
      if (member.hitPoint > 0) {
        if (member.containState(StateType.poison)) {
          member.hitPoint -= 10;
          if (member.hitPoint <= 0) {
            member.hitPoint = 0;
            member.addState(StateType.dead);
          }
        }
      }
    }
  ```

- 이럴 때 **`early continue를 사용해`** 가독성을 높일 수 있다.

  ```java
    for (Member member: members) {
      if (member.hitPoint == 0) continue;
      if (!member.containState(StateType.poison)) continue;

      member.hitPoint -= 10;
      if (member.hitPoint > 0) continue;

      member.hitPoint = 0;
      member.addState(StateType.dead);
    }
  ```

- 이 외에도 **`early break를`** 통해 중첩을 제거하고 가독성을 높일 수 있다.

### 응집도가 낮은 컬렉션 처리

- **`컬렉션에 대한 추가 작업도 응집도가 낮아지기 쉽다`**.
- RPG에서 필드 맵을 관리하는 클래스에 파티 멤버를 추가하는 아래 로직이 있다.

  ```java
    class FieldManager {
      void addMember(List<Member> members, Member newMember) {
        ...
      }
    }
  ```

- 그런데 필드 맵 말고도 게임에서 멤버를 추가하는 시점이 있다.

  ```java
    // 특별한 이벤트를 관리하는 클래스
    class SpecialEventManager {
      void addMember(List<Member> members, Member newMember) {
        ...
      }
    }
  ```

- 이처럼 **`컬렉션과 관련된 작업을 처리하는 코드가 여기저기에 구현되어 응집도가 낮아질 가능성이 높다`**.

#### 컬렉션 처리를 캡슐화하기

- 컬렉션과 관련된 응집도가 낮아지는 문제는 **`일급 컬렉션 패턴을 사용해 해결할 수 있다`**.
- **`일급 컬렉션이란`** 컬렉션과 관련된 로직을 캡슐화하는 디자인 패턴이다.
- 클래스의 설계 원리를 반영하면 일급 컬렉션은 아래 요소로 구성된다.
  - 컬렉션 자료형의 인스턴스 변수
  - 컬렉션 자료형의 인스턴스 변수에 잘못된 값이 할당되지 않게 막고, 정상적으로 조작하는 메서드
- 파티의 멤버 컬렉션을 인스턴스 변수로 가지는 `Party` 클래스를 설계해보자.

  ```java
    class Party {
      private final List<Member> members;

      Party() {
        members = new ArrayList<Member>();
      }

      void add(final Member newMember) {
        members.add(newMember);
      }
    }
  ```

- 위의 add 메서드는 **`members의 요소가 변화(추가)되는 부수 효과가 발생한다`**.
- 부수 효과를 막기 위해 **`새로운 리스트를 생성하는 형태로`** add 메서드를 구현한다.

  ```java
    class Party {
      void add(final Member newMember) {
        List<Member> members = new ArrayList<>(this.members);
        members.add(newMember);
        return new Party(members);
      }
    }
  ```

- 위와 같이 하면 원래 **`members를 변화시키지 않아 부수 효과를 막을 수 있다`**.
- 파티의 인원수를 확인하는 `isFull` 메서드 등, 컬렉션과 컬렉션을 조작하는 로직을 한 클래스에 응집한 구조로 만들 수 있다.

  ```java
    class Party {
      static final int MAX_MEMBER_COUNT = 4;
      private final List<Member> members;

      Party() {
        members = new ArrayList<Member>();
      }

      private Party(List<Member> members) {
        this.members = members;
      }

      void add(final Member newMember) {
        if (isFull()) throw new IllegalArgumentException("불라불라");

        List<Member> members = new ArrayList<>(this.members);
        members.add(newMember);
        return new Party(members);
      }

      boolean isFull() {
        return members.size() == MAX_MEMBER_COUNT;
      }
    }
  ```

#### 외부로 전달할 때 컬렉션의 변경 막기

- 파티 멤버 전원의 상태를 표시하는 기능이 추가된다면 members에 접근해 전체 데이터를 참조할 수 있어야 한다.
- **`인스턴스 변수 그대로 외부에 전달하면`** Party 클래스 **`외부에서 마음대로 멤버를 추가하고 제거할 수 있다`**.
- **`외부로 전달할 때는 컬렉션이 요소를 변경하지 못하게 막아두는게 좋다`**.

  ```java
    class Party {
      // 생략
      List<Member> members() {
        return members.unmodifiableList(); // 요소를 추가하거나 제거할 수 없다.
      }
    }
  ```

## 강한 결합

- **`결합도란`** 모듈 사이의 의존도를 나타내는 지표이다.
- 어떤 클래스가 **`다른 클래스에 많의 의존하고 있는 구조를 강한 결합이라고`** 부르며 변경하기 힘들어진다.

### 결합도와 책무

- 온라인 쇼핑몰에 할인 서비스가 추가되었다.
- 할인의 사양은 상품 하나당 `3,000`원을 할인하고 최대 `200_000`까지 상품 추가가 가능하다.

  ```java
    class DiscountManager {
      List<Product> discountProducts;
      int totalPrice;

      boolean add(Product product, ProductDiscount) {
        // 예외 처리 코드 생략

        int discountPrice = getDiscountPrice(product.price);
        int tmp;

        if (productDiscount.canDiscount) {
          tmp = totalPrice + discountPrice;
        } else {
          tmp = totalPrice + product.price;
        }

        if (tmp <= 200000) { // 가격 총합이 20만원 이내인 경우 상품 리스트에 추가
          totalPrice = tmp;
          discountProducts.add(product);
          return true;
        } else {
          return false;
        }
      }

      static int getDiscountPrice(int price) {
        int discountPrice = price - 3000;
        if (discountPrice < 0) {
          discountPrice = 0;
        }
        return discountPrice;
      }
    }
  ```

#### 로직의 위치에 일관성이 없음

- **`할인 서비스 로직은 로직의 위치 자체에 문제가 있다`**.
  - `DiscountManager`가 상품 정보 확인 말고도 할인 가격 계산, 할인 적용 여부 판단 등 **`너무 많은 일을 한다`**.
  - `Product`가 직접 해야 하는 유효성 검사 로직이 `DiscountManager` 내부에 있다.
- 이런 클래스 설계가 바로 **`책임을 고려하지 않은 설계라고`** 할 수 있다.

#### 단일 책임 원칙

- **`SW의 책임이란 자신의 관심사와 관련해서 정상적으로 동작하도록 제어하는 것이라고 생각할 수 있다`**.
- **`단일 책임 원칙은`** 클래스가 담당하는 책임은 하나로 제한해야 한다라는 설계 원칙이다.

#### 책임이 하나가 되게 클래스 설계하기

- 단일 책임 원칙 위반으로 만들어진 악마를 퇴치하려면 **`단일 책임 원칙을 지키도록 설계를 바꿔야 한다`**.
- 상품의 가격을 나타내는 `Price` 클래스를 만들고 유효성 검사 과정을 추가하자.
- 유효성 검사와 관련된 책임을 모두 `Price` 클래스에서 지므로, 다른 곳에 유효성 검사 코드가 중복될 일은 없어진다.

  ```java
    class RegularPrice {
      private static final int MIN_AMOUNT = 0;
      final int amount;

      RegularPrice(final int amount) {
        if (amount < MIN_AMOUNT) throw new IllegalArgumentException("불라불라");
        this.amount = amount;
      }
    }
  ```

- 일반 할인을 책임지는 클래스도 만들어보자.

  ```java
    class RegularDiscountedPrice {
      private static final int MIN_AMOUNT = 0;
      private static final int DISCOUNT_AMOUNT = 3000;
      final int amount;

      RegularDiscountedPrice(final RegularPrice price) {
        int discountedAmount = price.amount - DISCOUNT_AMOUNT;
        if (discountedAmount < MIN_AMOUNT) {
          discountedAmount = MIN_AMOUNT;
        }

        amount = discountedAmount;
      }
    }
  ```

- 여름 할인을 책임지는 클래스도 만들어보자.

  ```java
    class SummerDiscountedPrice {
      private static final int MIN_AMOUNT = 0;
      private static final int DISCOUNT_AMOUNT = 4000;
      final int amount;

      SummerDiscountedPrice(final RegularPrice price) {
        int discountedAmount = price.amount - DISCOUNT_AMOUNT;
        if (discountedAmount < MIN_AMOUNT) {
          discountedAmount = MIN_AMOUNT;
        }

        amount = discountedAmount;
      }
    }
  ```

- 클래스가 정가, 일반 할인가격, 여름 할인 가격으로 구분되어 있다.
- 할인과 관련된 **`사양이 변경되어도 서로 영향을 주지 않는 이런 구조를 느슨한 결합이라고 부른다`**.

#### DRY 원칙의 잘못된 적용

- 위 예시에서 `RegularDiscountedPrice, SummerDiscountedPrice`의 로직은 대부분 같다.
- 이를 보고 `중복 코드 아닌가?` 생각할 수 있지만 아래 요구사항으로 바뀌면 중복이 아니게 된다.
  - 여름 할인 가격은 정가에서 5% 할인한다
- **`책임을 생각하지 않고 중복을 제거하면 안된다`**. 무리하게 중복을 제거하려 하면 단일 원칙 책임이 깨지게 된다.
- **`DRY 원칙은 반복을 피해라는 의미인데`**, 일부 사람들은 코드 중복을 절대 허용하지 마라로 받아들인다.
- 같은 로직, 비슷한 로직이라도 **`개념이 다르면 중복을 허용해야 한다`**.
  - 일반 할인과 여름 할인은 서로 다른 개념이다.
  - DRY는 같은 개념 내에서 반복을 하지 말라는 의미이다.

### 다양한 결합 사례와 대처 방법

#### 상속과 관련된 강한 결합

- **`상속은 주의해서 다루지 않으면 곧바로 강한 결합 구조를 유발하는`** 위험한 문제가 발생한다.
- **`상속에서는 서브 클래스는 슈퍼 클래스에 굉장히 크게 의존한다`**.
  - 따라서 서브 클래스는 슈퍼 클래스의 구조를 신경써야 하고, 슈퍼 클래스의 변화를 놓치는 순간 버그가 발생할 수 있다.
- 슈퍼 클래스 의존으로 인한 강한 결합을 피하려면 **`상속보다 컴포지션(합성)을 사용하면 좋다`**.
  - 컴포지션이란 사용하고 싶은 클래스를 인스턴스 변수로 갖고 사용하는 것을 의미한다.
- 상속을 사용하면 서브 클래스가 슈퍼 클래스의 로직을 그대로 사용하므로 슈퍼 클래스가 공통 로직을 두는 장소로 사용된다.
  - 위의 예시에서 일반 할인과 여름 할인을 상속으로 사용하고 `getDiscountedPrice`를 공통으로 사용했다고 가정하자.
  - 일반 할인과 여름 할인이라는 두 가지 책임을 지게 되므로 단일 원칙 책임을 위반해서 좋은 구현이라고 말할 수 없다.
- 서브 클래스가 일부는 부모의 메서드를 그대로 쓰고 일부는 오버라이딩을 시작한다고 가정하자.
  - 이때 물려받아 그대로 쓰는 메서드 내부에서 서브 클래스가 오버라이딩 하는 메서드를 사용한다.
  - 이런 경우 각 서브클래스의 오버라이딩 메서드는 **`부모의 물려받은 메서드를 자세하게 알아야 한다`**.
- 이렇게 **`슈퍼/서브 클래스간 강한 결합이 되면`** 로직을 추적하기가 매우 어려워지며 요구사항 변경이 매우 힘들어진다.
- 상속도 설계를 잘하면 아무런 문제가 없다. 하지만 강한 결합과 로직 분산 등의 악마를 불러들이므로 신중하게 사용해야 한다.

#### 인스턴스 변수 별로 클래스 분할이 가능한 로직

- 아래 코드는 책임이 완전히 다른 메서드들이 하나의 클래스 안에 정의되어 있다.

  ```java
    class Util {
      private int reservationId;
      private ViewSettings viewSettings;

      void cancelReservation() {
        ... // reservationId 사용
      }

      void darkMode() {
        ... // viewSettings 사용
      }

    }
  ```

- 위와 같은 클래스는 각각의 역할에 따라 클래스를 분리해야 한다.

#### 특별한 이유 없이 public 사용하기

- 이유 없이 `public`으로 만들면 관계를 맺지 않길 원하는 클래스끼리 결합되어 영향범위가 확대된다.
- 강한 결합을 피하려면 **`외부에 공개할 필요가 있는 클래스와 메서드만 public으로 선언하자`**.

#### private 메서드가 너무 많다는 것은 책임이 많다는 것

- 규모가 점점 커진 온라인 쇼핑몰의 주문을 담당하는 클래스이다.

  ```java
    class OrderService {
      private int calculateDiscountPrice(int price) {
        // 할인 가격 계산 로직
      }

      private List<Product> getProductBrowsingHistory(int userId) {
        // 최근 본 상품 리스트를 확인하는 로직
      }
    }
  ```

- 위 코드는 주문 시 할인을 적용하거나, 최근 본 상품을 곧바로 주문하고 싶은 기능이 반영된 클래스이다.
- 책임의 관점에서 생각해보면 가격 할인과, 최근 본 상품은 주문과는 다른 책임이다.
- **`private 메서드가 너무 많이 쓰인 클래스는`** 많은 책임을 갖고 있을 가능성이 높으니 책임이 다르다면 분리하자.

#### 높은 응집도를 오해해서 생기는 강한 결합

- 높은 응집도를 잘못 이해해서 강한 결합이 발생하는 경우가 있다.

  ```java
    // 판매 가격 클래스
    class SellingPrice {
      final int amount;

      SellingPrice(final int amount) {
        if (amount < 0) throw new IllegalArgumentException("불라불라");
        this.amount = amount;
      }

      int calcSellingCommission() {
        // 판매 수수료 계산 로직
      }

      int calcDeliveryCharge() {
        // 배송비 계산하기
       }

      int calcShoppingPoint() {
        // 쇼핑 포인트 계산
      }
    }
  ```

- 일부 엔지니어는 판매 수수료와 배송비는 판매 가격과 관련이 깊을 것이다 생각해 위와 같이 작성했다.
- 하지만 판매 가격과 쇼핑 포인트, 배송비, 판매 수수료는 판매 가격과는 **`다른 개념이다`**.
- **`응집도를 생각해 관련이 깊다고 생각한 로직을 한 곳에 모으려 했지만`** 결과적으로 강한 결합을 만들었다.
  - 이런 상황은 자주 일어나고 누구라도 빠질 수 있는 함정이다.
  - 그렇기 때문에 결합이 느슨하고 응집도가 높은 설계를 한 덩어리로 묶어 이야기하곤 한다.
- 각 개념을 클래스로 분할하고 값 객체로 설계하는게 좋다.
- 어떤 개념의 값을 사용해 다른 개념의 값을 구하고 싶을 때는 생성자에 매개변수로 계산에 사용할 값을 전달한다.

  ```java
    class SellingCommission {
      private static final float SELLING_COMMISSION_RATE = 0.05f;
      final int amount;

      SellingCommission(final SellingPrice price) {
        amount = (int)(price.amount * SELLING_COMMISSION_RATE);
      }
    }

    class DeliveryCharge {
      private static final int DELIVERY_FREE_MIN = 20000;
      final int amount;

      DeliveryCharge(final SellingPrice price) {
        amount = DELIVERY_FREE_MIN <= price.amount ? 0 : 5000;
      }
    }

    class ShoppingPoint {
      private static final float SHOPPING_POINT_RATE = 0.01f;
      final int value;

      ShoppingPoint(final SellingPrice price) {
        value = price.amount * SHOPPING_POINT_RATE;
      }
    }
  ```

#### 스마트 UI, 거대 데이터 클래스, 트랜잭션 스크립트 패턴

- 스마트 UI는 **`화면과 관련없는 책임이 구현되어 있는 클래스이다`**.
  - 예를 들어 복잡한 금액 계산 로직을 프런트에 구현하면 디자인을 변경할 때 변경하기 힘들게 된다.
- 수 많은 인스턴스 변수와 많은 기능을 가진 **`거대 데이터 클래스도 다양한 버그를 발생하게 된다`**.
- 메서드 내부에 일련의 처리가 하나하나 길게 작성된 구조를 **`트랜잭션 스크립트 패턴이라 부른다`**.
  - 메서드 하나가 길에는 수백 줄의 거대한 로직을 갖게 되며 변경하기 매우 어려워진다.

## 설계의 건정성을 해치는 여러 악마

### 데드 코드

- **`데드 코드란 절대로 실행되지 않는 조건 내부에 있는 코드이다`**. 또한 도달 불가능한 코드라고 부른다.
- 이런 코드는 가독성을 떨어뜨리고 사양 변경에 의해 실행되면 버그가 될 가능성이 있다.
- 데드 코드는 발견 즉시 제거하는게 좋다.

### YAGNI(You Aren't Gonna Need It) 원칙

- 실제 개발을 할 떄 **`미래를 예측하고 미리 만들어 두는 경우가 있다`**.
- 이렇게 미리 구현한 로직은 실제로 거의 사용되지 않고 버그의 원인이 되기도 한다.
- **`YAGNI 원칙은`** 지금 필요 없는 기능은 만들지 말라는 원칙이다. 미리 구현하면 어떤 문제가 있을까?
  - SW에 대한 요구는 변하는데 과거에 예측해서 만들어둔 기능은 실행되어도 **`현재 사양에 없기 때문에 버그가 될 수 있다`**.
  - 읽는 사람을 혼란스럽게 만들고 가독성을 낮추게 된다.

### 매직 넘버

- **`설명이 없는 숫자는 개발자를 혼란스럽게 만든다`**.
- 아래 코드는 웹툰 서비스에 사용된다고 가정하고 만든 코드이다. 60이라는 숫자는 뭘까?

  ```java
    class ComicManager {
      boolean isOk() {
        return value >= 60;
      }
    }
  ```

- 60은 웹툰을 체험 구독할 때 필요한 포인트이다. 숫자에 대한 설명이 없다면 의도를 알기 어려워진다.
- 이처럼 로직 내부에 직접 작성되어 있어서, **`의미를 알기 힘든 숫자를 매직 넘버라고 부른다`**.
  - 매직 넘버는 구현자 본인만 의도를 이해할 수 있다.
  - 동일한 값이 여러 위치에 등장하여 중복 코드를 만들어 낸다.
- 이런 매직 넘버는 상수를 이용해 가독성을 높이고, 상수를 통해 코드의 중복을 줄일 수 있다.

### 전역 변수

- 모든 곳에서 참조할 수 있고 조작할 수 있는 변수는 어떻게 보면 편리하다고 생각할 수 있다.
- 하지만 전역 변수는 아래와 같은 문제들이 있다.
  - 어디에서 어떤 시점에 값을 변경했는지 파악하기 힘들다.
  - 전역변수를 참조하고 있는 로직을 변경할 때, 해당 변수를 참조하는 다른 로직에 버그가 발생하는지 검토해야 한다.

### null 문제

- 아래 캐릭터의 전체 방어력을 리턴하는 메서드를 보자.

  ```java
    class Member {
      private Equipment head;
      private Equipment body;
      private Equipment arm;
      private int defence;

      int totalDefence() {
        return defence + head + body + arm;
      }
    }
  ```

- 그런데 이 코드를 실행하면 방어구를 착용하지 않은 상태를 null로 표현하기에 `NullPointerException`이 발생하는 경우가 있다.

  ```java
    class Member {
      void takeOffAllEquipments() {
        head = null;
        body = null;
        arm = null;
      }
    }
  ```

- 물론 null 체크를 통해 계산하면 문제는 발생하지 않는다.
- 하지만 **`null이 들어갈 수 있다고 전제하고 로직을 만들면 모든 곳에서 null 체크를 해야 한다`**.
  - null 체크 코드가 많아져 가독성이 떨어지고 실수로 null 체크를 안하면 곧바로 버그가 된다.
- null은 메모리 접근과 관련된 문제를 방지하기 위한 구조로서 null 자체가 잘못된 처리를 의미한다.
  - 그런데 정보가 입력되지 않은 상태를 null로 표현하는 코드가 많다.

#### null을 리턴하거나 전달하지 않기

- null 체크를 하지 않으려면 **`애초에 null을 다루지 않게 만들어야 한다`**.
  - null을 리턴하지 않는 설계 / null을 전달하지 않는 설계
- 위 방어구 예시에서 방어구를 착용하지 않은 상태를 null이 아니라 `Equipment`의 EMPTY로 만들 수있다.

  ```java
    class Equipment {
      static final Equipment EMPTY = new Equipment("장비 없음", 0, 0, 0);

      final String name;
      final int price;
      final int defence;
      final int magicDefence;
    }
  ```

#### null 안전

- **`null 안전이란 null에 의한 오류가 아예 발생하지 않도록 만드는 구조다`**.
- 코틀린의 경우 null 안전 자료형이 있는데, null을 아예 저장할 수 없게 만드는 자료형이다.
- 언어에서 null 안전을 지원한다면 적극적으로 사용하는게 좋다.

### 예외를 catch 하고 무시하는 코드

- 아래처럼 예외를 무시하는 코드는 굉장히 사악한 코드이다.

  ```java
    try {
      reservations.add(product)
    } catch (Exception e) {
    }
  ```

#### 원인 분석을 어렵게 만듬

- 위 코드의 **`문제는 오류가 나도, 오류를 탐지할 방법이 없다는 것이다`**.
- 예외를 무시하면 잘못된 상태를 곧바로 확인할 수 없고 이후 서비스 사용자에 의해 보고될 가능성이 높아진다.
- 문제가 발생해도 어느 시점에 어떤 코드에서 문제가 발생했는지 빠르게 대응하기 힘들어진다.

#### 문제가 발생했다면 소리치기

- 잘못된 상태에서 계속 처리를 진행하는 것은 폭탄을 들고 돌아다니는 것과 같다.
- **`예외를 확인했다면 곧바로 통지하고 기록하는게 좋다`**.
- 문제가 발생하는 즉시 소리쳐서 잘못된 상태를 막는게 좋은 구조이다.

### 설계 질서를 파괴하는 메타 프로그래밍

- 프로그램 실행 중, **`프로그램 구조 자체를 제어하는 프로그래밍을 메타 프로그래밍이라고 부른다`**.
- 자바에서 리플렉션 API를 사용해 클래스 구조를 읽고 쓰는 메타 프로그래밍을 할 수 있다.
  - 리플렉션으로 final 변수의 값을 바꿀 수 있고 private 변수에도 접근할 수 있으며 이상 동작을 유발할 수 있다.
- 리플렉션을 남용하면 이 책에서 다룬 좋은 설계가 의미를 갖지 못할 수 있으니 시스템 분석 용도로 한정하는 등 최소화 해야 한다.

### 은 탄환

- 서양에서 늑대 인간과 악마는 은으로 만들어진 총알로 죽일 수 있다고 알려져 있다.
- 그래서 어떤 **`문제를 해결하는 비장의 무기를 은탄환이라고 부른다`**. 하지만 SW 개발에서 은 탄환은 없다.
- 중요한 것은 어떤 문제가 있을 때 어떤 방법이 해당 문제에 효과적인지, 비용이 더 들지 않는지 판단하는 자세다.
- **`설계에 Best라는 것은 없다. 항상 Better를 목표로 할 뿐이다`**.

## 이름 설계

- 이 장에서 이름을 짓는 기본적인 방법은 **`목적 중심 이름 설계이다`**.
- 이는 SW가 달성해야 하는 목적을 기반으로 이름을 설계하는 방법이다.

### 악마를 불러들이는 이름

- 온라인 쇼핑몰을 예로 들어서 `예약, 주문, 재고, 발송`을 모두 하나의 상품 클래스라고 이름을 붙이는 것이다.
- 온라인 쇼핑몰은 상품을 중심으로 출고, 예약, 주문, 발송 등 상품을 다루는 use case가 많다.
- 따라서 이름을 단순하게 상품 클래스라 붙이면 상품 클래스가 거대해지고 변경에 따른 영향 범위가 넓어지게 된다.

#### 관심사 분리

- **`결합이 느슨하고 응집도가 높은 구조로 만들려면 관심사 분리를 해야 한다`**.
- 관심사 분리는 관심사`(use case, 목적, 역할)`에 따라서 분리한다는 SW 공학의 개념이다.
- 따라서 상품 클래스는 괌심사에 따라 주문 상품, 예약 상품, 발송 상품으로 분리해야 한다.

### 이름 설계하기 - 목적 중심 이름 설계

- 저자는 클래스와 메서드에 이름을 붙이는 것을 명명이라 부르지 않고 **`이름 설계라고 부른다`**.
- **`목적 중심 이름 설계는 목적에 맞게 이름을 설계하는 것이다`**. 중요한 포인트들을 정리하자
  - 최대한 구체적이고, 의미 범위가 좁고, 특화된 이름 선택하기
  - 존재가 아니라 목적을 기반으로 하는 이름 생각하기
  - 어떤 관심사가 있는지 분석하기

#### 최대한 구체적이고, 의미 범위가 좁고, 특화된 이름 선택하기

- 목적을 달성하는데 특화된 의미 범위가 좁은 이름을 클래스에 붙인다.
- ex) `상품 -> 예약 상품, 주문 상품, 재고 상품, 발송 상품`

#### 존재가 아니라 목적을 기반으로 하는 이름 생각하기

- 목적에 특화되지 않은 경우를 생각해보자. 사람과 사용자처럼 인물이 존재하기 때문에 붙인 이름은 존재 기반 이름이다.
- 온라인 쇼핑몰에서 주소를 사용하는 목적은 배송 때문일 것이다.
  - 따라서 단순하게 주소가 아니라 발송지와 배송지처럼 목적에 특화된 이름을 사용하는게 좋다.
- 금액도 단순히 존재 기반의 이름이다.
  - 청구 금액, 소비 세액, 할인 금액 등 목적에 맞는 이름을 사용하는 것이 좋다.
- 아래 존재 기반과 목적 기반의 이름 예시를 참고하자

  ```text
    존재기반    | 목적기반
    주소       | 발송지, 배송지, 업무지
    금액       | 청구 금액, 소비 세액, 할인 금액
    사용자     |  계정, 개인 프로필, 직무
    사용자 이름  | 계정 이름, 닉네임, 본명
    상품       | 입고 상품, 예약 상품, 주문 상품, 발송 상품
  ```

#### 어떤 비지니스 목적이 있는지 분석하기

- 비지니스 목적에 특화된 이름을 만드려면 어떤 비지니스를 하는지 파악해야 한다.
  - **`온라인 쇼핑몰에는`** 판매 제품, 주문, 발송, 캠피엔 등이 있다.
  - **`게임에는`** 무기, 몬스터, 아이템, 기간, 이벤트 등이 있다.
- SW에 따라 목적과 내용이 다르다. 등장 인물과 관련되 내용을 나열해보고 관계를 정리하고 분석하자.

### 의미를 알 수 없는 이름

- 이름을 지을때 자주 부딫히는 나쁜 상황을 소개하고 해결 방법을 살펴보자.
- 우선 **`의도와 이름을 알수 없는 예시는 아래와 같다`**

  ```java
    int tmp3 = tmp1 - tmp2;
    if (tmp3 < tmp4) {
      tmp3 = tmp4;
    }
    int tmp5 = tmp3 * tmp6;
    return tmp5;
  ```

- 이름만 보고 목적이 무엇인지 알기 힘들다. 목적 중심 관점에서 보면 관심사 분리에 아무런 도움이 되지 않는다.
- 이해하기 어려우므로 관련 부분을 수정할 때마다 코드를 해석해야 한다.

#### 기술 중심 명명

- 프로그래밍, 컴퓨터와 같이 기술을 기반으로 이름 짓는 방법을 기술 중심 명명이라 부른다.
  - `ex) MemoryStateManger, ChangeIntValue01`
- 이렇게 이름을 지으면 의도를 알기 어렵기 때문에 비지니스 목적을 나타내는 이름을 짓도록 노력하자.

#### 로직 구조를 나타내는 이름

- 아래는 어떤 메서드일까?

  ```java
    class Magic {
      boolean isMemberHpMoreThanZeroAndIsMemberCanActAndIsMemberMpMoreThanMagicCostMp() {
        // 중첩 if문 로직
      }
    }
  ```

- 이는 게임에서 멤버가 마법을 사용할 수 있는 상태인지 판정하는 로직이다.
- 그런데 메서드의 이름은 로직 구조를 그대로 드러내고 있다. 무엇을 의도하는지 메서드 이름만 보고 알기 힘들다.
- **`의도와 목적을 이해하기 쉽게 이름을 붙이자`**.

  ```java
    class Magic {
      boolean canEnchant() {
        // early return + 로직
      }
    }
  ```

#### 놀람 최소화 원칙

- 다음 코드를 주문 상품 수를 리턴하는 것 처럼 보인다.

  ```java
    int count = order.itemCount();
  ```

- 그럼 `itemCount` 메서드 내부를 살펴보자

  ```java
    class Order {
      // 인스턴스 변수들

      int itemCount() {
        int count = items.count();

        if (count >= 10) {
          giftPoint = giftPoint.add(new GiftPoint(100));
        }

        return count;
      }
    }
  ```

- 놀랍게도 주문 **`상품 수를 리턴하는 로직과 기프트 포인트를 추가하는 로직이 있다`**.
- **`놀람 최소화 원칙이 있다`**. 사용자가 예상하지 못한 놀라움을 최소화하도록 설계하는 접근 방법이다.
- 처음에는 로직과 의도가 일치하게끔 구현했다고 해도, 사양을 변경하면서 별 생각없이 기존 메서드에 로직을 추가하는 경우가 있다.
  - 이는 메서드와 클래스 레벨에서 발생할 수 있으며 흔한 일이므로 주의해야 한다.
- 로직을 변경할 때는 놀람 최소화 원칙을 신경써야 한다.
  - 로직과 이름 사이에 괴리가 있다면 이름을 수정하거나, 의도에 맞게 따로 만들자.

### 구조에 악영향을 미치는 이름

#### 데이터 클래스처럼 보이는 이름

- `ProductInfo`는 상품 정보를 저장하는 클래스이다.

  ```java
    class ProductInfo {
      int id;
      String name;
      int price;
      String productCode;
    }
  ```

- `~info`, `~Data` 같은 이름은 데이터만 갖는 클래스니까 로직을 구현하면 안되는구나 생각하게 만들 수 있다.
- 데이터만 갖는다는 인상을 주는 이름은 피하는게 좋다. `ProductInfo`는 `Product`로 개선하는게 좋다.

#### 클래스를 거대하게 만드는 이름

- 클래스를 점점 더 거대하고 복잡하게 만드는 대표적인 이름으로 `Manager`가 있다.
- 문제의 원인은 Manager, 즉 관리라는 **`단어가 가진 의미가 넓고 애매하기 때문이다`**.
  - `Processor, Controller`와 같은 이름도 넓게 해석되어 거대한 클래스가 될 수 있다.
- MVC에서 `Controller`는 전달받은 요청 매개 변수를 다른 클래스에 전달하는 책무만 가져야 한다.
  - 금액을 계산하는 비지니스 로직이 들어가면 단일 책임 원칙을 위반하는 것이다.

### 이름을 봤을 때 위치가 부자연스러운 클래스

#### 동사 + 목적어 형태의 메서드 이름 주의하기

- 게임에서 적을 나타내는 `Enemy` 클래스이다. 이름을 주의깊게 살펴보자.

  ```java
    class Enemy {
      // 인스턴스 변수

      // 도망치기
      void escape() {
        ...
      }

      // MP 소비
      void consumeMagicPoint() {
        ...
      }

      // 주인공 파티에 아이템 추가하기
      boolean addItemToParty() {
        ...
      }
    }
  ```

- Enemy 클래스의 관심사는 적이다. MP를 다루는 `consumeMagicPoint`는 적의 관심사라 할 수 있다.
- 하지만 `addItemToParty` 메서드는 캐릭터의 소지품을 다루지 적의 관심사와는 상관없다.
- 다양한 환경에서 서둘러 구현하려고 기존 클래스만 가지고 끝내고자 무리하면, 관심사와 관계없는 메서드가 클래스에 추가되는 경우가 많다.
- **`동사 + 목적어로 이루어진 이름은 관계없는 책임을 가진 메서드일 가능성이 있으니`** 주의깊게 살펴보자.

#### 가능하다면 메서드의 이름은 동사 하나로 구성되게 하기

- 관심사가 다른 메서드가 섞이지 못하게 막으려면 **`메서드의 이름이 동사 하나로 설계하는게 좋다`**.
- 구체적으로 설명하면 아래와 같다

  ```text
    동사 + 목적어 형태의 메서드 ->
      목적어 개념을 나타내는 클래스 만들기 + 그 클래스에 동사 하나의 메서드 추가
  ```

- 위의 예시에서 `addItemToParty` 메서드는 일급 컬렉션을 통해 동사 하나의 메서드를 제공할 수 있다.

  ```java
    class PartyItems {
      final List<Item> items;

      // 생성자 로직 생략

      PartyItems add(final Item newItem) {
        // 예외 처리

        final List<Item> adding = new ArrayList<>(items);
        adding.add(newItem);
        return new PartyItems(adding);
      }
    }
  ```

### 이름 축약

#### 의도를 알 수 없는 축약

- 긴 이름이 싫어서 이름을 축약하는 경우가 있다. 아래 코드는 축약으로 의도를 이해하기 힘들다.

  ```java
    int trFee = brFee + LRF * dod
  ```

#### 기본적으로 이름은 축약하지 말기

- 과거에는 이름이 길면 타이핑을 많이 해야 해서 싫어하는 분위기였다.
- 하지만 최근에는 자동 완성 기능을 IDE에서 제공하므로 **`조금 귀찮더라도 축약하지 말고 쓰자`**.

  ```java
    int totalRentalFee = basicRentalFee + LATE_RENTAL_FEE_PER_DAY * daysOverDue;
  ```

- 이름 생략을 완전히 금지해야 한다는 말은 아니다. **`축약한 이름이 통용된다면 축약해도 괜찮다고 생각한다`**.
  - ex) `SNS, VIP`와 같은 관습적으로 축약해도 의미를 전달할 수 있는 경우

## 주석: 유지 보수와 변경의 정확성을 높이는 주석 작성 방법

### 내용이 낡은 주석

- **`주석의 설명과 실제 코드가 일치하지 않는 사례는 많이 찾아볼 수 있다`**. 왜 그럴까?
  - 코드에 비해 **`주석을 유지 보수 하는 것이 어렵기 떄문이다`**.
- 코드를 변경할 때 주석도 함께 변경하면 좋겠지만 업무가 바쁘고 주의하지 않으면 주석까지 유지보수 하기는 힘들다.
- 주석이 구현 시점과 멀어질수록, **`주석은 거짓말 할 가능성이 높아진다`**.

#### 로직의 동작을 설명하는 주석은 낡기 쉽다

- **`코드의 동작을 그대로 설명하는 주석은`** 코드를 변경할 때마다 주석도 변경해야 한다.
- 이처럼 로직을 그대로 설명하는 주석은 시간이 갈수록 별다른 도움이 되지 않을 수 있다.

### 주석 때문에 이름을 대충 짓는 예

- 의도를 전달하기 힘든 메서드에는 의미를 다시 설명하는 주석을 달기 쉽다.

  ```java
    class Member {
      // 수면, 혼란, 석화, 사망 이외의 상황에서 행동가능
      boolean isNotSleepingAndIsNotConfusedAndIsNotStoneAndIsNotDead() {
        ...
      }
    }
  ```

- 이런 주석은 나중에 행동 불능 상태로 `공포`가 추가될 경우 주석을 함께 변경해야 한다.
- 이런 메서드는 주석으로 설명을 추가하기 보다 메서드의 이름 자체를 수정하는게 좋다.
- **`메서드의 가독성을 높이면 주석으로 설명을 추가하지 않아도 된다`**. 그러면 낡은 주석이 생길 가능성도 줄어든다.

### 의도와 사양 변경시 주의사항을 읽는 이에게 전달하기

- 코드를 유지보수 할 때 사람의 관심사는 **`이 로직은 어떤 의도를 갖고 움직이는가`** 이다.
- 사양을 변경할 때 읽는 사람의 관심사는 **`안전하게 변경하려면 무엇을 주의해야 하는가`** 이다.
- 따라서 주석은 이러한 내용을 담는게 좋다.

  ```java
    class Member {
      // 고통받는 상황일 때 true를 리턴
      boolean isPainful() {
        // 이후 사양 변경으로 표정 변화를 일으키는 상태를 추가하면 이 메서드에 로직을 추가한다.
        if (
          state.contains(StateType.poison) ||
          state.contains(StateType.stone) ||
          state.contains(StateType.feat) ||
        ) {
          return true;
        }

        return false;
      }
    }
  ```

## 메서드: 좋은 클래스에는 좋은 메서드가 있다

### 반드시 현재 클래스의 인스턴스 변수 사용하기

- **`인스턴스 변수를 안전하게 조작하도록 메서드를 설계하면`** 클래스 내부가 정상 상태인지 보장할 수 있다.
- **`메서드는 인스턴스 변수를 사용하도록 설계하자`**. 예외도 있지만 이것이 기본 원칙이다.
- **`완전 생성자로`** 인스턴스 변수를 안전하게 만들고 다른 클래스에 인스턴스 변수를 제공한다면 **`새로운 인스턴스를 생성해 리턴하자`**.

### 불변을 활용해 예상할 수 있는 메서드를 만들기

- **`가변 인스턴스는`** 의도하지 않게 다른 부분에 영향을 줄 수 있고 유지보수 하기 어려워진다.
- **`불변을 활용해`** 예상치 못한 동작 자체를 막을 수 있게 설계하자.

### 묻지 말고 명령하라

- 인스턴스 변수 값을 추출하는 메서드를 `getter`, 값을 설정하는 메서드를 `setter` 라고 부른다.
- `getter/setter`는 다른 클래스를 확인하고 조작하는 구조가 되기 쉽기 때문에 좋지 않다.
- **`메서드를 호출하는 쪽에서는 복잡한 처리를 하지 않는게 좋다. 이때 묻지말고 명령하라 방법이 유효하다`**.

### 커맨트/쿼리 분리

- 아래 메서드는 **`상태 변경과 추출을 동시에 하고 있다`**.

  ```java
    int gainAndGetPoint() {
      point += 10;
      return point;
    }
  ```

- 상태 변경과 추출을 동시에 하는 메서드는 재 사용 하기도 힘들고 좋을게 없는 형태이다.
- **`커맨드/쿼리 분리(CQS, Command-Query Separation)`** 패턴이 있다.
  - 메서드는 커맨드 or 쿼리 중에 하나만 하도록 설계해야 한다는 패턴이다.
- 커맨드와 쿼리를 분리해보자. 쿼리가 하나의 책임만 가지며 단순해졌다.

  ```java
    int gainPoint() {
      point += 10;
    }

    int getPoint() {
      return point;
    }
  ```

### 매개 변수

#### 불변 매개변수로 만들자

- **`매개 변수를 변경하면 값을 추적해야 하고 의미를 유추하기 어렵다`**.
- 매개 변수에 final 수식자를 붙여 불변으로 만들자.

#### 플래그 매개변수를 사용하지 말자

- **`플래그 매개변수를 받는 메서드는`** 코드를 읽는 사람이 메서드가 무슨 일을 하는지 이해하기 어렵게 된다.
- 메서드 내부의 로직을 확인해야 하므로 가독성이 낮아진다.

#### null을 전달하지 말자

- null을 사용하는 로직은 `NullPointerException`이 발생할 수 있으며 null 체크가 필요해 복잡해진다.
- null 대신 초기화 상태를 `Equipment.EMPTY`로 표현했던 것처럼 구현하자.

#### 출력 매개변수를 사용하지 말자

- **`매개변수는 입력 값으로만 사용하는게 기본이다`**.
- **`매개변수를 리턴하면`** 코드를 읽는 사람에게 혼란을 줄 수 있다.

#### 매개변수는 최대한 적게 사용하자

- 매개변수가 많다는 것은 메서드가 여러 기능을 처리한다는 의미이다.

### 리턴 값

#### 자료형을 사용해서 리턴 값의 의도 나타내기

- 아래 `Price.add` 메서드는 `int` 자료형을 리턴하고 있다.

  ```java
    class Price {
      int add(final Price other) {
        return amount + other.amount;
      }
    }
  ```

- `int` 기본 자료형을 리턴하면 호출하는 쪽에 의미를 전달할 수 없다.
- 아래 코드를 보면 가격뿐 아니라 할인 금액과 배송비까지 모두 `int` 자료형을 사용하고 있다.

  ```java
    int price = productPrice.add(otherPrice);
    int discountPrice = calcDiscountedPrice(price);
    int deliveryPrice = calcDeliveryPrice(discountPrice);
  ```

- `int` 자료형을 리턴하게 만들면 어떤 값이 어떤 금액을 의미하는지 알기 힘들다.
- 따라서 **`매개변수를 잘못 전달하는 등의 실수가 발생할 수 있다`**.

  ```java
    // 배송비가 전달되어야 하는데 상품 가격이 전달되고 있음
    DeliveryCharge deliveryCharge = new Delivery(price);
  ```

- 따라서 기본 자료형을 사용하지 말고 **`값 객체를 사용해서 의도를 명확히 나타내는게 좋다`**.

  ```java
    class Price {
      Price add(final Price other) {
        return new Price(amount + other.amount);
      }
    }
  ```

#### null 리턴하지 않기

- **`매개변수로 null을 전달하지 않듯이, null을 리턴하지 않아야 좋다`**.

#### 오류는 리턴하지 말고 예외를 발생시키기

- 아래는 문제가 있는 오류 처리 코드이다.

  ```java
    class Location {
      Location shift(final int x, final int y) {
        ...

        // (-1, -1)은 오류 값
        return new Location(-1, -1);
      }
    }
  ```

- 이러한 구현은 호출하는 쪽에서 오류가 발생하면 `Location(-1, -1)`을 리턴하는 사실을 알고 있어야 한다.
- 만약 오류 처리를 잊으면 `Location(-1, -1)` 값이 후속 로직에서 정상 값처럼 사용되어 버그를 만들게 된다.
- **`잘못된 상태에서는 관용을 베풀어선 안된다. 오류 값을 리턴하지 말고 곧바로 예외를 발생시켜야 한다`**.

  ```java
    class Location {
      Location(final int x, final int y) {
        if (!valid(x, y)) {
          throw new IllegalArgumentException("불라불라");
        }
        this.x = x;
        this.y = y;
      }

      Location shift(final int x, final int y) {
        ...
      }
    }
  ```

## 리팩토링

### 리팩토링의 흐름

- 리팩토링은 **`실제 동작은 유지하면서 구조만 정리하는 작업이다`**.
- 리팩터링으로 코드를 변경할 때 실제 동작까지 바뀌어 버린다면 이는 리팩터링이라 할 수 없다.
- 아래는 리팩토링에 활용할 구매 결제를 나타내는 클래스다.

  ```java
    class PurchasePointPayment {
      final CustomerId customerId;
      final ComicId comicId;
      final PurchasePoint consumptionPoint;
      final LocalDateTime paymentDateTime;

      PurchasePointPayment(final Customer customer, final Comic comic) {
        if (customer.isEnabled()) {
          customerId = customer.id;
          if (comic.isEnabled()) {
            comicId = comic.id;
            if (comic.currentPurchasePoint.amount <= customer.possessionPoint.amount) {
              consumptionPoint = comic.currentPurchasePoint;
              paymentDateTime = LocalDateTime.now();
            } else {
              throw new RunTimeException('보유하고 있는 포인트가 부족합니다.');
            }
          } else {
            throw new IllegalArgumentException('현재 구매할 수 없는 만화입니다.');
          }
        } else {
          throw new IllegalArgumentException('유효하지 않은 계정입니다.');
        }
      }
    }
  ```

#### 중첩을 제거하여 보기 좋게 만들기

- if 조건문을 여러 번 중첩하고 있으니 `early return`을 적용해보자.

  ```java
    PurchasePointPayment(final Customer customer, final Comic comic) {
      if (!customer.isEnabled()) {
        throw new IllegalArgumentException('유효하지 않은 계정입니다.');
      }

      customerId = customer.id;
      if (!comic.isEnabled()) {
        throw new IllegalArgumentException('현재 구매할 수 없는 만화입니다.');
      }
      comicId = comic.id;
      if (comic.currentPurchasePoint.amount > customer.possessionPoint.amount) {
        throw new RunTimeException('보유하고 있는 포인트가 부족합니다.');
      }
      consumptionPoint = comic.currentPurchasePoint;
      paymentDateTime = LocalDateTime.now();
    }
  ```

#### 의미 단위로 로직 정리하기

- 결제 조건을 확인하는 로직과 변수에 대입하는 로직이 섞여 있으니 각각 분리해서 정리해준다.
- 조건 확인을 모두 완료한 이후에 값을 대입하는 순서로 바꾸자

  ```java
    PurchasePointPayment(final Customer customer, final Comic comic) {
      if (!customer.isEnabled()) {
        throw new IllegalArgumentException('유효하지 않은 계정입니다.');
      }
      if (!comic.isEnabled()) {
        throw new IllegalArgumentException('현재 구매할 수 없는 만화입니다.');
      }
      if (comic.currentPurchasePoint.amount > customer.possessionPoint.amount) {
        throw new RunTimeException('보유하고 있는 포인트가 부족합니다.');
      }

      customerId = customer.id;
      comicId = comic.id;
      consumptionPoint = comic.currentPurchasePoint;
      paymentDateTime = LocalDateTime.now();
    }
  ```

#### 조건을 읽기 쉽게 하기

- if 문에서 `!`로 부정 연산자를 사용하고 있다. 각 class에 `isDisabled` 메서드를 추가한다.

  ```java
    PurchasePointPayment(final Customer customer, final Comic comic) {
      if (customer.isDisabled()) {
        throw new IllegalArgumentException('유효하지 않은 계정입니다.');
      }
      if (comic.isDisabled()) {
        throw new IllegalArgumentException('현재 구매할 수 없는 만화입니다.');
      }
      if (comic.currentPurchasePoint.amount > customer.possessionPoint.amount) {
        throw new RunTimeException('보유하고 있는 포인트가 부족합니다.');
      }

      customerId = customer.id;
      comicId = comic.id;
      consumptionPoint = comic.currentPurchasePoint;
      paymentDateTime = LocalDateTime.now();
    }
  ```

#### 무턱대고 작성한 로직을 목적을 나타내는 메서드로 바꾸기

- 생성자에서 `amount` 비교로 포인트가 부족한지 판단하고 있다. 그런데 이 로직만 봐서는 목적을 알기 힘들다.
- **`무턱대고 로직을 작성하지 말고 목적을 나타내는 메서드로 만들어 사용하는게 좋다`**.
- `customer` 클래스에 보유 포인트가 부족한지 리턴하는 메서드를 추가한다.

  ```java
    class Customer {
      boolean isShortOfPoint(Comic comic) {
        return possessionPoint.amount < comic.currentPurchasePoint.amount;
      }
    }

    PurchasePointPayment(final Customer customer, final Comic comic) {
      if (customer.isDisabled()) {
        throw new IllegalArgumentException('유효하지 않은 계정입니다.');
      }
      if (comic.isDisabled()) {
        throw new IllegalArgumentException('현재 구매할 수 없는 만화입니다.');
      }
      if (customer.isShortOfPoint(comic)) { // 메서드로 목적 제공
        throw new RunTimeException('보유하고 있는 포인트가 부족합니다.');
      }

      customerId = customer.id;
      comicId = comic.id;
      consumptionPoint = comic.currentPurchasePoint;
      paymentDateTime = LocalDateTime.now();
    }
  ```

### 단위 테스트로 리팩터링 중 실수 방지하기

- 리팩터링시, **`실수를 줄일 수 있는 방법으로 단위 테스트가 있다`**.
- **`리팩터링을 할 때 단위 테스트는 필수다`** 라는 말이 있을 정도로 리팩터링과 단위 테스트는 세트로 이야기된다.

#### 테스트 코드를 사용한 리팩터링 흐름

- 안전하게 리팩터링하기 위한 테스트 코드 추가 방법은 여러 가지이다.
- 저자가 생각하는 방법 중 하나로는 아래와 같은 흐름이 있다.
  - 이상적인 구조의 클래스 기본 형태를 어느정도 잡는다.
  - 이 형태에서 테스트 코드를 작성한다.
  - 테스트를 실행해 실패시킨다. (로직이 완전히 작성이 안 되어 있기 떄문)
  - 테스트를 성공시키기 위한 코드를 작성한다.
  - 테스트가 성공하면 내부 코드를 리팩터링 한다.

### 불확실한 사양을 이해하기 위한 분석 방법

- 실무에선 사양을 제대로 모르는 경우도 많다. 이때는 리팩터링을 위한 테스트 코드를 작성할 수가 없다.
- **`레거시 코드 활용 전략에`** 나오는 안전하게 리팩터링 할 수 있는 방법 두 가지를 소개한다.

#### 문서화 테스트

- **`테스트를 통해 입력값에 따라 어떤 결과가 나오는지 확인하는 방법이다`**.
  - 테스트 코드에서 입력값에 따라 어떤 결과가 나오는지 확인하는 방식
- 물론 문서화 테스트만으로 사양을 완벽하게 밝히기는 어렵다. 그래도 사양의 단서를 찾는 방법 중 하나이다.

#### 스크래치 리팩터링

- **`스크래치 리팩터링은`** 정식 리팩터링이 아니라 로직의 의미와 구조를 파악하기 위해 시험삼아 리팩터링 한 것이다.
- **`테스트 코드를 작성하지 않고 로직을 먼저 리팩토링 한다`**.
- 코드가 정리되어 가독성이 좋아지면 아래의 장점이 나타난다.
  - 코드의 가독성이 좋아져 로직의 사양을 이해할 수 있게 된다.
  - 어떻게 리팩토링 할 것인지 감을 잡을 수 있고 테스트 코드를 어떻게 작성해야 할지 알 수 있다.
- 스크래치 리팩터링으로 분석한 결과로 테스트 코드를 작성해 정식으로 리팩터링 하면 된다.
- 추가로 **`스크래치 리팩터링은 분석용이므로, 리포지터리에 병합하면 안된다`**.

### 리팩터링 시 주의사항

#### 기능 추가와 리팩터링 동시에 하지 않기

- **`기능 추가와 리팩터링은 동시에 하지 말고 하나만 집중해야 한다`**.
- 리포지터리에 커밋할 때도 기능추가와 리팩터링을 구분해두지 않으면 이후에도 구분하기 힘들어진다.
  - 그럼 버그가 발생했을 때 기능 추가로 버그가 발생한 것인지 리팩터링으로 인해 발생한 것인지 분석하기 힘들어진다.

#### 작은 단계로 실시하기

- 리팩터링은 작은 단계로 커밋하는게 좋다.
- 변경이 많으면 다른 사람의 코드와 충돌할 수도 있고, 리팩터링한 코드가 불완전하면 롤백하기도 어려워진다.

## 설계의 의의와 설계를 대하는 방법

### 이 책은 어떤 설계를 주제로 집필한 것인가?

- **`설계는 어떠한 문제를 효율적으로 해결하는 구조를 만드는 것을 의미한다`**.
- 이 책에서 나오는 악마와 가장 관련 있는 품질 특성은 **`유지 보수성으로`** 볼 수 있다.
- 유지 보수성은 시스템이 정상 운용되도록 유지 보수하기가 얼마나 쉬운가를 나타내는 정도이다.

### 설계하지 않으면 개발 생산성이 저하된다

- 변경하기 어렵고 버그가 생기기 쉬운 코드를 레거시 코드라고 부른다.
- 그리고 레거시 코드가 축적되어 있는 상태를 기술 부채라고 한다.
- 변경이 용이한 설계를 하지 않으면 개발 생산성이 저하되고, 저하되는 요인으로는 크게 두 가지가 있다.

#### 버그가 발생하기 쉬운 구조

- 코드 변경시 버그가 발생하기 쉽다면 정확히 변경하는데 시간이 오래 걸린다.
- 응집도가 낮은 구조로 인해 사양 변경 시, 수정 누락이 발생하기 쉬워지고 결국 버그가 발생함

#### 가독성이 낮은 구조

- 가독성이 낮으면 의도를 이해하는데 오랜 시간이 걸린다.
- 잘못된 값이 들어와 버그가 발생했을 떄, 가독성이 낮아 출처를 추적하기 어려워진다.

#### 열심히 일했지만 생산성이 나쁨

- 개발 생산성이 낮으면 새로운 기능을 릴리즈하는데 굉장히 오래 걸린다.
- 개발 현장에선 일정을 맞추기 위해 장시간 노동하고 무작정 어떻게든 일단 작동하게 만드려고 노력한다.
- 하지만 성과를 내기 쉬운 구조를 설계하는데 노력하지 않았다면 이를 열심히 했다라고 이야기하기 어려울 것이다.
- SW 개발에선 **`나무꾼의 딜레마가 꽤 많이 발생한다`**. 제대로 설계하지 않으면 로직 변경과 디버그에 더 많은 시간을 쏟게 된다.

  ```text
    나무꾼의 딜레마

    나무꾼은 열심히 나무를 벴지만 나무가 잘 베어지지 않고 있었다.
    여행자가 지켜보다 도끼의 날이 무딘 것 같으니 도끼를 갈고 나무를 베는게 좋지 않냐 물었다.
    나무꾼은 대답했다.
    알고 있지만, 나무를 베는 것이 바빠서 도끼를 갈 시간이 없어요!
  ```

### 소프트웨어와 엔지니어의 성장 가능성

- 코드의 **`변경 용이성이 높을수록 SW가 빠르게 성장해서 가치를 높일 수 있다`**.
- SW의 성장 가능성을 높이는 것이 바로 이 책의 핵심 주제이자 의의이다.

#### 엔지니어에게 자산이란 무엇인가?

- 저자는 **`엔지니어의 자산은 기술력이라고 생각한다`**. 엔지니어의 기술력은 부를 창출하는 원천이라고 이야기 할 수 있다.
- 그런데 레거시 코드는 이러한 자산의 축적, 즉 기술력의 성장을 방해하는 무서운 존재다.

#### 레거시 코드는 발전을 막음

- 신입 사원이나 후임이 레거시 코드가 많은 프로젝트에서 개발을 담당하게 됬다.
- 신입 사원은 이 코드가 레거시 코드라는 것을 알아보기 어렵다.
- 오히려 **`선배가 작성한 좋은 코드라고 착각해서 레거시 코드를 추가로 양산하게 된다`**.
- 레거시 코드는 다음 사람들이 레거시 코드를 작성하게 만든다. 즉 낮은 수준의 기술만 사용하게 된다.

#### 레거시 코드는 고품질의 경험을 막음

- 이것이 레거시 코드임을 깨달은 후임도 있을 것이다. 그렇다면 설계부터 수정하려 노력하지만 개선하기가 매우 힘들다.
- 결국 프로젝트 일정으로 인해 설계 개선을 포기하게 된다. 고품질 설계 경험을 할 수가 없어 능력이 향상되지 않는다.

#### 레거시 코드는 시간을 낭비하게 만듬

- 레거시 코드는 이해하는데 시간이 오래 걸린다. 더 가치 있는 일에 써야 할 시간이 줄어들게 된다.
- 결국 레거시 코드는 기술 향상을 막고, 엔지니어에게 중요한 기술력의 축적을 막는다.

### 문제 해결하기

#### 문제를 인식하지 못하면 설계에 대한 생각 자체가 떠오르지 못함

- 현재 코드에 대해 문제를 인식하지 못하면 설계에 대한 생각 자체가 떠오르지 못한다.

#### 이상적인 형태를 알아야 문제를 인식할 수 있음

- 이상적인 설계를 알면 현재 설계와 비교해 기술 부채를 인식할 수 있다.

### 코드의 좋고 나쁨을 판단하는 지표

- 미래의 개발 생산성을 측정할 수 있는 방법은 현재 시점에서는 없다.
- 하지만 현재 코드의 좋고 나쁨을 판단하는 지표들이 있는데 알아보자.

#### 실행되는 코드의 줄 수

- 주석을 제외하고 실행되는 로직을 포함하는 코드의 라인 수를 의미한다.
- 라인 수가 많으면 많을수록 너무 많은 일을 하고 있는 가능성이 높다.
- 루비의 코드 분석 라이브러리는 **`적절한 코드 라인 수의 상한을 아래와 같이 정의한다`**.
  - 메서드: 10줄 이내
  - 클래스: 100줄 이내
- 저자는 `C, C++, C#, 자바스크립트, Ruby` 등의 실무 경험을 가지고 있다.
  - 저자도 대체로 위 기준으로 구현하는게 좋다고 생각한다.

#### 순환 복잡도, 응집도, 결합도

- 순환 복잡도는 코드의 구조적인 복잡함을 나타내는 지표로 조건 분기, 반복 처리, 중첩이 많아지면 복잡도가 커진다.
  - 저자가 설계할 때 클래스의 복잡도는 일반적으로 10~15 사이를 나타낸다.
- 응집도와 결합도도 분석도구로 계측할 수 있다.
  - 응집도가 높고 결합도가 낮아야 좋은 구조이다.

### 설계 대상과 비용 대비 효과

- 우리는 진행중인 프로젝트에서 코드를 모두 리팩터링하거나 차라리 다시 작성하고 싶은 마음이 들 수 있다.
- 하지만 회사의 예산은 유한하기 때문에 설계와 리팩터링을 무한히 할 수는 없다.
- 현실적으로는 프로젝트의 소스 코드 전반이 나쁘다고 해도, 설계를 개선할 수 있는 부분은 일부에 불과한 경우가 많다.
- 비용 제약이 있다면 어느 부분의 설계 품질을 높여야 하는 걸까?
  - 예를 들어 구조가 나쁘더라도 버그없이 동작하고 있으며 기능 변경이 거의 일어나지 않는 부분을 개선해야 할까?
  - 사양 변경도 없는데 비용을 들여 변경 용이성을 높이는 것은 낭비다.
- 따라서 **`비용 대비 효과가 높은 부분을 노려야 한다`**.

#### 파레토의 법칙

- **`파레토의 법칙은`** 80:20의 법칙으로 아래와 같이 비유할 수 있다.
  - 매출의 80%는 전체 상품 중 20%의 상품이 만들어 낸다
  - SW 처리 시간 중 80%는 소스 코드 전체의 20%가 차지한다
- 이 **`20에 해당하는 중요하고 사양 변경이 빈번한 곳의 설계를 개선하면 비용 대비 효과가 높을 것이다`**.

#### 코어 도메인: 서비스의 중심 영역

- 모든 상품과 서비스에는 이것이 우리가 판매하는 것. 이라고 말할 수 있는 중심 가치가 있다.
- **`서비스에서 중심이 되는 비지니스 영역을 도메인 주도 설계에서는 코어 도메인이라고 한다`**. 코어 도메인이란
  - 시스템에서 가장 큰 가치를 창출하는 곳
  - 가치있고 중요하고 비용 대비 효과가 가장 큰 곳
  - 경쟁 우위에 있고, 차별점을 만들며, 비니지스 우위를 만들 수 있는 곳
- 코어 도메인은 설계에 비용을 투자하는 곳, 비용 대비 효과가 큰 곳이라고 할 수 있다.

## 설계를 방해하는 개발 프로세스와의 싸움

- 개발 프로세스 자체가 레거시 코드의 발생 원인이 되기도 한다.
- 설계 품질을 떨어뜨리는 문제는 기술 부족 이외에도 **`심리적 안정감, 커뮤니케이션, 조직적 요인등 다양한 원인으로 발생한다`**.

### 커뮤니케이션

#### 커뮤니케이션이 부족하면 설계 품질에 문제가 발생

- 팀 개발에서 팀원과 어떤 코드를 작업할 때 서로의 로직이 맞물리지 않아 버그로 이어지는 상황이 흔하다.
- 왜 이런 현상이 일어났을까? 서로가 무엇을 하고 있는지 잘 모르기 때문이다. 왜 모를까? 커뮤니케이션이 부족하기 떄문이다.
- 바쁘다던지, 팀원 사이가 원만하지 않다던지, 정보를 바라보는 관점이 다른 것 등등 의사소통에 문제가 있으면 버그가 많아지는 경향이 있다.

#### 심리적 안정성

- 팀원간 관계 개선에는 **`심리적 안정성이 중요하다`**.
- **`심리적 안정성은 어떤 발언을 했을 때, 부끄럽거나 거절당하지 않을 것이라는 확신을 느낄 수 있는 심리 상태를 말한다`**.
  - 1999년 하버드 대학교에서 처음 소개되었고 2012년에 Google이 채용되면서 널리 알려졌다.
- 의견을 내고 제안을 하는데 팀원들이 냉소하고 제대로 귀를 기울이지 않는다면 정보 공유가 잘 이루어지지 않을 것이다.
- 커뮤니케이션에 문제가 있을 때는 일단 심리적 안정성 향상에 힘쓰는 것이 좋다.

### 설계

- 책에서 언급한 것처럼 **`설계는 굉장히 중요한 개발 프로세스이다`**. 그러나 이 설계에 다양한 함정들이 있다.

#### 빨리 끝내고 싶다는 심리가 품질 저하의 함정

- 품질이 나쁜 시스템을 만드는 팀은 클래스 설계와 관련된 습관이 애초에 없는 경우가 많다.
- **`일이 바쁘면 구현을 빨리 끝내고 싶은 마음이 앞서게 되고, 동작하기만 한다면 코드를 어떻게든 구현해 버린다`**.
- 품질을 무시하고 구현하는 과정이 반복되면 코드를 이해하는데 시간이 더 걸리고 버그가 발생하며 생산성이 점점 떨어진다.

#### 나쁜 코드를 작성하는 것이 좋은 코드를 작성하는 것보다 오래 걸린다

- 클린 아키텍처 책을 보면, TDD를 사용해 구현한 경우와 아닌 경우 어느 쪽이 더 빨리 개발을 완료하는지 실험했다.
- 결과는 TDD를 사용하는 편이 전체적으로 보았을 때 더 빠르다는 결론이 나온다.
- 이런 실험 결과를 보면 일단 어떻게든 움직이는 코드를 빨리 작성하는게 좋다라는 생각에는 동의할 수 없다.

#### 한 번에 완벽하게 설계하려 하지 말고 사이클을 돌리며 완성하기

- 사양을 대규모로 변경할 때 **`한 번에 완벽하게 설계하려는 욕심은 버리기를 권장한다`**.
- 처음부터 완벽하게 설계하려고 하면, 구현이 설계와 달리질 때 충격을 받고 설계 자체를 하지 않게 될 수도 있다.
- 단 한번의 설계로 완벽한 구조를 만들어 낼 수는 없다.
- 설계 품질은 설계와 구현 피드백 사이클을 돌리다 보면 조금씩 향상되는 것이다.

#### 성능이 떨어질 수 있으니 클래스를 작게 나누지 말자는 맞는 말일까

- 클래스 인스턴스 생성은 비용이 발생해 성능을 떨어뜨릴 수 있으므로 클래스를 많이 만들면 안 된다고 생각하는 사람들이 있다.
- 클래스가 많으면 비용이 발생하는 것은 맞지만 최근에는 무시할 수 있는 정도이다.
- 성능은 실제로 측정해 보기 전까지는 제대로 알 수 없다.
- **`병목이 어디인지 모른 채 성능이 빠른 코드를 작성하려고만 하는 것은 너뿌 빠른 최적화를 하는 안티 패턴에 해당된다`**.

#### 설계 규칙을 다수결로 결정하면 코드 품질은 떨어진다

- 팀에서 코딩 규칙과 설계 규칙을 정할 때, 팀 전체의 합의를 이루고자 규칙을 다수결로 정하는 경우가 있다.
- **`일반적으로 설계 규칙으로 다수결을 따르면 결과가 좋지 않다`**.
  - 다수결로 코드와 설계를 결정하려고 하면 아무래도 **`수준이 낮은 쪽에 맞춰서 하향평준화 되기 쉽기 때문이다`**.
- 설계 기술이 미숙한 팀원이 제안된 규칙의 좋고 나쁨을 제대로 판단할 수 있을까? 아마도 없을 것이다.

#### 설계 규칙을 정할 때 중요한 점

- **`팀원들의 능력 차이가 큰 경우에는`** 다수결보다, 시니어 엔지니어처럼 설계 역량이 뛰어난 팀원이 중심이 되는게 좋다.
- 각각의 설계 규칙에는 이유와 의도를 함께 적는 것이 좋다.
  - 규칙이 형식만 강제하고 아무런 의미를 갖지 못하는 상황을 방지하기 위해서다
- 팀의 설계 역량이 성숙하지 않으면 개인에게만 맡기지 말고, 설계 역량이 있는 팀원이 설계 리뷰와 코드 리뷰를 돕게 해서 설계 품질을 관리할 수 있다.
- 리뷰만으로는 역량을 높이기 어려운 경우가 있다.
  - 이때는 팀원들과 스터디를 진행하며 팀 전반의 설계 역량을 조금씩 높이는 것도 좋다.

### 구현

#### 깨진 유리창 이론과 보이스카우트 규칙

- 범죄학에는 **`깨진 유리창 이론이 있다`**.
  - 깨진 유리창을 방치하면 아무도 신경쓰지 않는 건물이라는 인식이 생겨 범죄가 일어난다는 의미이다.
- 이 이론은 SW 개발에도 적용된다.
  - 복잡하고 조악한 코드가 방치되면 SW 전체가 더 무질서해 진다.
  - **`저 코드를 보니까 내 코드도 그렇게 나쁜 편은 아니라는 마음의 틈이 생겨난다`**.
- 미국의 보이 스카우트에는 **`캠핑장을 떠날때 자신이 왔을 때보다 더 깨끗하게 치우고 가기라는 규칙이 있다`**.
- 이를 프로그래밍에 적용해 코드를 변경할 때 더 꺠끗한 상태로 만들어 커밋하는 것이다.
  - 이러한 작은 개선이 반복되면 코드의 전체적인 질서가 점점 더 좋아질 것이다.

#### 기존의 코드를 믿지 말고 냉정하게 파악하기

- 많은 사람들이 조악한 코드를 보고도 특별한 의심 없이 따라 하는 경우가 많다.
- 특히 신입사원들은 선배가 작성한 코드를 레거시 코드라고 생각하지 않고 같은 레거시 코드를 양산하게 된다.
- 레거시 코드를 박멸하려면 기존의 코드를 맹신하지 않는 마음가짐이 중요하다.
- 코드가 무엇을 해결하고 싶은지, 코드의 목적이 무엇인지 분석하고 이상적인 설계를 다시 생각하길 바란다.
  - 저자는 이 작업을 정체를 파악하는 행위라고 부른다.

#### 코딩 규칙 사용하기

- **`코드 작성 방식에 통일성이 없다면`** 코드를 읽기가 매우 힘들어진다.
- 코드를 읽기 쉽게 하려면 코드의 구조와 이름에 질서를 만들 수 있는 코딩 규칙을 잘 지키자.

### 리뷰

#### 코드를 설계 시점에 리뷰하기

- 많은 사람들이 코드 리뷰를 `로직이 기능을 만족하는지`, `결함이 존재하는지`, `컨벤션을 지켰는지`를 확인하는 것이라 생각한다.
- 이보다 설계적 타당성을 중심으로 리뷰해야 한다. 설계 품질은 코드에 하나하나 나타난다.

#### 존중과 예의

- 코드 리뷰 시, 기술적 올바름을 두고 공격적인 커멘트를 다는 사람이 있다.
  - 아무리 맞는 말이라도 공격적 커멘트를 허용해서는 안된다.
- 이런 리뷰는 사람에게 상처를 주고, 생산성을 저하시키는 것은 물론 코드를 좋게 만든다는 본래의 목표도 저해한다.
- **`코드 리뷰에서 가장 중요한 것은 존중과 예의다`**. 기술적 올바름보다 중요한 건 함께 일하는 동료를 존중하는 것이다.
- 아래 구글 크로미움 프로젝트의 리뷰 지침을 참고하자

#### 코드 리뷰 해야 할 지침

- 개발자는 능력이 있고 선의를 갖고 있다고 생각하자. **`실수는 정보 부족에서 발생하는 것이라 생각해야 한다`**.
- **`왜 잘못되었는지, 어떤 변경이 좋은지 설명하라`**. 이렇게 하면 안됩니다. 라는 말만으로는 의견이 상대에게 전달되지 않는다.
- 상대방의 의도를 모른다면 주저하지 말고 변경 이유를 물어보자. 의견을 교환하면 의도도 잘 알게 되고 좋은 구현을 생각할 기회가 된다.
- **`완벽을 위해 철저하게 리뷰하면, 리뷰를 받는 사람이 힘들어진다`**. 완벽을 찾지 말고 이게 좋을것 같다라는 의견으로 적절히 끝맺자
- 리뷰를 방치하지 말자. 24시간 이내로 답변할 수 없다면 언제까지 답변하겠다고 커멘트를 남겨라

#### 코드 리뷰 하지 말아야 할 지침

- 상대방이 최선을 다하고 있다는 것을 전제로 한다. **`왜 신경쓰지 못했죠?`** 같은 공격적인 말을 달아서는 안 된다.
- **`일반적인 사람이라면 이러지 않는다`** 같은 부정적인 표정을 사용해서는 안된다. **`사람이 아닌 코드에 대해서 이야기 해야 한다`**.
- 어떻게 해도 괜찮을 것 같은 경우, 리뷰에서 둘 중 하나로 결정하려 하지말자. **`리뷰의 목적은 이기고 지는 것이 아니다`**.

### 팀의 설계 능력 높이기

- 팀에서 설계를 잘 아는 팀원이 아예 없을 수도 있다.
- 이런 경우 설계를 개선하려 해도 팀원이 이해하지 못하거나 받아들이지 않아 힘든 상황들이 있다.
- 팀 전체의 설계 능력이 부족하면, 설계 능력을 높이기 위한 활동이 필요하다.

#### 영향력을 갖는 규모까지 동료 모으기

- 혼자서 품질을 높이려고 시도하는 것은 거의 효과가 나지 않는다.
- **`설계 뿐 아니라 일의 방식을 개선하려면 주위의 협력이 반드시 필요하다`**.
- 이 협력은 왜 필요할까? 서로 협력해야 일하는 방식을 바꿀 만큼 영향력이 생기기 때문이다.
- 방향성이 같은 동료에게 말을 걸고, 고민을 나눠고, 협력해 줄 동료를 만들자.

#### 천리길도 한 걸음부터

- 동료를 발견하면 책의 내용을 한번에 전달하고 싶은 충동이 생길 수 있다.
- 조급하게 굴지말자. 사람은 한 번에 대량의 정보를 받아들일 수 없고 큰 변화에 불안을 느낀다.
- 매일매일 조금씩 설계 지식을 공유해보자.

#### 백문이 불여일견

- 동료와 설계 지식을 공유했다면, 함께 클래스를 설계하고 구현한 뒤 리뷰해 보자.
- 설계도 실제로 개선 전과 후의 가독성이나 수정 범위를 비교해 보면 확실하게 알 수 있다.

#### 팔로우업 스터디 하기

- 동료를 더 모으려면 설계 스터디를 진행해 보는 것도 좋다.
- 단순히 책을 보는 것 보다는 실제 코드를 개선해 보는게 효과적이다.
- 저자가 추천하는 스터디 흐름
  - 책에 적혀 있는 노하우를 1-2개 정도 읽어 본다.
  - 프로덕션 코드에서 노하우를 적용해 볼 수 있는 부분을 찾느다. 직접 다루는 코드라면 더 좋다.
  - 노하우를 사용해 코드를 개선해 본다.
  - 어떻게 개선했는지 비포&애프터를 비교할 수 있게 발표한다.
  - 발표 내용에 대해 질의 응답과 논의를 한다.
- 스터디는 한 번에 한 시간 정도, 이렇게 하면 인풋과 아웃풋, 피드백을 빠르게 반복할 수 있어 매우 효과적이다.

#### 리더와 매니저에게 설계의 중요성과 비용 대비 효과 설명하기

- **`팀 리더와 매니저가 변경 용이성에 대한 지식이 없으면`** 개발 리소스에 변경 용이성과 관련된 설계 비용이 포함되지 않는다.
- 조직 차원에서 설계 품질을 향상하려면, 개발 프로세스 흐름에 설계를 추가해야 한다.
  - 그러러면 리더와 매니저도 설계의 필요성을 공유해야 한다.
- 매니저에게는 비용 대비 효과를 중심으로 이야기를 하자.
  - 개발 효율 저하를 이야기하고 이 문제를 해결할 수 있는 설계의 중요성을 알리자.
- **`매니저에게 설명할 때는 혼자가 아니라 신뢰를 얻는 동료와 함께 하는 것이 더 바람직하다`**.

## 설계 기술을 계속 공부하려면

### 추천 도서

- 우선 추천 도서에서 일본어 서적은 읽을 수 없기 때문에 제외했다.

#### 읽기 좋은 코드가 좋은 코드다

- 프로그래밍 세계에서 **`3일 후의 나는 타인이다`**. 라는 말이 있다.
- 자신이 직접 작성한 코드라도 3일이 지나면 의도를 잊어서 코드를 읽는데 어려움을 겪는다는 점이다.
- 이 책은 타인이나 미래의 자신이 읽더라도 의도를 이해하기 쉬운 코드를 작성하는 방법에 대해 설명한다.

#### 리팩터링 2판

- 다양한 리팩터링 기법을 알려주고 실천적인 대응 방법을 공부하기 좋은 책

#### 클린 코드

- 미숙한 나쁜 코드를 개선해서 개발 생산성이 좋은 성숙한 코드로 만드는 방법에 대해 정리한 책

#### 레거시 코드 활용 전략

- **`사양을 알 수 없고 테스트도 없는 코드를 어떻게 분석해서 리팩토링 할지`** 중점적으로 다룬 책
- 실제 현장에서 볼 수 있는 레거시 코드에 대한 대처 방법이 많이 수록되어 있다

#### 클린 아키텍처

- **`더 좋은 설계를 목표로 하려면 어떻게 해야할까?`** 라는 생각으로 이어진다.
- SOLID 원칙을 시작으로 아키텍처 전체의 변경 용이성을 향상시키기 위한 원칙, 관점, 접근방법을 설명한다.

#### 도메인 주도 설계

- 서비스에서 판매하는 것을 정의하고 이를 설장시킬 수 있는 구조를 설계하는 방법을 다룬 책

#### 도메인 주도 설계 철저 입문

- 도메인 주도 설계에 등장하는 디자인 패턴을 중심으로, 쉬운 용어를 사용해 설명하는 입문서

#### 테스트 주도 개발

- 테스트 코드를 먼저 작성하고, 코드를 작성하고, 리팩터링 하면서 코드를 발전시키는 방법
- 테스트 주도 개발을 통해서 자연스럽게 좋은 코드를 설계하는 방법을 배울 수 있는 책

### 설계 스킬을 높이는 학습 방법

- 실전 설계 능력을 높이기 위해 효율적으로 학습하는 방법을 소개한다.

#### 학습을 위한 지침

- **`인풋은 2, 아웃풋은 8`**
  - 첫번째 지침은 **`인풋보다 아웃풋을 중시하는 것이다`**. 이는 설계 뿐 아니라 학습에도 적용된다.
  - 프로그래밍 책을 읽는다고 실전 프로그래밍 스킬이 몸에 익는 것은 아니다.
    - 실제로 기능을 마주하고, 요건을 만족하는 로직을 생각하면서 코드를 작성해 보는 경험이 필요하다.
  - 설계도 책만 읽어서는 제대로 이해하기 힘들다. 코드를 보며 시행 착오를 경험해야 설계 스킬을 제대로 익힐 수 있다.
  - **`한 가지 내용을 새로 배웠다면 곧바로 코드를 작성해보며 여러 시행착오를 경험해보자`**.
- 설계 효과를 반드시 머리속에 새겨두기
  - **`설계 전 후에 효과를 반드시 확인해야 한다`**.
  - 예를 들어 전략 패턴을 사용한 뒤, 전략 패턴의 효과가 일치하는지 확인하자.
  - **`제대로 효과를 보지 못했다면 무엇이 문제인지 생각해 보는 것이 좋다`**. 이런 생각이 쌓일수록 설계에 대한 이해가 깊어질 것이다.

#### 악마의 구조를 파악하는 연습

- 책의 내용과 비교해보면서, **`평소 개발에서 다루는 프로덕션 코드를 살펴보자`**.
- 구조적으로 나쁜 부분이 어디인지, 왜 안좋은지 분석하는 연습을 하자.

#### 리팩터링으로 설계 기술력 높이기

- 설계 스킬을 높이는 가장 좋은 방법은 리팩터링이라고 생각한다.
- 우선 연습용 브랜치를 만들어서 리팩터링을 연습하자.
- **`이어서 리팩터링 대상을 선택한다`**.
  - 줄 수가 적은 private 메서드와 static 메서드가 좋은 연습 소재이다.
  - 책에서 설명한 테크닉을 활용해 여러 시도를 해보자.
- **`반복해서 언급하지만 아웃풋이 중요하다. 또 많은 연습도 중요하다. 또한 의도대로 효과가 발생하는지 확인해보자`**.
- 저자의 경험

  ```txt
    저자는 레거시 코드 때문에 계속된 야근으로 고통받는 시절이 있었다.
    어느 날 좋은 기술서라도 읽어 봐야겠다고 생각했고, 우연히 본 리팩터링이라는 책이 저자의 운명을 바꾸었다.

    책을 읽자마자 책엫서 배운 내용을 시험해 보고 싶었다.
    프로덕션 코드에 연습용 브랜치를 만들고 프로덕션 코드를 기반으로 리팩터링 연습을 매일 반복했다.
    이때 설계 스킬이 가장 많이 늘었던 것 같다.

    난잡했던 로직이 깔끔하게 정돈되어 가는게 너무 재밌어서 설계에 대한 흥미가 늘어 다양한 설계 기술서를 구매하러 돌아다녔다.
  ```

#### 동작하는 코드를 작성했다면 다시 설계하고 커밋하기

- 일단 **`제대로 동작하는 코드를 빠르게 작성하는 것을 추천한다`**.
- 시간을 들여 신중히 설계하더라도 실제 코드를 동작시키면 몇몇 요소를 빠뜨리는 경우가 많기 때문이다.
- 동작하는 코드를 구현했다고 곧바로 커밋해서는 안 된다. 그때부터 이상적인 구조를 차근차근 설계한다.
- 처음에 작성한 코드를 기반으로 설계 측면에서 좋은 클래스를 만들어 커밋한다.
  - 이렇게 코드의 품질도 높일 수 있고, 설계 스킬도 높일 수 있다.

#### 설계 기술서를 읽으며 더 높은 목표 찾기

- 계속해서 언급하지만 **`아웃풋을 내려면, 책을 읽자마자 손을 움직여 시행착오를 겪어보아야 한다`**.
- 그리고 설계 효과도 직접 확인해봐야 한다.
