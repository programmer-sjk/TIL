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

- 테스트를 준비하기 위한 `given` 절에 너무 많은 코드를 작성해야 할 때가 있다.
- 이런 경우 **`별도의 메서드나 클래스로 도출한 후 테스트 간에 재사용하는 것이 좋다`**.
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

- 위와 같이 생성자에서 데이터를 준비하면 **`두 가지 중요한 단점이 있다`**.
- **`테스트 간 결합도가 높아지고 가독성이 떨어진다`**.

#### 테스트 간의 높은 결합도는 안티 패턴이다

- 위 예시에선 모든 테스트가 서로 결합돼 있어서 테스트의 준비 로직을 수정하면 클래스의 모든 테스트에 영향을 미친다.
- **`테스트를 수정해도 다른 테스트에 영향을 주어서는 안 된다`**.

#### 테스트 가독성을 떨어뜨리는 생성자 사용

- 테스트 코드만 보고는 전체 그림을 볼 수 없다.
- 테스트가 무엇을 하는지 이해하려면 클래스의 다른 부분도 봐야 한다.

#### 더 나은 테스트 픽스처 재사용법

- 생성자보다 더 나은 방법은 **`비공개 팩토리 메서드를 두는 것이다`**.

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

### 단위 테스트 명명법

- 테스트에 표현력이 있는 이름을 붙이는 것이 중요하다.
- 표현력 있고 **`읽기 쉬운 테스트 이름을 지으려면 다음 지침을 따르자`**.
  - 복잡한 동작에 대한 높은 수준의 설명은 엄격한 명명 정책에 넣기 힘들다. 표현의 자유를 허용하자.
  - 비개발자들에게 시나리오를 설명하는 것처럼 테스트 이름을 짓자.
- 하나의 예시로 `IsDelivery_InvalidDate_ReturnsFalse()` 테스트 이름이 있다면
  - `Delivery_with_a_past_date_is_invalid()` 라는 테스트 이름이 훨씬 낫다.

### 매개변수화된 테스트 리팩터링하기

- 테스트 코드의 양을 줄이고자 테스트를 묶을 수 있다.

```c#
public class DeliveryServiceTests
{
 [InlineData(-1, false)]
 [InlineData(0, false)]
 [InlineData(1, false)]
 [InlineData(2, true)]
 [Theory]
 public void Can_detect_an_invalid_delivery_date(
  int daysFromNow,
  bool expected
 )
 {
  ...
 }
}
```

- 매개변수화된 테스트를 사용하면 테스트 코드의 양을 줄일 수 있지만 내용을 파악하기가 어려워졌다.
- 절충안으로는 긍정적인 테스트 케이스는 고유한 테스트로 도출하고 좋은 이름을 짓는 것이다.

```c#
public class DeliveryServiceTests
{
 [InlineData(-1, false)]
 [InlineData(0, false)]
 [InlineData(1, false)]
 [Theory]
 public void Can_detect_an_invalid_delivery_date(
  int daysFromNow,
  bool expected
 )
 {
  ...
 }

 [Fact]
 public void The_soonest_delivery_date_is_two_days_from_now()
 {
  ...
 }
}
```

- 입력 매개변수만으로 테스트 케이스를 판단할 수 있다면 긍정과 부정 테스트 모두 하나의 메서드로 두는 것이 좋다.
- 테스트 파악이 어렵다면 긍정적인 테스트 케이스를 도출하자.
- 그럼에도 동작이 너무 복잡하다면 매개변수화된 테스트를 사용하지 말고, 각각의 테스트 메서드로 나누자.

### 요약

- 모든 단위 테스트는 `AAA 패턴(준비, 실행, 검증)`을 따라야 한다.
- 실행 구절이 한 줄 이상이면 SUT의 API에 문제가 있다는 뜻이다.
- 테스트 픽스처 초기화 코드는 생성자에 두지 말고 비공개 메서드나 팩토리 메서드를 도입해 재사용하자.

## 좋은 단위 테스트의 4대 요소

### 좋은 단위 테스트의 4대 요소 살펴보기

