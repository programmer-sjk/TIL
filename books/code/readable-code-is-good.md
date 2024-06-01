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

- **`이름에 모호하고 보편적인 단어를 피하고, 구체적인 단어를 사용하자`**
- 변수명에 ms 같은 세부적인 정보를 붙이고, 사용범위가 넓다면 짧은 이름을 사용하지 말자.

## 오해할 수 없는 이름들

- 본인이 선택한 이름을 **`다른 사람들이 다른 의미로 해석할 수 있을까?`** 라는 질문을 던져야 한다.
- 경계를 포함하는 한계갑을 다룰 때는 `min/max`를 사용하자.
- 경계를 포함하는 범위에는 `first/last`를 사용하자.
- 경계를 포함하고 배제하는 범위에는 `begin/end`를 사용하자.

  ```js
  if (shoppingCart.numItems > MAX_ITEMS_IN_CART);
  function numberRange(first, last) {} // first <= x <= last

  function printEventInRange(begin, end) {} // begin <= x < end
  printEventInRange('2024-01-01', '2024-01-02'); // 1일을 포함해서 2일까지
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
    // 주어진 이름과 깊이를 이용해서 서브트리에 있는 노드를 찾는다.
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

- 이 책은 코드를 처음으로 읽는 동료의 입장에 자기 자신을 놓는 기법을 다루고 있다.
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
    if (length >= 10)
    if (10 >= length)
  ```

- 대부분은 첫 번째 코드가 더 읽기 쉽다고 느낀다. 그렇다면 아래는 어떤가?

  ```js
    if (bytesReceived < bytesExpected)
    if (bytesExpected > bytesReceived)
  ```

- 이 경우에도 첫 번째가 읽기 쉽다. 이 **`가이드 라인은 영어 어순과 일치한다`**.
  - `당신이 18세라면, 당신이 1년에 10만불을 번다면`
- 때문에 **`왼쪽에는 질문을 받는 표현, 오른쪽은 비교대상으로 표현되는 값이 오는게 좋다`**.

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

- **`더 흥미롭고 확실한 것을 먼저 다루자`**. (early return의 경우는 반대)
  - if에 더 중요한 것을 두고 else에 아닌것을 배치하자.

### do/while 루프를 피하자

- 일반적으로 `if, while, for` 동작이 위에서 아래로 읽는다.
- `do/while` 문은 역순으로 코드를 두 번 읽어야 하기 때문에 부자연스럽다.

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
  - `userResult와 permission` 결과를 머릿속에 저장한 상태에서 코드를 읽어나가야 한다.
  - 또한 결과가 SUCCESS 인 경우와 아닌 경우를 계속해서 왔다 갔다 해야 한다.
- **`early return 하여 중첩을 제거하라`**.

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
  - `if (...) return` 형태처럼 `if (...) continue` 구문으로 중첩을 제거할 수 있다.

### 요약

- 흐름제어 코드에서 변하는 값을 왼쪽에 두고, 안정적인 값을 오른쪽에 두는 것이 좋다.
- `if/else` 에서 긍적적이고 중요한 경우를 앞에 두어라.
- 과도한 삼항 연산자, `do/while`, `goto` 문은 종종 코드의 가독성을 떨어뜨린다.
- **`중첩된 구조보다는 선형적인 코드를 추구하자`**. 함수 중간에 반환하면 코드를 더 깔끔하게 작성할 수 있다.

## 거대한 표현을 잘게 쪼개기

### 설명 변수 / 요약 변수

- 하위 표현을 설명하는 설명 변수 예시를 보자.

  ```js
    if (line.split(':')[0].stipe() == 'root') {...}

    // 설명 변수 사용
    const username = line.split(':')[0].stipe();
    if (username === 'root') {...}
  ```

- 코드의 덩어리를 변수로 쉽게 관리 및 파악하는 변수를 요약 변수라고 한다.

  ```js
    if (req.user.id === document.ownerId) {...}

    // 개선
    const userOwnsDocument = req.user.id === document.ownerId;
    if (userOwnsDocument) {...}
  ```

### 드모르간 법칙 사용하기

- 수학이나 과학시간에 드모르간 법칙을 기억할 것이다.
- 이 법칙으로 불리언 표현을 간단하게 만들 수 있다.

  ```js
    if (!(fileExists && !isProtected)) console.error(...)

    // 개선
    if (!fileExists || isProtected) console.error(...)
  ```

### 쇼트 서킷 논리 오용하지 않기

