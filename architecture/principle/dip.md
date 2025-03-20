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

- 위 코드에서는 새로운 커피를 판매하더라도 CoffeeShop을 호출하는 클라이언트 쪽에서 추가되는 커피만 넣으면 되고 고수준 모듈인 CoffeeShop은 변경되지 않는다.
- 결과적으로 상위 모듈이 **`구체적인 하위 모듈이 아닌 인터페이스에 의존함으로써 결합도를 낮출 수 있었다`**.
- 또한 DIP 원칙에 맞게 의존성이 상위에서 하위 모듈로 가는게 아닌, 하위 모듈과 상위 모듈이 모두 인터페이스에 의존하도록 의존성의 방향이 바뀜을 확인할 수 있다.

## service <-> repository

- Layer 아키텍처에서 서비스와 레포지터리를 비교해보면 서비스는 고수준 모듈에 속하고 레포지터리는 DB와 통신하는 저수준 모듈에 해당한다.
- 보통 NestJS에서 서비스 생성자에서 레포지터리를 받아 사용하고는 한다. 이때 서비스에서는 주입받은 레포지터리가 제공하는 수 많은 API들을 사용할 수 있다. **`만약 ORM이 바뀌거나 DB가 바뀌게 되면 비지니스 로직들이 저수준 모듈인 레포지터리 변경에 의해 모두 영향받게 된다`**. 여기에 DIP를 적용해보자.
- 아래는 Nestjs에서 볼 수 있는 서비스에 레포지터리를 생성자로 주입받는 코드이다.

```ts
@Injectable()
export class UserService {
  constructor(private readonly userRepository: UserRepository) {}

  async find(id: number) {
    return this.userRepository.findOneBy({ id });
  }
}
```

- findOneBy 메서드는 TypeOrm 레포지터리가 제공하는 메서드로 만약 ORM을 변경한다면 service 코드가 수정된다.
- DIP 적용을 위해 인터페이스를 선언한다.

```ts
export interface IUserRepository {
  findOneBy(id: number): Promise<User>;
}
```

- 그리고 IUserRepository 인터페이스를 구현하는 어댑터 클래스를 추가한다.

```ts
@Injectable()
export class UserRepositoryAdaptor implements IUserRepository {
  constructor(private readonly userRepository: UserRepository) {}

  async findOneBy(id: number) {
    return this.userRepository.findOneBy({ id });
  }
}
```

- 어댑터 클래스는 생성자로 userRepository를 주입받아 기능을 사용한다. 이제 서비스가 이 어댑터를 사용하면 어댑터 클래스가 제공하는 메서드만 사용할 수 있지 userRepository가 제공하는 수 많은 API는 숨겨지게 된다.
- 서비스에는 생성자로 IUserRepository 인터페이스를 주입해야한다. Spring과 달리 **`NestJS는 인터페이스를 Provider로 제공하려면 user 모듈에 아래와 같이 provide 설정을 추가해야 한다`**.

```ts
@Module({
  providers: [
    UserService,
    UserRepository,
    { provide: "IUserRepository", useClass: UserRepositoryAdaptor },
  ],
})
export class UserModule {}
```

- 이제 아래와 같이 service에 Inject 데코레이터를 통해 인터페이스를 주입받아 사용한다.

```ts
@Injectable()
export class UserService {
  constructor(
    @Inject("IUserRepository") private readonly userRepository: IUserRepository
  ) {}

  async find(id: number) {
    return this.userRepository.findOneBy(id);
  }
}
```

## 정리하며

- DIP를 활용해 서비스가 실제 구현체인 repository가 아닌 인터페이스에 의존하도록 설계해봤다.
- 다만 위 코드에서 알 수 있듯, **`module, interface, adaptor 클래스 등 추가적인 코드 작업이 필요하다`**. 실제 운영 환경에서 ORM이 바뀌거나 DB가 바뀌는 경우는 정말 드물기에 팀과 프로젝트 미래를 생각해 적용할지 말지 고민해 볼 필요는 있어 보인다.