- 회귀 방지
- 리팩터링 내성
- 빠른 피드백
- 유지 보수성

#### 회귀 방지

- 일반적으로 **`테스트가 실행하는 코드가 많을수록 테스트에서 회귀가 나타날 가능성이 높다`**.
- 복잡한 비지니스 로직을 검증하는 테스트는 가치 있지만, 단순한 코드를 테스트하는 것은 가치가 거의 없다.
  - 이런 코드는 짧고, 비니지스 로직을 거의 담지 않기 떄문이다.

#### 리팩터링 내성

- 리팩터링을 통해 **`기능은 의도한대로 동작하지만 테스트가 실패하는 것을 거짓양성이라고 한다`**.
- **`거짓양성은 테스트를 통해 얻을 수 있는 이점을 방해한다`**.
  - 테스트가 타당한 이유 없이 실패하면, 시간이 흐르면서 실패에 익숙해지고 신경을 많이 쓰지 않는다.
  - **`거짓 양성이 빈번하면 테스트에 대한 신뢰가 떨어지며, 리팩터링이 줄어든다`**. 회귀를 피하려고 코드 변경을 최소한으로 하기 때문이다.

#### 거짓양성의 원인은 무엇인가

- 거짓 양성은 테스트 구성 방식과 관련이 있다. 테스트가 SUT의 **`구현 세부 사항에 많이 결합할수록 거짓양성은 늘어난다`**.
- 거짓 양성을 줄이는 방법은 구현 세부 사항에서 테스트를 분리하는 것 뿐이다.
- 구현 세부 사항에서 멀어지기 위해서는 **`최종 결과를 목표로 테스트해야 한다`**.

### 회귀 방지와 리팩터링 내성

- **`좋은 단위 테스트의 두 요소(회귀 방지와 리팩터링 내성)는`** 본질적으로 관계가 있다.
- 프로젝트가 시작된 직후에는 회귀 방지가 중요한 것에 비해 리팩터링 내성은 상대적으로 중요하지 않다.
  - 프로젝트 초반에는 리팩터링이 크게 중요하지 않으며 시간이 지나면서 점차 중요해진다.
- 시간이 흐를수록 코드베이스는 나빠지고 복잡하므로 정기적으로 리팩터링이 필요하다.
  - 결국 테스트에서 리팩터링 내성도 점점 더 중요해진다.
- 테스트에서 계속 늑대라고 울리면 리팩터링을 할 수 없고, 존재하지 않는 버그에 경고를 받게 되니 신뢰를 잃게 된다.

### 빠른 피드백과 유지 보수성

- 빠른 피드백은 단위 테스트의 필수 속성이다. **`테스트 속도가 빠를수록 더 많은 테스트를 수행할 수 있고 자주 실행할 수 있다`**.
- 느린 테스트는 자주 실행하지 못하기 때문에 피드백을 느리게 하고 시간을 더 많이 낭비하게 된다.
- 마지막 유지 보수성은 테스트를 얼마나 이해하기 쉽고 얼마나 실행하기 쉬운가와 관련이 있다.

### 이상적인 테스트

- **`좋은 단위 테스트의 4대 요소를 모두 만족하는 이상적인 테스트를 만드는 것은 불가능하다`**.
- 회귀 방지, 리팩터링 내성, 빠른 피드백은 상호 배타적이기 때문이다.
  - 엔드 투 엔드 테스트는 회귀 방지와 리팩터링 내성은 강하지만 빠른 피드백을 받기 어렵다.
  - 버그가 없는 간단한 테스트는 리팩터링 내성과 빠른 피드백은 가능하지만, 실수할 여지가 없기에 회귀를 나타내진 않는다.
  - 깨지기 쉬운 테스트는 회귀 방지와 빠른 피드백이 가능하지만 리팩터링 내성은 낮다.
- 회귀 방지, 리팩터링 내성, 빠른 피드백 중 좋은 테스트를 만드는 균형을 만드는 것은 쉽지 않다.
  - 각각 일부를 희생해야 하지만 **`실제론 리팩터링 내성을 포기할 수 없어서 회귀방지와 빠른 피드백 사이에서 절충해야 한다`**.
  - 리팩터링 내성을 포기할 수 없는 이유는 내성이 있거나 없거나의 문제고 중간 단계가 없기 때문이다.
  - 반면에 **`회귀 방지와 빠른 피드백에 대한 지표는 중간 단계에서 조절이 가능하다`**.

