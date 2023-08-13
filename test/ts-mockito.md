# ts-mockito

## ts-mockito란

- 자바 진영에는 [mockito](https://site.mockito.org/)라고 불리는 mock 테스트를 편하게 도와주는 도구가 있다.
- ts-mockito는 자바 진영의 mockito에서 영감을 받아 typescript에서 동작하도록 만들어진 라이브러리다.

## ts-mockito 장점

- ts-mockito의 장점에 대해 이야기하려면 상대적으로 jest를 이용한 mock 테스트랑 비교해야 한다.

### 장점 1. IDE 지원

- jest
  - 아래는 jest를 사용하여 bannerRepository의 getCount 함수가 호출되면 bannerCount가 리턴되도록 stub한다.
  - getCount가 문자열로 작성하기 때문에 IDE의 자동완성의 혜택을 받을 수 없다.
  - getCount가 리팩토링되어 이름이 getTotalCount로 수정되었다고 하자. 네이밍만 수정하고 테스트를 돌리면 문자열 getCount가 수정되지 않아 실패한 테스트를 볼 수 있을 것이다.

```ts
jest.spyOn(bannerRepository, 'getCount').mockResolvedValue(bannerCount);
```

- ts-mockito
  - 아래는 ts-mockito를 사용해 동일한 코드를 보여준다.
  - 문자열이 아닌 실제 메서드를 사용하기해 자동완성, 리팩토링 지원을 받을 수 있다.

```ts
  when(bannerRepository.getCount(args)).thenResolve(bannerCount);
```

### 장점 2. 직관적인 코드

- 테스트 코드를 많이 작성해 봤음에도 jest 코드를 처음 볼 때 직관적으로 이해되지는 않았다. `jest.fn(), jest.spyOn(), mockImplementation()` 등 코드를 처음 보고 머리속에 물음표가 많았다.
- 반대로 ts-mockito는 처음 볼 때부터 when/then 구조로 어떤 코드인지 직관적으로 파악이 되었다.

  ```ts
    // bannerRepository.getCount가 args 인자로 호출되면 bannerCount를 리턴해라
    when(bannerRepository.getCount(args)).thenResolve(bannerCount);
  ```

### 장점 3. 불 필요한 코드

- jest를 사용한 mock 테스트를 쓴다면 회사마다 다를 수 있겠지만 최근 합류한 회사에서는 아래와 같은 테스트를 기존에 작성하고 있었다.
- jest로 bannerRepository를 mocking 하려면 따로 mock repository를 아래와 같이 생성해야 했다.

  ```ts
    // service 테스트에서 이 MockRepository를 사용
    export const MockBannerRepository = () => ({
      getCount: jest.fn().mockResolvedValue(0),
      getMany: jest.fn().mockResolvedValue([]),
      getOneById: jest.fn().mockResolvedValue(undefined),
      add: jest.fn().mockResolvedValue(undefined),
    })
  ```

  - 만약 새로운 repository에 새로운 메서드가 추가된다면 테스트를 위해 여기서도 가짜 함수를 선언해줘야 했다.
  - 그리고 왜 이렇게하지? 라는 생각을 하지만 일단은 적응해야 하니 타이핑을 했던 내가 있었다.
- ts-mockito를 사용하면 실제 repo기반으로 위와 같은 불필요한 코드없이도 테스트를 할 수 있다.

  ```ts
  beforeEach(async () => {
    await Test.createTestingModule({
      providers: [BannerService, BannerRepository],
    })
      .setLogger(new MockLogger())
      .compile();

    bannerRepository = mock(BannerRepository);
    service = new BannerService(instance(bannerRepository));
  });

  .
  .

  describe('getBannerCount()', () => {
    it('배너 개수를 구할 수 있다.', async () => {
      // given
      when(bannerRepository.getCount(type)).thenResolve(bannerCount);

      // when
      const result = await service.getBannerCount(type);

      // then
      verify(bannerRepository.getCount(type)).once();
      expect(result).toEqual(bannerCount);
    });
  });
  ```

- 위의 코드에서는 jest를 사용할 떄 처럼 테스트를 위해 불필요한 코드 없이 테스트에 필요한 코드만 작성하면 된다.

## ts-mockito 사용법

- 위에서 ts-mockito의 장점을 이야기 했으니 사용법에 대해 알아보자.
- 더 자세한 내용을 알고 싶다면 [github docs](https://github.com/NagRock/ts-mockito#readme)를 보고 확인할 수 있다.

### when

- 어떤 메서드가 어떤 인자값으로 호출되면 어떤 값을 반환할지를 정한다.

```ts
  // 리턴되는 결과가 동기라면 thenReturn
  when(bannerRepository.getBanner(1)).thenReturn(bannerCount);
  when(bannerRepository.getBanner('first')).thenReturn(bannerCount);

  // 리턴되는 결과가 비동기라면 thenResolve
  when(bannerRepository.getBanner({ id: 1})).thenResolve(bannerCount);
  when(bannerRepository.getBanner(anyFunction())).thenResolve(bannerCount);
  when(bannerRepository.getBanner(anyString())).thenResolve(bannerCount);
  when(bannerRepository.getBanner(anything())).thenResolve(bannerCount);
```

### verify

- 검증하는 로직으로 테스트 대상 함수에 대해 상호작용 테스트를 지원한다.

```ts
verify(bannerRepository.getBanner(1)).times(1); // 한 번 호출, 아래와 동일
verify(bannerRepository.getBanner(1)).once(); // 한 번 호출
verify(bannerRepository.getBanner(1)).twice(); // 두 번 호출
verify(bannerRepository.getBanner(1)).times(4); // 네 번 호출

verify(bannerRepository.getBanner(1)).atLeast(2); // 최소 두번 호출
verify(bannerRepository.getBanner(1)).atMost(4);  // 최대 네번 호출
verify(bannerRepository.getBanner(1)).never();    // 호출되지 않음

// remove 호출전에 getBanner 호출 검증
verify(bannerRepository.getBanner(1)).calledBefore(bannerRepository.remove(1));
// getBanner가 호출된 후에 updateTitle 호출 검증
verify(bannerRepository.getBanner(1)).calledAfter(bannerRepository.updateTitle('짱'));
```

### 전반적인 코드 sample

```ts
describe('BannerService', () => {
  let service: BannerService;
  let bannerRepository: BannerRepository;

  beforeEach(async () => {
    sentryService = MockSentryService();
    await Test.createTestingModule({
      providers: [
        BannerService,
        BannerRepository,
      ],
    })
      .compile();

    bannerRepository = mock(BannerRepository);
    service = new BannerService(instance(bannerRepository));
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('getBannerCount()', () => {
    it('배너 개수를 구할 수 있다.', async () => {
      // given
      const bannerCount = 10;
      const args = { bannerType: BannerType.MAIN_TOP };
      when(bannerRepository.getCount(args)).thenResolve(bannerCount);

      // when
      const result = await service.getBannerCount(args);

      // then
      verify(bannerRepository.getCount(args)).once();
      expect(result).toEqual(bannerCount);
    });
  });
})
```

## 마치며
