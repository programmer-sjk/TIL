# 함수의 리턴타입과 파라미터 타입

```ts
function getProduct({ id: number, ownerId: number }) {
  return "getProduct";
}

type returnType = ReturnType<typeof getProduct>;
type parameterType = Parameters<typeof getProduct>;
```