### 대중적인 테스트 자동화 개념 살펴보기

#### 테스트 피라미드

- **`테스트 피라미드는`** 테스트 유형 간의 일정한 비율을 일컫는 개념이다.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/books/code/test-pyramid.png" width="250">

- **`피라미드에서 넓을수록 테스트는 많아지며`** 위에 있을수록 사용자의 동작을 유사하게 흉내내는 테스트다.
- 피라미드 상단의 테스트는 회귀 방지에 유리한 반면 하단은 실행 속도를 강조한다.
- 피라미드에서 테스트는 빠른 피드백과 회귀 방지 사이에서 선택하며, **`어떤 계층도 리팩터링 내성을 포기하지 않는다`**.
- e2e 테스트는 빠른 피드백 관점에서 낮은 점수를 받기 때문에 가장 중요한 기능에만 적용하거나 긍정적인 케이스에만 적용한다.

#### 블랙박스 테스트와 화이트박스 테스트

- **`블랙박스 테스트는 시스템 내부 구조를 모르는 상태에서`** 기능을 검사하는 테스트 방법이다.
- 화이트박스 테스트는 정반대로 내부 작업을 검증하는 테스트 방식이다.
- 화이트박스 테스트는 더 철저하게 테스트하지만 리팩터링 내성이 나쁘다. 블랙박스 테스트는 정반대의 장단점을 제공한다.
- 앞에서 언급했듯, **`리팩터링 내성은 타협할 수 없기에 블랙박스 테스트를 기본으로 선택하자`**.

### 요약

- 좋은 단위 테스트는 **`회귀 방지, 리팩터링 내성, 빠른 피드백, 유지 보수성까지`** 네 가지 특성이 있다.
- 거짓 양성은 허위 경보다. 허위 경보에 익숙해져서 주의를 기울이지 않게 되고 테스트에 대한 신뢰를 잃게 된다.
- 거짓 양성은 테스트와 SUT의 구현 세부 사항에 결합도가 강하기 때문에 발생한다. 따라서 SUT가 만든 최종 결과를 검증해야 한다.
- 유지 보수성은 테스트의 이해와 실행 난이도로 결정된다.
- 리팩터링 내성은 모 아니면 도이기에 타협할 수 없다. 절충은 회귀방지와 빠른 피드백 사이의 선택으로 귀결된다.
- 테스트 피라미드는 단위/통합/e2e 테스트의 일정한 비율을 일컫는다.

## 목과 테스트 취약성

- 테스트에서 목을 사용하는 것은 논란의 여지가 있는 주제다. 목과 테스트 취약성 사이에는 깊고 불가피한 관련이 있다.

### 목과 스텁 구분

- 목은 테스트 대상 시스템(SUT)과 협력자 사이의 상호 작용을 검사할 수 있는 테스트 대역이라고 했다.
- 또 다른 테스트 유형으로 스텁(stub)이 있다. 목과 스텁이 어떻게 다른지 알아보자.

#### 테스트 대역 유형

- 테스트 대역은 모든 가짜 의존성을 설명하는 포괄적인 용어다.
- **`테스트 대역의 주 용도는 테스트를 편리하게 하는 것이다`**.
- 테스트 대역에는 **`더미, 스텁, 스파이, 목, 페이크`** 다섯가지가 있지만 **`실제로는 목과 스텁 두 유형으로 나눌 수 있다`**.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/books/code/test-double.png" width="500">

- **`목은 외부로 나가는 상호 작용을 모방하고 검사하는데`** 도움이 된다.
- 반면 **`스텁은 내부로 들어오는 상호 작용을 모방하는데`** 도움이 된다.

 <img src="https://github.com/programmer-sjk/TIL/blob/main/images/books/code/mock-vs-stub.png" width="500">

