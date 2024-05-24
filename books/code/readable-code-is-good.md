# 읽기 좋은 코드가 좋은 코드다

- [책 링크](https://m.yes24.com/Goods/Detail/6692314)

## 코드는 이해하기 쉬워야 한다

### 가독성의 기본정리

- 코드는 다른 사람이 그것을 **`이해하는데 들이는 시간을 최소화하는 방식으로`** 작성되어야 한다.
- 여기서 이해란 동료가 내 코드에서 쉽게 버그를 찾고, 수정하고, 수정한 내용이 기존 코드와 어떻게 상호작용 하는지 아는 것이다.

### 분량이 적으면 항상 더 좋은가?

- 일반적으로 **`분량이 적은 코드로 똑같은 문제를 해결할 수 있다면 그것이 더 낫다`**.
- 하지만 분량이 적다고 항상 좋은 것은 아니다. 이 코드를 이해하는데 걸리는 시간은

  ```c++
  assert((!(bucket = FindBucket(key))) || !bucket->IsOccupied());
  ```

- 아래 두 줄짜리 코드를 이해 할 때보다 훨씬 많은 시간이 걸린다.

  ```c++
  bucket = FindBucket(key);
  if (bucket != NULL) assert(!bucket->IsOccupied())
  ```

- 적은 분량으로 코드를 작성하는 것 보다, **`이해를 위한 시간을 최소화하는게 더 좋은 목표다`**.

## 이름에 정보 담기

### 특정한 단어 고르기

- 이름에 정보를 담는 방법 중 하나는 **`구체적인 단어를 선택하고 무의미한 단어를 피하는 것이다`**.
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

- 다음은 `BinaryTree` 클래스의 예다.

  ```java
    class BinaryTree {
      int size();
    }
  ```

- 위 size 메서드는 무엇을 반환할까? 트리의 높이, 노드의 개수, 트리의 메모리 용량?
- 문제는 **`size 라는 이름이 의도한 정보를 전달하지 못한다는데 있다`**. `height(), numNodes()`가 더 의미있을 것이다.

### tmp, retval 같은 보편적인 이름 피하기

- `tmp, retval` 같은 무의미한 이름이 아니라, 목적을 설명하는 이름을 골라야 한다.
- 아래 값을 교환하는 알고리즘 내에서는 임시 저장소로 사용된 `tmp` 이름이 명확하다.

  ```js
  if (right < left) {
    tmp = right;
    right = left;
    left = tmp;
  }
  ```

- 하지만 아래 코드에서 `tmp` 이름은 게으름의 산물에 불과하다.

  ```js
  const tmp = user.name();
  tmp += ' ' + user.phoneNumber();
  tmp += ' ' + user.email();
  template.set('user_info', tmp);
  ```

- 변수의 scope가 짧긴 하지만 임시적인 저장소의 역할로 국한되지 않았기에 `userInfo` 같은 이름이 더 적절하다.

### 추상적인 이름보다 구체적인 이름을 선호하라

- 서버가 어떤 포트를 사용하는지 검사하는 `serverCanStart` 메소드가 있다고 가정하자.
  - 이 이름은 다소 추상적이기에 `canListenOnPort`로 더 구체적인 이름을 사용할 수 있다.
- 프로그램에서 `--run_locally` 옵션을 사용한 적이 있다.
  - 이 플래그가 켜지면 디버깅 정보를 출력하고 대신 동작 속도는 다소 느려진다.
  - **`이 이름에는 몇가지 문제가 있다`**.
    - 팀에 새로 합류한 사람은 무엇을 위한 플래그인지 모르고, 왜 필요한지 알 수 없다.
    - 원격에서 실행할 때도 디버깅 정보를 출력할 수 있다.
  - 문제는 실제 내용보다, 주로 사용되는 환경을 나타내는 방식으로 지어졌다는 점이다.
  - 기존 이름보다는 `--extra_logging` 같은 이름이 더 직접적이고 명확하다.

### 추가적인 정보를 이름에 추가하기

- **`변수의 이름은 작은 설명문이다`**. 사용자가 알아야 하는 정보를 추가적인 단어로 전달하는게 좋다.

  ```java
    // not good
    string id = 'af84ef845cd8' // 16진수

    // good
    string hex_id = 'af84ef845cd8'
  ```

- 변수가 시간이나 바이트 같은 측정치를 담고 있다면 변수명에 포함시키는게 도움이 된다.

  ```js
  // not good
  const start = new Date().getTime();

  // good
  const startMs = new Date().getTime();
  ```

### 이름은 얼마나 길어야 하는가?

- 이름이 지나치게 길면 좋지 않다는 인식이 암묵적으로 존재한다. 이름 길이에 대한 조언을 살펴보자.
- 좁은 범위에서는 짧은 이름이 괜찮다.
- 이름은 짧을수록 좋지만, 짧은 이름으로 설명이 부족하다면 긴 이름을 사용할 수 있다.
- 약어나 축약의 경우, 새로운 사람이 합류해도 이해할 수 있다면 괜찮다.
  - ex) `evaluation -> eval, document -> docs, string -> str`