- 대부분의 프로그래밍 언어는 쇼트 서킷 평가를 지원한다.
- 예로 `if (a || b)` 에서 a가 참이면 b를 평가하지 않는다.
- 매우 편리하지만 **`복잡한 연산을 수행할 때는 오용될 수 있다`**.

  ```c++
    assert((!(bucket = FindBucket(key))) || !bucket->IsOccupied());
  ```

- 위 코드는 한 줄이지만, 대부분의 **`프로그래머는 이해하기 위해 손을 멈추고 생각을 해야 한다`**.
- 아래 코드는 동일한 일을 수행한다. 코드는 두 줄로 늘어났지만 훨씬 이해하기 쉬워졌다.

  ```c++
    bucket = FindBucket(key);
    if (bucket != NULL) assert(!bucket->IsOccupied())
  ```

### 요약

- 거대한 표현을 쪼개서 코드를 읽는 사람이 더 쉽게 소화하는 방법을 알아보았다.
- 하위표현을 대체하는 설명 변수는 코드를 읽는 사람이 코드의 핵심 개념을 파악하는 것을 돕는다.
- 드모르간 법칙을 이용해 복잡한 if 문을 개선할 수 있다.

## 변수와 가독성

- 변수는 아래 3가지 경우 기억하고 다루기 어려워지는 문제가 있다.
  - **`변수의 수가 많을수록 / 변수의 범위가 늘어날수록 / 변수 값이 자주 바뀔수록`** 어려워진다.

### 변수 제거하기

- 아래 코드에서 `now` 변수가 필요할까?

  ```python
    now = datetime.datetime.now()
    root_message.last_view_time = now
  ```

- 그렇지 않다. `datetime.now()`는 그 자체로 명확하고, now가 재 사용되지 않기 떄문이다.

  ```python
    root_message.last_view_time = datetime.datetime.now()
  ```

### 변수의 범위를 좁혀라

- 전역 변수는 어디에서 어떻게 사용되는지 파악하기 어려우니 피하라는 조언을 한 번쯤 들었을 것이다.
- 사실 전역 변수뿐 아니라 **`모든 변수의 범위를 좁히는 일은 언제나 좋다`**.

### 값을 한 번만 할당하는 변수를 선호하라

- 변수들의 **`값이 변한다면 프로그램을 따라가는 일이 더욱 어려워지기 때문에`** 불변 변수를 사용하자.
- 다음과 같은 상수는 코드를 읽는 사람에게 추가적인 생각을 요구하지 않는다.

  ```java
    static final int NUM_THREADS = 10;
  ```

- 만약 불변 변수가 아니더라도, 재할당을 최대한 줄이는 것은 여전히 도움이 된다.

### 요약

- 변수를 덜 사용하고, 최대한 가볍게 만들어 가독성을 높일 수 있다.
- **`변수의 범위를 최대한 줄이고 불변 변수 (const, final 등등)를 사용하자`**

## 상관없는 하위문제 추출하기

- 어떤 코드가 있을 때, 상위 수준에서 이 코드의 목적이 무엇인지 물어야 한다.
  - **`이 코드는 직접적인 목적을 위해서 존재하는가?`**
  - **`목적을 위해 필요하긴 하지만 목적과 직접적으로 상관없는 하위 문제를 해결하는가?`**
- 예를 들어, 파일을 업로드 하는 기능이 있다.
  - 헌데 **`데이터를 파싱해서 로깅하는 하위 로직들이`** 코드 대부분을 차지하며 가독성을 어지럽힌다.
  - 이런 경우 하위 문제를 함수나 메서드로 추출하자.

### 특정한 목적을 위한 기능

- 아래 파이썬 코드는 `Business` 객체를 만들고, `name, url` 값을 설정한다.

  ```python
    business = Business()
    business.name = req.POST['name']
    url_path_name = business.name.lower()
    url_path_name = re.sub(r"['\.]", "", url_path_name)
    url_path_name = re.sub(r"[^a-z0-9]+", "-", url_path_name)
    business.url = "/biz/" + url_path_name
  ```

- 이때 name을 기반으로 url을 만드는 하위 문제가 있으므로 이를 쉽게 추출할 수 있다.

  ```python
    CHARS_TO_REMOVE = re.compile(r"['\.]+")
    CHARS_TO_DASH = re.compile(r"[^a-z0-9]+")

    def make_url_friendly(text):
      text = text.lower()
      text = CHARS_TO_REMOVE.sub('', text)
      text = CHARS_TO_DASH.sub('-', text)
      return text

    business = Business()
    business.name = req.POST['name']
    business.url = "/biz/" + make_url_friendly(business.name)
  ```