- 크게 목과 스텁으로 나뉘며 나머지는 미미한 구현 사항의 차이다.
  - 스파이는 목과 같은 역할을 하지만, 스파이는 수동으로 작성하는 반면, 목은 목 프레임워크의 도움을 받는다.
  - 스텁, 더미, 페이크의 차이는 얼마나 똑똑한지에 있다.
    - 더미는 단순히 하드코딩된 값이고, 스텁은 시나리오마다 다른 값을 반환하게끔 필요한 것을 다 갖춘 완전한 의존성이다.
    - 페이크는 대부분 스텁과 같으나 아직 존재하지 않는 의존성을 대체하고자 구현한다.
- 목은 SUT와 관련 의존성 간의 상호 작용을 모방하고 검사하는 반면, 스텁은 모방만한다. 이는 중요한 차이점이다.

#### 도구로서의 목과 테스트 대역으로서의 목

- **`목이라는 용어는`** 목 라이브러리에 있는 Mock 클래스와, 테스트 대역으로서의 목이 있다.

```c#
[Fact]
public void Sending_a_greetings_email()
{
 var mock = new Mock<IEmailGateway>(); // Mock(도구)으로 mock(목) 생성
 var sut = new Controller(mock.Object);

 sut.GreetUser("user@email.com");

 // 테스트 대역으로 SUT의 호출을 검사
 mock.Verify(x => x.SendGreetingsEmail("user@email.com"), Times.Once);
}
```

- 위 예제에서 **`Mock 클래스는 도구로서의 목에 비해, 인스턴스인 mock은 테스트 대역으로서의 목이다`**.
- 도구로서의 목을 사용해 목과 스텁. 두 가지 유형의 테스트 대역을 생성할 수 있기 때문에 혼동하지 않는 것이 중요하다.
- 아래 예제도 Mock 클래스를 사용하지만 해당 클래스의 인스턴스는 목이 아니라 스텁이다.

```c#
[Fact]
public void Creating_a_report()
{
 var stub = new Mock<IDatabase>(); // Mock(도구)으로 stub(스텁) 생성
 stub.Setup(x => x.GetNumberOfUsers()).Returns(10); // 준비한 응답 설정
 var sut = new Controller(stub.Object);

 Report report = sut.CreateReport();

 Assert.Equal(10, report.NumberOfUsers);
}
```

- 위 예제에서 테스트 대역 스텁은 내부로 들어오는 상호 작용(SUT에 입력 데이터를 제공하는 호출)을 모방한다.
- 반면 이전 예제(`SendGreetingsEmail`)에서 목은 외부로 나가는 상호 작용이고 목적은 사이드 이펙트(이메일 발송)뿐이다.

#### 스텁으로 상호 작용을 검증하지 말라

- 목은 SUT에서 의존성으로 나가는 상호 작용을 모방하고 검사한다.
- 스텁은 내부로 들어오는 상호 작용만 모방하고 검사하지 않는다.
- **`스텁과의 상호 작용을 검증하는 것은 취약한 테스트를 야기하는 일반적인 안티 패턴이다`**.
- 밖으로 나가는 의존성에 대해 아래 코드는 실제 결과에 부합하며, 도메인 전문가에게 의미가 있다.
  - `mock.Verify(x => x.SendGreetingsEmail("user@email.com"), Times.Once);`
  - 즉, 인사 메일을 보내는 것은 비지니스 담당자가 시스템에 하길 원하는 것이다.
- 스텁에서 `GetNumberOfUsers`를 검증하는 것은 결과가 아니고, 입력을 위한 내부 구현 세부사항이다.

#### 목과 스텁 함께 쓰기

- 떄로는 목과 스텁을 모두 나타내는 테스트 대역을 만들 필요가 있다.

```c#
[Fact]
public void Purchase_fails_when_not_enough_inventory()
{
 var storeMock = new Mock<IStore>();
 storeMock.Setup(x => x.HasEnoughInventory(Product.Shampoo, 5)).Returns(false);
 var sut = new Customer();

 bool success = sut.Purchase(storeMock.Object, Product.Shampoo, 5);

 Assert.False(success);
 storeMock.Verify(x => x.RemoveInventory(Product.Shampoo, 5), Times.Never);
}
```

