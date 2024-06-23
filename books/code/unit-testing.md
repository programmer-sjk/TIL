# Unit Testing

- [책 링크](https://product.kyobobook.co.kr/detail/S000001805070)

## 단위 테스트의 목표

- 단위 테스트의 목표는 **`SW 프로젝트의 지속 가능한 성장을 가능하게 하는 것이다`**.
- 흔히 단위 테스트 활동이 더 나은 설계로 이어진다고 한다. 이는 사실이다.
  - 하지만 **`단위 테스트의 주 목표는 아니다`**. 더 나은 설계는 단지 좋은 사이드 이펙트일 뿐이다.
- 아래 그림은 테스트가 없는 프로젝트의 성장 추이를 보여준다. 처음에는 빨리 시작할 수 있지만 시간이 지나면서 진척도가 많이 떨어진다.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/books/code/test-graph.png" width="350">

- 테스트 코드는 회귀에 대한 보험을 제공하기 때문에 기존 기능이 잘 동작하는지 확인하는데 도움이 된다.
- 한 가지 단점은, 테스트는 초반에 노력이 필요하다는 것이다.
- 그러나 **`프로젝트 후반에도 잘 성장할 수 있도록 하므로, 장기적으로 보면 그 비용을 메울 수 있다`**.

### 좋은 테스트와 좋지 않은 테스트를 가르는 요인

- 테스트가 잘못 작성된 프로젝트도 결국 테스트가 없는 프로젝트처럼 침체 단계에 빠진다.
  - 모든 테스트를 작성할 필요는 없다. 일부 테스트는 아주 중요하고 품질에 기여를 한다.
  - 어떤 테스트는 잘못된 경고가 발생하고, 회귀 오류를 알아내는데 도움이 되질 않으며, 유지 보수가 어렵고 느리다.
- 사람들은 종종 제품 코드와 테스트 코드가 다르다고 생각한다. **`하지만 테스트도 역시 코드다`**.
- 애플리케이션 정확성을 보장하는 것을 목표로 하는 코드베이스의 일부로 봐야 한다.

### 코드 커버리지 지표에 대한 이해

- 가장 많이 사용되는 커버리지 지표로는 코드 커버리지가 있으며, 테스트 커버리지로도 알려져있다.
- `테스트 커버리지 = 실행 코드 라인 수 / 전체 라인 수`

```c#
public static bool IsStringLong(string input)
{
  if (input.length > 5)
    return true;
  return false;
}

public void Test()
{
  bool result = IsStringLong("abc");
  Assert.Equal(false, result);
}
```

- 위 코드에서 코드 커버리지를 쉽게 계산할 수 있다.
  - 메서드 전체 라인수는 5이고, 테스트가 실행하는 라인 수는 4이다.
  - 따라서 코드 커버리지는 `4/5 = 80%`이다.
- 이제 메서드를 리팩토링 해서 아래와 같이 수정해보자.

```c#
public static bool IsStringLong(string input)
{
  return input.length > 5;
}
```

- 이제 테스트는 코드 세 줄을 모두 점검하기 때문에 커버리지가 100%로 증가했다.
- **`이 예제는 커버리지 숫자에 얼마나 쉽게 장난칠 수 있는지를 보여준다`**.

#### 특정 커버리지 숫자를 목표로 하기

- 위 설명에서 테스트 커버리지 지표만으로는 충분치 않다는 사실을 알 수 있다.
  - 100, 90, 70 같은 특정 커버리지 숫자를 목표로 삼기 시작하면 안 된다.
  - 커버리지 지표는 지표로만 봐야하지, 목표로 여겨서는 안 된다.
- **`커버리지 숫자를 강요하면`**, 개발자들은 테스트 대상에 신경쓰지 못하고, 적절한 단위테스트를 만들기 어려워진다.
- 테스트 커버리지 숫자가 낮으면 문제라 할 수 있다. 하지만 높은 숫자도 별 의미는 없다.

### 요약

- 코드는 점점 나빠지고, 시스템은 점점 복잡해지며 흐트러진다. 테스트를 통해 회귀에 대한 보험을 제공할 수 있다.
- 단위 테스트의 목표는 프로젝트가 지속적으로 성장하게 하는 것이다.
- 특정 커버리지 숫자를 부과하면 동기 부여가 잘못된 것이다. 시스템 핵심 부분에 높은 커버리지를 갖는게 좋지만 이를 목표로 하지 않는다.

## 단위 테스트란 무엇인가

- **`단위 테스트에 접근하는 방법이 고전파, 런던파 두 가지 견해로 나뉘었다`**.
- 고전파는 테스트에 대해 원론적으로 접근하는 방식이기에 고전이라 하고, 런던파는 런던의 프로그래밍 커뮤니티에서 시작되었다.

### 단위 테스트의 정의

- 단위 테스트는 가장 중요한 세 가지 속성이 있다.
  - 작은 코드 조각을 검증하고
  - 빠르게 수행하고
  - 격리된 방식으로 처리하는 자동화된 테스트다.
- 여기서 **`고전파와 런던파의 의견이 다른 점은 세 번째 격리 문제이다`**.

#### 격리 문제에 대한 런던파의 접근

- 런던파에서는 하나의 클래스가 **`여러 클래스에 의존하면, 이 모든 의존성을 테스트 대역으로 대체해야 한다`**.
- 이런식으로 동작을 외부 영향과 분리해서 테스트 대상 클래스에만 집중할 수 있다.
- 이 방법의 한 가지 이점은 테스트가 실패하면 확실히 테스트 대상 시스템이 고장난 것이다.
  - 클래스의 모든 의존성은 테스트 대역으로 대체됐기 때문에 의심할 여지가 없다.
- 고전적인 스타일이 사람들에게 더 익숙하기 때문에, 고전적인 스타일로 작성된 테스트 코드를 보고, 런던 스타일로 다시 작성해보자.

```c#
// 고전적인 스타일로 작성된 테스트
public void Purchase_succeeds_when_enough_inventory()
{
 // given
 var store = new Store();
 store.addInventory(Product.Shampoo, 10);
 var customer = new Customer();

 // when
 bool success = customer.Purchase(store, Product.Shampoo, 5);

 // then
 Assert.True(success);
 Assert.Equal(5, store.GetInventory(Product.Shampoo));
}
```

- 위 예시에서 `given` 절은 의존성과 테스트 대상 시스템을 모두 준비하는 부분이다.
  - `테스트 대상 시스템(SUT)`과 협력자를 준비한다.
  - 이 경우 customer이 SUT가 되고, store이 협력자가 된다.
- **`위 코드는 고전 스타일의 예시로, 테스트는 협력자를 대체하지 않고 운영용 인스턴스를 사용한다`**.
- `Customer, Store` 모두 검증한다. **`그러나 Customer가 올바르게 작동해도 Store 내부에 버그가 있으면 단위 테스트에 실패할 수 있다. 테스트에서 두 클래스는 서로 격리돼 있지 않다`**.
- 이제 런던 스타일로 예제를 수정해보자.

```c#
// 런던 스타일로 작성된 테스트
public void Purchase_succeeds_when_enough_inventory()
{
 // given
 var storeMock = new Mock<IStore>();
 storeMock.Setup(x => x.HasEnoughInventory(Product.Shampoo, 5)).Returns(true);
 var customer = new Customer();

 // when
 bool success = customer.Purchase(storeMock.Object, Product.Shampoo, 5);

 // then
 Assert.True(success);
 storeMock.Verify(x => x.RemoveInventory(Product.Shampoo, 5), Times.Once)
}
```

- 런던 스타일에서는 `given` 절에서 Store의 실제 인스턴스 대신 `Mock<T>`를 사용해 대체한다.
- 검증 단계에서 고전파는 상점의 상태를 검증했지만 런던파에서는 `Customer <-> Store` 간 상호 작용을 검사한다.
  - 즉 고객이 상점으로 호출해야 할 메서드와 호출 횟수까지 검증할 수 있다.

#### 격리 문제에 대한 고전파의 접근

- 고전적인 방법에서 코드를 꼭 격리하는 방식으로 테스트해야 하는 것은 아니다.
- 대신 **`단위 테스틑 서로 격리해서 실행해야 한다`**.
  - 이렇게 하면 테스트를 어떤 순서로 실행하든 서로의 결과에 영향을 미치지 않는다.
  - 테스트를 격리하는 것은 공유 상태에 도달하지 않는 한, 여러 클래스를 한 번에 테스트 할 수 있다.
- **`공유 의존성은`** 테스트 간에 공유되고, **`서로의 결과에 영향을 미칠 수 있는 의존성이다`**.
  - **`DB, 파일 시스템`** 등의 외부 의존성이 공유 상태의 대표적인 예시이다.
  - 테스트 준비 단계에서 DB의 고객을 생성할 수 있고, 다른 테스트에서 고객을 삭제할 수도 있다.
    - 이 두가지 테스트를 병렬로 실행하면 테스트가 실패할 수 있다.
- **`공유 의존성을 대체하는 또 다른 이유는 테스트 실행 속도를 높이는데 있다`**.
  - DB, 파일 시스템 등의 공유 의존성 호출은 비공개 의존성에 비해 더 오래 걸린다.

### 단위 테스트의 런던파와 고전파

- **`런던파와 고전파로 나눠진 원인은 격리 특성에 있다`**.
- 런던파는 협력자도 격리하는 반면, 고전파는 단위 테스트끼리 격리하는 것으로 본다.

#### 고전파와 런던파가 의존성을 다루는 방법

- 의존성은 불변 의존성, 비 공개 의존성도 존재한다.
- 고전파에서는 공유 의존성만 교체 대상이라면, 런던파에서는 추가적으로 변경 가능한 의존성도 교체 대상이다.
- 런던파에서는 변경 가능한 비 공개 의존성도 테스트 대역으로 교체할 수 있다.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/books/code/dependency.png" width="500">

### 고전파와 런던파의 비교

- **`고전파와 런던파의 차이는`** 단위 테스트의 정의에서 **`격리 문제를 어떻게 다루는지에 있다`**.
- **`저자 개인적으로는 고전파가 고품질의 테스트를 만들고 지속 가능한 성장을 달성하는 데 더 적합하다고 생각한다`**.
- **`그 이유는 취약성에 있다`**. Mock을 사용하는 테스트는 고전적인 테스트에 비해 불안정한 경향이 있기 떄문이다.
- 런던파의 장점을 취합하면 아래와 같다.
  - 테스트가 세밀해서 한 번에 한 클래스만 확인한다.
  - 테스트가 실패하면 어떤 기능이 실패했는지 확실히 알 수 있다. 테스트 내 다른 의존성을 제거했기 때문에 SUT에 포함된 버그만 실패한다.

### 두 분파의 통합 테스트

- 런던파와 고전파는 통합 테스트의 정의에도 차이가 있다.
- **`런던파는 실제 협력자 객체를 사용하는 모든 테스트를 통합 테스트로 간주한다`**.
  - 고전 스타일로 작성된 대부분의 테스트는 런던파에게는 통합 테스트로 느껴질 것이다.
- 이 책은 고전적인 정의로 단위 테스트와 통합 테스트를 정의한다.
- 고전파의 관점에서 단위 테스트를 정의해보자.
  - 단일 동작을 검증하고
  - 빠르게 수행하고
  - 다른 테스트와 격리해서 처리한다.
- **`통합 테스트는 이러한 기준 중 하나를 충족하지 않는 테스트이다`**.
  - 예를 들어 **`DB(공유 의존성)에 접근하는 테스트는`** 다른 테스트와 격리해 실행할 수 없다.
  - 또 외부 의존성에 접근하면 테스트가 느려진다. DB 호출은 처음에는 미미하지만 테스트 코드가 커질수록 시간이 지연된다.
  - 둘 이상의 동작 단위를 검증할 때의 테스트도 통합 테스트이다.

#### 통합 테스트의 일부인 e2e(end-to-end) 테스트

- 위에서 통합 테스트는 공유 의존성, 프로세스 외부 의존성, 다른 팀이 개발한 코드와 통합해 작동하는지도 검증하는 테스트다.
- **`e2e는 통합 테스트의 일부다`**. e2e는 최종 사용자의 관점에서 검증하며 통합 테스트에 비해 의존성을 더 많이 포함한다.
- e2e는 유지 보수 측면에서 비용이 많이 들기 때문에 `단위 테스트 -> 통합 테스트`를 통과한 후 마지막에 실행하는 것이 좋다.

### 요약

- **`테스트 대상 시스템(SUT)의`** 의존성 처리 방식에 따라 고전파와 런던파로 나뉠 수 있다.
- 런던파 테스트의 가장 큰 문제는 SUT 세부 구현에 결합된 테스트 문제다.
- 통합 테스트는 단위 테스트의 기준 중 하나 이상을 충족하지 못한 테스트이다.

## 단위 테스트 구조

### 단위 테스트를 구성하는 방법

- 이 절에서는 **`단위 테스트를 구성하는 방법, 피해야 할 함정, 테스트를 읽기 쉽게 만드는 방법들을`** 알아본다.

#### AAA 패턴 사용

- **`AAA 패턴은`** `Given-When-Then` 패턴과 같이 테스트를 준비, 실행, 검증이라는 세 부분으로 나눌 수 있다.

```c#
public class CalculatorTests // 응집도 있는 테스트 세트를 위한 클래스 컨테이너
{
 [Fact] // 테스트를 나타내는 xUnit 속성
 public void Sum_of_two_numbers()
 {
  // 준비
  double first = 10;
  double second = 20;
  var calculator = new Calculator();

  // 실행
  double result = calculator.Sum(first, second);

  // 검증
  Assert.Equal(30, result);
 }
}
```

- `Given-When-Then` 패턴과 같이 AAA 패턴은 모든 테스트가 단순하고 동일한 구조를 갖는데 도움이 된다.
- 익숙해지면 테스트를 쉽게 읽고 이해할 수 있으며, 테스트 유지 보수 비용이 줄어든다.

#### 여러 개의 준비, 실행, 검증 구절 피하기

- 여러 개의 준비, 실행, 검증은 **`테스트가 너무 많은 것을 한 번에 검증한다는 의미다`**.
- 여러 실행 구절을 보면, 여러 동작을 테스트 한다는 의미이므로 각 동작을 고유의 테스트로 도출해라.
- 통합 테스트에서는 속도를 높이기 위해 여러 실행 구절을 두는게 선택지일 수 있다.
- 그러나 최적화 기법은 더 느려지게 하고 싶지 않은 통합 테스트에만 적용할 수 있다. 단위 테스트에서는 여러개의 테스트로 나누는게 좋다.

#### 테스트 내 if 문 피하기

- **`if 문이 있는 단위 테스트는 안티 패턴이다`**. 모든 테스트는 분기가 없는 간단한 일련의 단계여야 한다.
- if 문은 테스트가 한 번에 너무 많은 것을 검증한다는 것을 뜻하며 if 문은 테스트를 읽고 이해하는 것을 더 어렵게 만든다.

#### 각 구절은 얼마나 커야 하는가?

- **`일반적으로 준비 구절이 가장 크다`**. 하지만 너무 길면 같은 **`테스트 클래스 내 비공개 메서드 or 별도의 팩토리 클래스로 도출하는게 좋다`**.
- 준비 구절에서 코드 재사용에 도움이 되는 패턴으로 **`오브젝트 마더와 테스트 데이터 빌더가 있다`**.
- **`실행 구절은 보통 코드 한 줄이다`**. 두 줄 이상인 경우 공개 API의 캡슐화에 문제가 있을 수 있다.
- 실행 구절을 한 줄로 하는 지침은 비지니스 로직을 포함하는 대부분의 코드에 적용되지만 유틸리티나 인프라 코드는 덜 적용되기에 절대라고 표현할 순 없다.

#### 검증 구절에서 검증문이 얼마나 있어야 하는가

- 단위 테스트에서 테스트하는 동작은 여러 결과를 낼 수 있으며, 하나의 테스트로 그 모든 결과를 평가하는 것이 좋다.
- 일반적으로 검증 구절이 커지는 것을 경계해야 한다. 결과 객체의 모든 속성을 검증하는 대신 equal로 단일 검증을 할 수 있다.

#### 테스트 대상 시스템 구별하기

- SUT는 테스트에서 중요한 역할을 하는데, 어플리케이션에서 호출하고자 하는 동작에 대한 진입점을 제공한다.
- SUT가 많은 경우, 테스트 대상을 쉽게 찾기 위해 테스트 코드에서 sut로 지정할 수 있다.

```c#
public class CalculatorTests
{
 [Fact]
 public void Sum_of_two_numbers()
 {
  // 준비
  double first = 10;
  double second = 20;
  var sut = new Calculator();

  // 실행
  double result = sut.Sum(first, second);

  // 검증
  Assert.Equal(30, result);
 }
}
```

### 테스트 간 테스트 픽스처 재사용

- 테스트를 준비하기 위한 given 절에 너무 많은 코드를 작성해야 할 때가 있다.
- 이런 경우 별도의 메서드나 클래스로 도출한 후 테스트 간에 재사용하는 것이 좋다.
- 테스트 픽스처를 재사용하는 잘못된 방법은 테스트 생성자에서 픽스처를 초기화 하는 것이다.

```c#
public class CustomerTests
{
 private readonly Store _store; // 공통 테스트 픽스처
 private readonly Customer _sut;

 public CustomerTests()
 {
  // 클래스 내 각 테스트 이전에 호출
  _store = new Store();
  _store.AddInventory(Product.Shampoo, 10);
  _sut = new Customer();
 }

 [Fact]
 public void Purchase_succeeds_when_enough_inventory()
 {
  bool success = _sut.Purchase(_store, Product.Shampoo, 5);

  Assert.True(success);
  Assert.Equals(5, _store.GetInventory(Product.Shampoo));
 }

 [Fact]
 public void Purchase_fails_when_not_enough_inventory()
 {
  bool success = _sut.Purchase(_store, Product.Shampoo, 15);

  Assert.False(success);
  Assert.Equals(10, _store.GetInventory(Product.Shampoo));
 }
}
```

- 위와 같이 생성자에서 데이터를 준비하면 두 가지 중요한 단점이 있다.
- 테스트 간 결합도가 높아지고 가독성이 떨어진다.

#### 테스트 간의 높은 결합도는 안티 패턴이다

- 위 예시에선 모든 테스트가 서로 결합돼 있어서 테스트의 준비 로직을 수정하면 클래스의 모든 테스트에 영향을 미친다.
- 테스트를 수정해도 다른 테스트에 영향을 주어서는 안 된다.

#### 테스트 가독성을 떨어뜨리는 생성자 사용

- 테스트 코드만 보고는 전체 그림을 볼 수 없다.
- 테스트가 무엇을 하는지 이해하려면 클래스의 다른 부분도 봐야 한다.

#### 더 나은 테스트 픽스처 재사용법

- 생성자보다 더 나은 방법은 비공개 팩토리 메서드를 두는 것이다.

```c#
public class CustomerTests
{
 [Fact]
 public void Purchase_succeeds_when_enough_inventory()
 {
  Store store = CreateStoreWithInventory(Product.Shampoo, 10);
  Customer sut = CreateCustomer();

  bool success = _sut.Purchase(_store, Product.Shampoo, 5);

  Assert.True(success);
  Assert.Equals(5, _store.GetInventory(Product.Shampoo));
 }

 [Fact]
 public void Purchase_fails_when_not_enough_inventory()
 {
  Store store = CreateStoreWithInventory(Product.Shampoo, 10);
  Customer sut = CreateCustomer();

  bool success = _sut.Purchase(_store, Product.Shampoo, 15);

  Assert.False(success);
  Assert.Equals(10, _store.GetInventory(Product.Shampoo));
 }

 private Store CreateStoreWithInventory(Product product, int quantity)
 {
  Store store = new Store();
  store.AddInventory(product, quantity);
  return store;
 }

 private static Customer CreateCustomer()
 {
  return new Customer();
 }
}
```

- 공통 초기화 코드를 비공개 메서드로 추출해 테스트 코드를 짧게 하면서, 테스트 전체 맥락을 유지할 수 있다.
- 비공개 메서드는 테스트간 서로 결합되지 않고, 읽기 쉬우며 재사용이 가능하다.
