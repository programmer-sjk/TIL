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

- 위 코드를 전부 모은 Money 클래스를 살펴보자

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

- 검증해보면 기존 악마가 잘 퇴치된 것을 볼 수 있다.
  - 필요한 로직이 클래스 내부에 모여 중복 코드가 작성될 일이 줄어듬
  - 중복 코드가 없으므로 수정 시, 누락이 발생할 일이 줄어듬
  - 필요한 로직이 모두 클래스에 있어 가독성이 높아짐
  - 생성자에서 변수의 값을 확정하므로 초기화되지 않은 상태가 있을 수 없음
  - 불변 변수를 사용하여 부수 효과로부터 안전함

### 프로그램 구조에 도움을 주는 디자인 패턴

- 응집도가 높은 구조, 방어적인 프로그램의 구조를 개선하는 설계 방법을 디자인 패턴이라 부른다.
- 위 Money 클래스는 사실 완전 생성자와 값 객체라는 두 가지 디자인 패턴을 적용한 것이다.

#### 완전 생성자

- 완전 생성자는 잘못된 상태로부터 클래스를 보호하기 위한 디자인 패턴이다.
- 생성자 내부에 잘못된 값이 들어오지 않도록 검사하고, 인스턴스 변수를 모두 초기화해야 객체를 생성할 수 있게 한다.
  - 이렇게 설계하면 값이 모두 정상인 완전한 객체만 만들어지게 된다.

#### 값 객체 (Value Object)

- 값 객체란 값을 클래스로 나타내는 디자인 패턴이다.
  - 애플리케이션에서 사용하는 금액, 날짜, 주문 수, 전화번호 등 다양한 값을 값 객체로 만들 수 있다.
- 값 객체의 장점은 각각의 값과 로직을 응집도가 높은 구조로 만들 수 있다.
  - 예를 들어 금액을 단순히 int 자료형으로 사용할 경우 금액 계산 로직이 여기저기 분산될 수 있다.
- 애플리케이션 내부에서 다루는 값과 개념들은 모두 값 객체로 만들 수 있다.
  - ex) 세외 제금 금액, 상품명, 전화번호, 배송지, 연령, 성별, 몸무게, 공격력, 아이템 가격, 아이템 이름 등
- 값 객체 + 완전 생성자는 객체지향 설계에서 폭 넓게 사용되는 기법이다.
- 언어와 상관없이 중요한 것은 데이터와 로직을 한 곳에 모아 응집도를 높이는 것, 필요한 조작만 외부에 공개해 캡슐화 하는 것이다.
