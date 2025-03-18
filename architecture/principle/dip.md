# DIP (Dependency Inversion Principle)

- DIP는 SOLID 원칙의 마지막 D에 해당하는 원칙이다.
- 고수준 모듈이 저수준 모듈에 의존해서는 안된다를 내포하며 어떤 객체에 대한 의존성을 가질 경우 구체적인 클래스에 의존하기 보다 추상화(추상 클래스, 인터페이스)에 의존해야 한다는 의미를 내포한다.
- 고수준 모듈이 저수준 모듈에 의존할 경우, **`저수준 모듈의 구체적인 구현이 변경될 때 마다 고수준 모듈도 변경이 되는 문제가 있기 때문이다`**.

## 코드로 보는 DIP

- 커피샵은 오픈 후 아메리카노만 판매해서 처음에는 아메리카노만 판매하고 있었다.
- 하지만 시간이 지나 바닐라 라떼, 모카 등등도 판매하게 된다면 CoffeeShop 클래스는 계속 수정되어야 한다.

```ts
export class CoffeeShop {
  private readonly coffee: Americano;

  constructor() {
    this.coffee = new Americano();
  }

  showPrice() {
    return this.coffee.price;
  }
}

export class Americano {
  readonly price = 4000;
}
```

- 위 코드에서 변경이 빈번하게 발생하는 **`원인은 CoffeeShop 클래스가 구체적인 클래스에 의존하고 있기 때문이다`**.
- DIP를 활용해서 다른 커피가 추가되어도 변경이 없도록 수정해보자.
- 아래처럼 Coffee 인터페이스를 만들고 저수준 모듈이 인터페이스를 구현하도록 강제한다.

```ts
interface Coffee {
  price: number;
}

export class Americano implements Coffee {
  readonly price = 4000;
}

export class VanillaLatte implements Coffee {
  readonly price = 5000;
}

export class MochaLatte implements Coffee {
  readonly price = 5500;
}
```

- 그 후 CoffeeShop에는 생성자를 통해 인터페이스를 받는다.

```ts
export class CoffeeShop {
  private readonly coffee: Coffee;

  constructor(coffee: Coffee) {
    this.coffee = coffee;
  }

  showPrice() {
    return this.coffee.price;
  }
}
```

- 그리고 CoffeeShop을 호출하는 쪽에서 구체적인 클래스를 전달한다.

```ts
const coffeShop = new CoffeeShop(new Americano());
console.log(coffeShop.showPrice);
```

- 결과적으로 상위 모듈이 **`구체적인 하위 모듈이 아닌 인터페이스에 의존함으로써 결합도를 낮출 수 있었다`**.
- 또한 DIP 원칙에 맞게 의존성이 상위에서 하위 모듈로 가는게 아닌, 하위 모듈과 상위 모듈이 모두 인터페이스에 의존하도록 의존성의 방향이 바뀜을 확인할 수 있다.

## service <-> repository