- 결과적으로 코드를 읽으면서 **`복잡한 문자열 처리를 신경쓰지 않아도 되서 가독성이 더 좋아졌다`**.

### 지나치게 추출하기

- **`함수를 너무 많이 쪼개면 오히려 가독성을 해친다`**.
  - 사용자가 신경써야 하는 내용이 늘어나고, 실행 경로를 추적하려면 코드 곳곳을 돌아다녀야 하기 떄문이다.
- **`코드에 새로운 함수를 더하는 일은 약간의 가독성 비용이 든다`**.
- 작게 쪼갠 함수들이 다른 곳에서도 사용되는게 아니라면, 너무 지나치게 분리할 필요는 없다.

### 요약

- 일반적인 목적의 코드(파싱, 로깅, 변환)를 특정 코드에서 분리해야 한다.
- 문제를 해결하기 위한 라이브러리, 헬퍼 함수를 분리하면 작은 핵심들만 남을 것이다.

## 한번에 하나씩

- **`한 번에 여러가지 일을 수행하는 코드는 이해하기 어렵다`**.
- 블록 안에서 객체 초기화, 데이터 청소, 입력 분석, 비지니스 로직이 섞이는 경우 그렇다.

### 작업은 작을 수 있다

- 댓글에 추천/반대 할 수 있는 투표가 있다고 가정하자. 사용자가 투표 버튼을 누르면 아래 함수가 호출된다.

  ```js
  function voteChanged(oldVote, newVote) {
    const score = getScore();

    if (oldVote != newVote) {
      if (newVote === 'UP') {
        score += oldVote === 'DOWN' ? 2 : 1;
      } else if (newVote === 'DOWN') {
        score += oldVote === 'UP' ? 2 : 1;
      } else if (newVote === '') {
        score += oldVote === 'UP' ? -1 : 1;
      }
    }

    setScore(score);
  }
  ```

- 위 함수는 **`코드의 길이는 짧지만 두 가지 작업을 수행한다`**.
  - `oldVote, newVote`로 score 값을 구한다.
  - 점수가 반영된다.
- 각각의 작업을 분리하여 코드를 더 읽기 편하게 만들 수 있다.

  ```js
  function voteChanged(oldVote, newVote) {
    const score = getScore();
    score -= voteValue(oldVote);
    score += voteValue(newVote);
    setScore(score);
  }

  function voteValue(vote) {
    if (vote === 'UP') {
      return 1;
    }

    if (vote === 'DOWN') {
      return -1;
    }

    return 0;
  }
  ```

### 요약

- 작성한 코드가 읽기 어렵다면, **`일단 수행하는 작업을 모두 나열하라`**.
- 나열된 작업 중 일부는 별도의 함수나 클래스로 분리할 수 있을 것이다.

## 생각을 코드로 만들기

- 자신의 생각을 상대에게 쉬운말로 전하는 기술은 매우 소중하다.
- 코드도 동료에게 말하듯이 쉽고 핵심적인 문구를 사용하자

### 논리를 명확하게 설명하기

- 다음은 PHP 코드로 사용자가 페이지를 볼 수 있는지 확인한다.

  ```php
    $is_admin = is_admin_request();
    if ($document) {
      if (!is_admin && ($document['username'] != $_SESSION['username'])) {
        return not_authorized();
      }
    } else {
      if (!$is_admin) {
        return not_authorized();
      }
    }
  ```

- 위 코드에는 상당한 논리가 있지만 단순화하면 관리자와 소유자만 문서를 볼 수 있다.

  ```php
    if (is_admin_request()) {
      // 허가
    } else if ($document && $document['username'] != $_SESSION['username']) {
      // 허가
    } else {
      return not_authorized();
    }
  ```

- 위에 개선된 **`코드는 분량이 적고, 중첩이 없으며, 부정문이 없어 논리가 더 간단해졌다`**.

### 논리를 쉬운말로 표현하기

- 구체적인 코드 예시는 너무 길어, 이 장의 핵심을 말로 간략하게 정리하겠다.
- **`구체적인 코드를 감춰서 상위 수준의 코드를 읽기 쉽게 가독성을 높여야 한다`**.
- 일반적인 설계원리는 **`덜 중요한 세부사항은 사용자가 볼 필요 없게 숨겨서 더 중요한 내용이 눈에 잘 띄게 해야 한다`**.

### 요약

- 프로그램을 평범한 말로 설명하고, 그 설명으로 자연스러운 코드를 작성하는 테크닉을 살펴보았다.
- 자신의 **`문제를 쉬운 말로 설명할 수 없으면, 해당 문제를 제대로 이해하지 못한 것이다`**.