- 경우에 따라 정보를 손실하지 않고 이름에 포함된 단어를 제거할 수 있다.
  - `convertToString()` 이름 대신 `toString()` 이라고 짧게 써도 실질적인 정보를 사라지지 않는다.

### 요약

- 이름에 모호하고 보편적인 단어를 피하고, 구체적인 단어를 사용하자
- 변수명에 ms 같은 세부적인 정보를 붙이고, 사용범위가 넓다면 짧은 이름을 사용하지 말자.

## 오해할 수 없는 이름들

- 본인이 선택한 이름을 **`다른 사람들이 다른 의미로 해석할 수 있을까?`** 라는 질문을 던져야 한다.
- 경계를 포함하는 한계갑을 다룰 때는 `min/max`를 사용하자.
- 경계를 포함하는 범위에는 `first/last`를 사용하자.
- 경계를 포함하고 배제하는 범위에는 `begin/end`를 사용하자.

  ```js
    if (shoppingCart.numItems > MAX_ITEMS_IN_CART)
    function numberRange(first, last) // first <= x <= last

    function printEventInRange(begin, end) // begin <= x < end
    printEventInRange('2024-01-01', '2024-01-02') // 1일을 포함해서 2일까지
  ```

### 불리언 변수에 이름 붙이기

- 일반적으로 **`is, has, can, should`** 단어를 활용하면 불리언 값의 의미가 더 명확해진다.
- 이름에서는 의미를 부정하는 용어는 피하는게 좋다.
  - `disableSsl` 보다는 `useSsl` 이름이 더 읽기 좋다.

### 요약

- **`언제나 의미가 오해되지 않는 이름이 좋다`**.
- 값의 상한과 하한을 정할때는 `min/max`가 좋은 접두사 역할을 한다.
- 경계를 포함한다면 `first/last`를, 경계의 마지막을 배제한다면 `begin/end`가 널리 사용되는 이름이다.
- 불리언 이름을 정할 때는 불리언이라는 사실을 드러내기 위해 `is, has` 같은 단어를 사용하는게 좋다.
- 사람들이 일반적으로 기대하는 `get, size` 같은 함수는 복잡하지 않은 가벼운 함수라고 기대할지도 모른다.

## 미학

- 좋은 코드는 눈을 편하게 해야 한다.

### 미학이 무슨 상관인가?

- 미학적으로 **`보기 좋은 코드가 읽기 쉽고 사용하기 쉽다는 사실은 명백하다`**.
- 가독성이 떨어지는 긴 코드와, 논리적으로 쪼개지고 가독성이 좋은 긴 코드를 상상해보자.

### 도움이 된다면 코드의 열을 맞춰라

- 경우에 따라 열 정렬을 통해 코드를 더 읽기 쉽게 할 수 있다.

  ```js
  // POST 파라미터를 지역변수에 저장. (lint 때문에 열 정렬 안되는 중)
  details = request.POST.get('details');
  location = request.POST.get('location');
  phone = request.POST.get('phone');
  email = request.POST.get('email');
  url = request.POST.get('url');
  ```

### 코드를 문단으로 쪼개라

