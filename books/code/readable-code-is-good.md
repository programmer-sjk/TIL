# 읽기 좋은 코드가 좋은 코드다

- [책 링크](https://m.yes24.com/Goods/Detail/6692314)

## 코드는 이해하기 쉬워야 한다

### 가독성의 기본정리

- 코드는 다른 사람이 그것을 **`이해하는데 들이는 시간을 최소화하는 방식으로`** 작성되어야 한다.
- 여기서 이해란 동료가 내 코드에서 쉽게 버그를 찾고, 수정하고, 수정한 내용이 기존 코드와 어떻게 상호작용 하는지 아는 것이다.

### 분량이 적으면 항상 더 좋은가?

- 일반적으로 **`분량이 적은 코드로 똑같은 문제를 해결할 수 있다면 그것이 더 낫다`**.
- 하지만 분량이 적다고 항상 좋은 것은 아니다. 이 코드를 이해하는데 걸리는 시간은

  ```java
  assert((!(bucket = FindBucket(key))) || !bucket->IsOccupied());
  ```

- 아래 두 줄짜리 코드를 이해 할 때보다 훨씬 많은 시간이 걸린다.

  ```java
  bucket = FindBucket(key);
  if (bucket != NULL) assert(!bucket->IsOccupied())
  ```

- 적은 분량으로 코드를 작성하는 것 보다, **`이해를 위한 시간을 최소화하는게 더 좋은 목표다`**.

## 이름에 정보 담기

### 특정한 단어 고르기

- 이름에 정보를 담는 방법 중 하나는 구체적인 단어를 선택하고 무의미한 단어를 피하는 것이다.
- 아래 get은 캐시, DB, 인터넷 중 어디서 페이지를 가져오는지 애매하다.

  ```js
    // 모호
    function getPage(url) {...}
  ```

- 인터넷에서 가져온다면 아래가 더 의미있는 이름이 된다.

  ```js
    function fetchPage() {...}
    function downloadPage() {...}
  ```

- 다음은 BinaryTree 클래스의 예다.

  ```java
    class BinaryTree {
      int size();
    }
  ```

- 위 size 메서드는 무엇을 반환할까? 트리의 높이, 노드의 개수, 트리의 메모리 용량?
- 문제는 size 라는 이름이 의도한 정보를 전달하지 못한다는데 있다. height(), numNodes()가 더 의미있을 것이다.

### tmp, retval 같은 보편적인 이름 피하기

- tmp, retval 같은 무의미한 이름이 아니라, 목적을 설명하는 이름을 골라야 한다.
- 아래 값을 교환하는 알고리즘 내에서는 임시 저장소로 사용된 tmp 이름이 명확하다.

  ```js
  if (right < left) {
    tmp = right;
    right = left;
    left = tmp;
  }
  ```

- 하지만 아래 코드에서 tmp 이름은 게으름의 산물에 불과하다.

  ```js
  const tmp = user.name();
  tmp += ' ' + user.phoneNumber();
  tmp += ' ' + user.email();
  template.set('user_info', tmp);
  ```

- 변수의 scope가 짧긴 하지만 임시적인 저장소의 역할로 국한되지 않았기에 userInfo 같은 이름이 더 적절하다.