- 목과 스텁이 각기 다른 메서드를 다룬다. 따라서 스텁과의 상호작용을 검증하지 말라는 규칙을 위배하지 않았다.

#### 목과 스텁은 명령과 조회에 어떤 관련이 있는가?

- 목과 스텁은 **`명령 조회 분리(CQS, Command Query Separation)`** 원칙과 관련이 있다.
- CQS 원칙에서는 모든 메서드는 명령이거나 조회여야 한다.
  - 명령은 사이드 이펙트를 일으키고 어떤 값도 반환하지 않는다.
  - 조회는 사이드 이펙트를 일으키지 않고 값을 반환한다.
- **`CQS 원칙에서 명령을 대체하는 테스트 대역은 목이다. 반대로 조회를 대체하는 테스트 대역은 스텁이다`**.

```c#
var mock = new Mock<IEmailGateway>();
mock.Verify(x => x.SendGreetingsEmail("user@email.com"), Times.Once);

var stub = new Mock<IDatabase>();
stub.Setup(x => x.GetNumberOfUsers()).Returns(10);
```

- `SendGreetingsEmail`은 사이드 이펙트가 있는 명령으로 목이 대체한다.
- `GetNumberOfUsers`은 값을 반환하고 DB 상태를 변경하지 않으므로, 해당 테스트의 대역은 스텁이다.

### 식별할 수 있는 동작과 구현 세부 사항

- 단위 테스트에 리팩터링 내성 지표가 있는지 여부는 이진 선택이므로, 리팩터링 내성 지표가 가장 중요하다.
- 이를 위해 구현 세부 사항과 테스트를 떨어뜨려야 한다. 그렇다면 구현 세부 사항은 무엇이고 식별할 수 있는 동작은 뭘까?
- 코드가 식별할 수 있는 동작이라면 다음 중 하나를 해야 한다.
  - 클라이언트가 목표를 달성할 수 있는 연산(계산이나 사이드 이펙트) or 상태를 노출한다.
- 이상적으로 공개 API는 식별할 수 있는 동작과 일치해야 하며, 모든 구현 세부 사항은 클라이언트 눈에 보이지 않아야 한다.

```c#
public class User
{
 public string Name { get; set; }
 public string NormalizeName(string name)
 {
  string result = (name ?? "").Trim();
  if (result.Length > 50)
   return result.Substring(0, 50);

  return result;
 }
}

public class UserController
{
 public void RenameUser(int userId, string newName)
 {
  User user = GetUserFromDatabase(userId);
  string normalizedName = user.NormalizeName(newName);
  user.Name = normalizedName;

 SaveUserToDatabase(user);
 }
}
```

- 위 코드에선 속성과 메서드 모두 공개되어 있다. 클라이언트 입장에선 Name 속성만 필요한 작업이다.
- API를 잘 설계하기 위해 user 클래스는 NormalizeName 메서드를 숨기고 속성 세터를 내부적으로 호출해야 한다.

```c#
public class User
{
 private string _name;
 public string Name
 {
  get => _name;
 set => _name = NormalizeName(value);
 }

 private string NormalizeName(string name)
 {
  string result = (name ?? "").Trim();
  if (result.Length > 50)
   return result.Substring(0, 50);

  return result;
 }
}

public class UserController
{
 public void RenameUser(int userId, string newName)
 {
  User user = GetUserFromDatabase(userId);
  user.Name = newName;
  SaveUserToDatabase(user);
 }
}
```

- 위 예제는 식별할 수 있는 동작만 공개돼 있고, 구현 세부 사항은 비공개 API 뒤에 숨겨져있다.

#### 잘 설계된 API와 캡슐화

- 장기적으로 코드베이스 유지 보수에는 캡슐화가 중요하다.
- 계속해서 증가하는 코드 복잡도에 대처할 수 있는 방법은 캡슐화 말고는 실절적으로 없기 때문이다.
- 캡슐화는 궁극적으로 단위 테스트와 동일한 목표. SW의 지속적인 성장을 가능하게 하는 것이다.
- 잘 설계된 API 정의에서 연산과 상태를 최소한으로 노출해야 한다.
  - 클라이언트가 목표를 달성하는데 직접적으로 도움이 되는 코드만 공개해야 하며, 다른 세부 사항은 비공개 API 뒤로 숨겨야 한다.