- 비슷한 생각을 묶어서, 성격이 다른 생각과 구분한다.

  ```python
    # bad
    def suggest_new_feinds(user, email_password):
      friends = user.friends()
      friend_emails = set(...)
      contacts = ...
      non_friend_emails = ...
      suggested_friends = ...
      display['user'] = ...
      display['friends'] = ...
      display['suggested_friends'] = ...
      return render(...)

    # good
    def suggest_new_feinds(user, email_password):
      # 사용자 친구들의 정보를 얻음
      friends = user.friends()
      friend_emails = set(...)

      # 사용자 이메일 계정으로부터 작업
      contacts = ...

      # 친구가 아닌 사용자 찾기
      non_friend_emails = ...
      suggested_friends = ...

      # 사용자 정보를 화면에 출력
      display['user'] = ...
      display['friends'] = ...
      display['suggested_friends'] = ...

      return render(...)
  ```

### 개인적인 스타일 대 일관성

- 아래 두가지 스타일 중 하나를 선택한다고 가독성에 실질적인 영향을 주진 않는다.

  ```java
    class Logger {

    }

    class Logger
    {

    }
  ```

- 하지만 **`두 스타일이 섞이면 가독성에 영향을 준다`**.
- 저자 기준에 잘못된 스타일을 사용하는 경우도 많았지만, 일관성 유지가 훨씬 중요하므로 해당 프로젝트의 스타일을 따랐다.

### 요약

- **`코드를 일관성 있게, 의미있게 정렬하면 읽기 더 편하고 빠르게 만들 수 있다`**.
- 여러 코드가 비슷한 일을 수행하면, 실루엣이 동일해 보이게 만들자.
- 열로 만들어서 줄을 맞추면 코드를 한 눈에 보기 편하다.
- 빈 줄을 이용해 커다란 문단을 논리적인 문단으로 나누자.

## 주석에 담아야 하는 대상

- 주석의 목적은 코드를 읽는 사람이, 코드를 작성한 사람만큼 이해하도록 돕는데 있다.

### 설명하지 말아야 하는 것

- **`설명 자체를 위한 주석을 달지 말자`**. 아래는 무가치한 주석에 속한다.

  ```c++
    // 주어진 이름과 깊이를 이용해서 서브트르에 있는 노드를 찾는다.
    Node* FindNodeInSubtree(Node* subtree, String name, int depth)
  ```

- 이 함수를 위한 주석을 달고 싶다면 중요한 세부사항을 적는 것이 낫다.

  ```c++
    // 주어진 name으로 노드를 찾거나 아니면 NULL을 반환한다.
    // 만약 depth <= 0이면 subtree만 반환된다.
    // 만약 depth == N 이면 N 레벨과 그 아래만 검색된다.
    Node* FindNodeInSubtree(Node* subtree, String name, int depth)
  ```

- **`나쁜 이름에 주석을 달지 말고 이름을 고쳐라`**

  ```c++
    // 키의 핸들을 해제한다. 실제 레지스트리를 삭제하진 않는다.
    void deleteRegistry(RegistryKey* key);
  ```

- 주석 대신 이름이 정확한 설명을 하는편이 낫다.

  ```c++
    void ReleaseRegistryHandle(RegistryKey* key);
  ```

- 일반적으로 사람들은 **`나쁜 가독성을 메우려고 노력하는 애쓰는 주석을 원하지 않는다`**.

### 생각을 기록하라

- 좋은 주석은 단순히 생각을 기록하는 것만으로 탄생할 수 있다.

  ```js
  // 여기서 이진트리는 해시 테이블보다 40% 정도 빠르다.
  // 이 클래스는 점점 엉망이 되고 있어서, ResourceNode 하위 클래스를 리팩토링 해야 한다.
  ```

- 코드에 있는 개선사항이나 결함을 설명하라.

  ```js
  // TODO: 더 빠른 알고리즘을 찾아 사용하라.
  // TODO(더스틴): JPEG 말고 다른 이미지 포맷도 처리할 수 있어야 한다.
  ```

- 상수에 대한 설명

  ```js
  const IMAGE_QUALITY = 0.72; // 사용자들은 0.72가 크기/해상도 대비 최선이라고 생각한다.
  ```

### 코드를 읽는 사람의 입장이 되어라