## 코드 분량 줄이기

- **`가장 읽기 쉬운 코드는 아무것도 없는 코드다`**.

### 기능을 구현하려고 애쓰지 마라

- 프로그래머는 **`정말로 필요한 기능인지 잘못 판단하는 경우도 많다`**.
- 또 어떤 기능을 구현하는데 필요한 노력을 과소평가 하는 경향도 있다.
- 구현 시간을 낙관적으로 예측하고, 코드를 추후 유지보수하고, 문서를 만드는데 필요한 시간을 잊어버린다.

### 요구사항에 질문을 던지고, 질문을 나누어 분석하라

- 프로그램이 반드시 빠르게 동작하고 100% 정확하며, 모든 입력을 처리해야 하는 것은 아니다.
- **`요구사항을 잘 분석하면 적은 코드로 구현할 수 있는 간단한 문제를 정의할 수 있다`**.
- 위치 추적을 100% 정확하게 구현하려면 매우 복잡하다.
  - 애플리케이션이 텍사스 주에 있는 가게 30곳만 대상으로 한다면 요구사항을 축소할 수 있다.
  - 1마일 마다 경도/위도를 고려할 필요가 없으므로 해결책은 훨씬 더 수월해졌다.
- **`요구사항을 제거하기와 더 간단한 문제를 해결하기가 제공하는 이점은`** 아무리 강조해도 지나치지 않는다.

### 자기 주변에 있는 라이브러리에 친숙해져라

- 라이브러리가 할 수 있는 일을 알고 활용하는 것은 중요하다.
- 다음은 실제로 도움을 주는 조언이다.
  - 매일 15분씩 표준 라이브러리에 있는 함수/모듈의 이름을 읽어라.
- 라이브러리가 대충 무엇을 제공하는지 감을 잡아놓고 나중에 필요할 때 생각할 수 있기를 바라는 것이다.

### 요약

- 새로 작성하는 코드는 모두 테스트, 문서화, 유지보수 해야 한다.
- 아래와 같은 방법으로 새로운 코드를 작성하는 일을 피할 수 있다.
  - 제품에 꼭 필요하지 않은 기능은 제거하고, 과도한 작업 피하기
  - 요구사항을 다시 생각해, 가장 단순한 형태의 문제를 찾기
  - 주기적으로 라이브러리 API 흝어보기

## 테스트와 가독성

### 읽거나 유지보수하기 쉽게 테스트 코드를 만들어라

- **`테스트 코드가 읽기 쉬우면`**, 사용자는 실제 코드가 어떻게 동작하는지 쉽게 이해할 수 있다.
- 다른 프로그래머가 수정하거나, 새로운 테스트를 더하는걸 쉽게 느낄 수 있도록, **`테스트 코드는 읽기 쉬워야 한다`**.

### 이 테스트는 어떤 점이 잘못되었을까?

- 점수가 있는 페이지를 필터링하는 아래 테스트 코드를 보자.

  ```c++
    void Test1() {
      vector<ScoredDocument> docs;
      docs.resize(3);
      docs[0].url = 'https://example.com'
      docs.score = -5.0;
      docs[1].url = 'https://example.com'
      docs.score = 1;
      docs[2].url = 'https://example.com'
      docs.score = 4;

      sortAndFilterDocs($docs);

      assert(docs.size() == 3)
      assert(docs[0].score == 4)
      assert(docs[1].size() == 3)
      assert(docs[2].size() == 1)
    }
  ```

- **`테스트가 너무 길고 중요하지 않은 자세한 내용으로 가득 찼다`**.
- assert가 제공하는 테스트 실패 메시지가 별로 도움이 되지 않는다.
- 이름이 의미가 없으며 비정상 값을 가지는 입력을 테스트하지 않는다.
- 너무 많은 기능을 테스트 하고 있다.

### 테스트에 친숙한 개발

- 어떤 코드는 테스트하기 쉽다.
- 테스트하기 좋은 코드는 잘 정의된 인터페이스를 가지고, 지나치게 많은 상태를 가지지 않는다.
- 반대로 전역 변수나, 외부 컴포넌트를 많이 호출하고, 논리가 어렵다면 테스트 코드를 작성하기 어렵다.

### 요약

- 테스트 코드에서 가독성은 여전히 중요한 자리를 차지한다.
- **`테스트가 읽기 편하면, 작성하기 쉬워지고, 다른 사람들도 더 많은 테스트를 작성하게 된다`**.
- 실제 코드를 **`테스트하기 쉬운 방식으로 작성하면, 실제 코드도 전반적으로 좋은 설계를`** 가지게 된다.
