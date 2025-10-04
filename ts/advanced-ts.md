# 함수의 리턴타입과 파라미터 타입

```ts
function getProduct({ id: number, ownerId: number }) {
  return "getProduct";
}

type returnType = ReturnType<typeof getProduct>;
type parameterType = Parameters<typeof getProduct>;
```

# as const

```ts
// 아래는 각 타입이 string으로 추론됨
const color = {
    Red: 'red',
    Blue: 'blue',
    Green: 'green'
}

// 아래는 각 타입이 red, blue, green으로 추론됨
const color = {
    Red: 'red',
    Blue: 'blue',
    Green: 'green'
} as const

keyof typeof color // "Red" | "Green" | "Blue"

// 이렇게 하면 color에 key가 추가되더라도 계속 사용이 가능
typeof color[keyof typeof color] // "red" | "green" | "blue"
```