- 이 책은 코드를 처음으로 읽는 외부인의 입장에 자기 자신을 놓는 기법을 다루고 있다.
- 나올 것 같은 질문을 예측해서 주석을 달 수도 있다.

  ```c++
    // 벡터가 메모리를 반납하도록 강제한다. ("STL swap trick을 보라")
    vector<float>().swap(data);
  ```

- 사람들이 쉽게 빠질 함정을 경고하라.

  ```js
    // 외부 서비스를 호출해 이메일 서비스를 호출한다. (1분 이후 타임아웃 된다.)
    void sendEmail(string to, string subject, string body);
  ```

### 요약

- 주석을 다는 목적은, **`코드를 작성하는 사람이 코드를 읽는 사람에게 정보를 전달하는 것이다`**.
- 코드에서 빠르게 알 수 있는 사실이나, 나쁜 이름은 주석을 달지 말자
- 코드가 작성된 이유를 설명하거나, 코드 결함(TODO), 상수가 값을 갖게 된 사연엔 주석을 달자
- 코드를 처음 본느 사람들을 위한 주석을 남겨두자
  - 코드를 읽는 사람이 왜? 라고 생각할 부분을 예측해 주석을 달자.
  - 일반적으로 예상하지 못할 특이한 동작을 기록하자.

## 읽기 쉽게 흐름제어 만들기

### 조건문에서 인수의 순서

- 다음 두 코드 중 어떤 코드가 읽기 쉬운가?

  ```js
    if (lenght >= 10)
    if (10 >= length)
  ```

- 대부분은 첫 번째 코드가 더 읽기 쉽다고 느낀다. 그렇다면 아래는 어떤가?

  ```js
    if (bytesReceived < bytesExpected)
    if (bytesExpected > bytesReceived)
  ```

- 이 경우에도 첫 번째가 읽기 쉽다. 이 가이드라이은 영어 어순과 일치한다.
  - 당신이 18세라면, 당신이 1년에 10만불을 번다면
- 때문에 왼쪽에는 질문을 받는 표현, 오른쪽은 비교대상으로 표현되는 값이 오는게 좋다.

### if/else 순서

- 부정이 아닌 긍정을 다루자.

  ```js
  if (a == b) {
  } else {
  } // ok
  if (a != b) {
  } else {
  } // ok
  ```

- 더 흥미롭고 확실한 것을 먼저 다루자. (early return의 경우는 반대)
  - if에 더 중요한 것을 두고 else에 아닌것을 배치하자.

### do/while 루프를 피하자

- 일반적으로 if, while, for 동작이 위에서 아래로 읽는다.
- do/while 문은 역순으로 코드를 두 번 읽어야 하기 때문에 부자연스럽다.

### 중첩을 최소화하기

- 코드의 중첩이 심할수록 이해하기 어렵다.

  ```js
  if (userResult == SUCCESS) {
    if (permission != SUCCESS) {
      reply.writeErrors('...');
      reply.done();
      return;
    }
    reply.writeErrors('');
  } else {
    reply.writeErrors(userResult);
  }
  reply.done();
  ```

- 위 코드는 중첩되지 않은 코드에 비해 읽기 어렵다.
  - userResult와 permission 결과를 머릿속에 저장한 상태에서 코드를 읽어나가야 한다.
  - 또한 결과가 SUCCESS 인 경우와 아닌 경우를 계속해서 왔다 갔다 해야 한다.
- early return 하여 중첩을 제거하라.

  ```js
  if (userResult != SUCCESS) {
    reply.writeErrors(userResult);
    reply.done();
    return;
  }

  if (permission != SUCCESS) {
    reply.writeErrors(permission);
    reply.done();
    return;
  }

  reply.writeErrors('');
  reply.done();
  ```

- 루프 내부에 중첩을 제거하자
  - if (...) return 형태처럼 if (...) continue 구문으로 중첩을 제거할 수 있다.

### 요약

- 흐름제어 코드에서 변하는 값을 왼쪽에 두고, 안정적인 값을 오른쪽에 두는 것이 좋다.
- if/else 에서 긍적적이고 중요한 경우를 앞에 두어라.
- 과도한 삼항연산자, do/while, goto 문은 종종 코드의 가독성을 떨어뜨린다.
- 중첩된 구조보다는 선형적인 코드를 추구하자. 함수 중간에 반환하면 코드를 더 깔끔하게 작성할 수 있다.