### 목과 테스트 취약성 간의 관계

#### 육각형 아키텍처 정의

- 전형적인 어플리케이션은 도메인과 어플리케이션 서비스라는 두 계층으로 구성된다.
- 어플리케이션 서비스 계층과 도메인 계층의 조합은 육각형을 형성하며, 이 육각형은 어플리케이션을 나타낸다.
- 어플리케이션은 다른 어플리케이션과 소통할 수 있고 다른 어플리케이션도 육각형으로 나타낸다.
  - 예를 들어 SMTP, 서드파티 시스템, 메시지 버스 등이 될 수 있다.
- 이렇게 육각형이 서로 소통하면서 육각형 아키텍처를 구성한다.

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/books/code/hexagonal-architecture.png" width="500">

- 육각형 아키텍처라는 용어는 앨리스터 코오번이 세가지 중요한 지침을 강조하기 위해 처음 소개했다.
  - 도메인 계층과 어플리케이션 서비스 계층 간의 관심사 분리
    - 비지니스 로직은 어플리케이션의 가장 중요한 부분으로 도메인 계층만 해당 책임을 지고 이 외에 모든 책임에서는 제외 되어야 한다
    - 외부와 통신하거나 DB에 대한 책임은 어플리케이션 서비스가 담당해야 한다.
  - 어플리케이션 내부 통신
    - 육각형 아키텍처에서 어플리케이션 서비스 계층에서 도메인 계층으로 흐르는 단방향 의존성 흐름을 규정한다.
    - 도메인 계층은 도메인 계층 내부 클래스끼리 의존하고 어플리케이션 서비스 계층에 의존하지 않는다.
    - 도메인 계층은 외부 환경에서 완전히 격리돼야 한다.
  - 어플리케이션 간의 통신
    - 외부 어플리케이션은 어플리케이션 서비스 계층을 통해 연결된다.

#### 시스템 내부 통신과 시스템 간 통신

- 시스템 내부 통신은 어플리케이션 내부의 클래스 간의 통신이고 시스템 간 통신은 다른 어플리케이션과 통신하는 것을 말한다.
- 내부에서 도메인 클래스간 협력은 식별할 수 있는 동작이 아니라서 구현 세부 사항에 해당한다.
- 시스템 외부와 통신하는 방식은 전체적으로 시스템의 식별할 수 있는 동작을 나타내기에 목을 사용해서 확인하면 좋다.

```c#
public class CustomerController
{
 public bool Purchase(int customerId, int productId, int quantity)
 {
  Customer customer = _customerRepository.GetById(customerId);
  Product product = _productRepository.GetById(productId);

  bool isSuccess = customer.Purchase(_mainStore, product, quantity);
  if (isSuccess)
  {
    _emailGateway.SendReceipt(customer.email, product.Name, quantity);
  }

  return isSuccess;
 }
}
```

- 위 예제에서 이메일을 보내는 동작은 외부 환경에서 볼 수 있는 사이드 이펙트이므로 식별할 수 있는 동작을 나타낸다.
- 이메일을 보내는 호출을 목으로 하는 이유는 타당하다. 리팩터링 후에도 이러한 통신 유형이 유지되기에 테스트 취약성을 야기하지 않는다.
- 아래는 목을 사용하는 타당한 테스트를 보여준다.

```c#
[Fact]
public void Successful_purchase()
{
  var mock = new Mock<IEmailGateway>();
  var sut = new CustomerController(mock.Object);

  bool isSuccess = sut.Purchase(customerId: 1, productId: 2, quantity: 5);

  Assert.True(isSuccess);
  mock.Verify(x => x.SendReceipt("customer@email.com", "Shampoo", 5), Times.Once);
}
```

- 반대로 mock을 사용해 내부 클래스(Customer, Product)간의 상호작용을 검증하면 취약성을 야기한다.
