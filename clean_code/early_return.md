# early return 정리

## 정리하게 된 배경

- 최근에 기능 개발을 하다가 아래와 같은 흐름의 함수를 작성했다.

```js
public doSomething() {
  if (this.isUpdatable()) {
    // 업데이트 하는 로직
    return;
  }

  // 초기화하는 로직
}
```

- 그리고 코드 리뷰에서 밑에 초기화하는 로직이 early return 안에 있어야 한다는 말을 들었다. 생각해보니 예전에 early return은 덜 중요한 부분을 빠르게 뱉어내고 상대적으로 중요한 내용을 아래 넣어야 한다는 이야기를 들었던 것 같아 수정하겠다고 comment를 남겼다.
- 이에 대해 시니어 개발자분들 의견이 아래와 같았는데 조금 더 정리할 필요성을 느껴서 내용을 정리해본다.

```text
시니어 A
Early return은 중요도가 아니고 먼저 확인할 사항순이고 조건문 안 코드 최소화가 중점입니다. 긴 코드가 Depth가 깊어져 가독성을 해치지않기위함

시니어 B
이정도 코드에서는 어떻게해도 큰 상관은 없긴한데 로직적으로 중요한 부분(주로 관심을 가져야 하는 부분)이 뎁스가 깊게 안들어가도록 만드는 것도 일부 맞는 말입니다
```

## early return

- early return은 정상적인 결과를 함수의 마지막에 위치시키고, 조건에 맞지 않는다면 남은 코드들은 빠르게 종료되도록 코드를 작성하는 방법이다.
- 극단적으로 아래와 같은 중첩 조건문 보다는

```js
public String returnStuff(SomeObject argument1, SomeObject argument2) {
  if (argument1.isValid()) {
    if (argument2.isValid()) {
      SomeObject otherVal1 = doSomeStuff(argument1, argument2)

      if (otherVal1.isValid()) {
				SomeObject otherVal2 = doAnotherStuff(otherVal1)

				if (otherVal2.isValid()) {
					return "Stuff";
				} else {
					throw new Exception();
				}
			} else {
				throw new Exception();
			}
		} else {
			throw new Exception();
		}
	} else {
		throw new Exception();
	}
}
```

- 아래와 같은 형태가 읽기 쉽다는 것은 다들 알고 있을 것이다.

```js
public String returnStuff(SomeObject argument1, SomeObject argument2){
  if (!argument1.isValid()) {
    throw new Exception();
  }

  if (!argument2.isValid()) {
    throw new Exception();
  }

  SomeObject otherVal1 = doSomeStuff(argument1, argument2);

  if (!otherVal1.isValid()) {
    throw new Exception();
  }

  SomeObject otherVal2 = doAnotherStuff(otherVal1);

  if (!otherVal2.isValid()) {
    throw new Exception();
  }

  return "Stuff";
}
```

- 바뀐 코드에서 주목할 내용은 아래와 같다.
  - 코드의 인덴트가 1단계만 유지되어 선형적으로 쉽게 읽을 수 있다.
  - 함수의 마지막에 기대되는 정상 결과를 빠르게 찾을 수 있다.
  - 비 정상적인 인자나 상황에 대해 에러를 먼저 뱉음으로, 후에 실행되는 비지니스 로직에 버그가 발생할 확률을 낮춘다.
  - 실패 먼저라는 개념은 TDD에서 실패하는 테스트를 만들어라와 유사하다.
  - 바로 종료되어 버리므로, 의도되지 않은 상황에서 더 많은 코드의 실행을 막을 수 있다.
- early return 은 디자인 패턴과도 관련이 있다.
  - Fail Fast
    - Jim Shore와 Martin Fowler는 2004년에 [Fail Fast](https://www.martinfowler.com/ieeeSoftware/failFast.pdf) 개념을 고안했는데, 이 컨셉이 return early 규칙의 근간이 되었다.
  - Guard Clause
    - 함수 내부에서 check 해서 조건에 맞으면 return or 예외를 발생시켜 종료하는 방법이다.
    - Guard Clause를 사용하면 발생가능한 오류를 쉽게 식별할 수 있다.
  - Bouncer Pattern
    - Bouncer Pattern은 몇몇 조건을 만족할 때 함수 내부에서 return 하거나 예외를 발생시키는 방법이다.
    - validation code가 복잡하거나 다양한 시나리오가 존재할 떄 유용하고 early return을 보완할 수 있는 방법이다.

    ```js
    private void validateArgument1(SomeObject argument1, SomeObject argument2){
      if(!argument1.isValid()) {
        throw new Exception();
      }

      if(!argument2.isValid()) {
        throw new Exception();
      }
    }

    public void doStuff(String argument1, String argument2) {
      validateArgument1(argument1, argument2);

      // do more stuff
    }
    ```


## 레퍼런스

- https://medium.com/swlh/return-early-pattern-3d18a41bba8
