# DFS와 BFS

<img src="https://github.com/programmer-sjk/TIL/blob/main/images/algorithm/dfs-bfs.png" width="600">

- 위 그래프에 대해 DFS와 BFS 알고리즘으로 코드를 작성한다.

## 공통 그래프

- DFS, BFS 모두 그래프는 아래와 같이 동일하다.

```js
const graph = {
  A: ["B", "C"],
  B: ["A", "D"],
  C: ["A", "G", "H", "I"],
  D: ["B", "E", "F"],
  E: ["D"],
  F: ["D"],
  G: ["C"],
  H: ["C"],
  I: ["C", "J"],
  J: ["I"],
};
```

## DFS (Depth-First Search)

### 재귀적으로 출력하는 경우

```js
const graph = {
  A: ["B", "C"],
  B: ["A", "D"],
  C: ["A", "G", "H", "I"],
  D: ["B", "E", "F"],
  E: ["D"],
  F: ["D"],
  G: ["C"],
  H: ["C"],
  I: ["C", "J"],
  J: ["I"],
};

const DFS = (graph, startNode, visited) => {
  console.log(startNode);
  visited[startNode] = true;

  for (const node of graph[startNode]) {
    if (!visited[node]) {
      DFS(graph, node, visited);
    }
  }
};

DFS(graph, "A", {});
```

### 배열에 담아 출력하는 경우

```js
const graph = {
  A: ["B", "C"],
  B: ["A", "D"],
  C: ["A", "G", "H", "I"],
  D: ["B", "E", "F"],
  E: ["D"],
  F: ["D"],
  G: ["C"],
  H: ["C"],
  I: ["C", "J"],
  J: ["I"],
};

const DFS = (graph, startNode) => {
  const visited = [];
  let needVisit = [];

  needVisit.push(startNode);

  while (needVisit.length !== 0) {
    const node = needVisit.shift();
    if (!visited.includes(node)) {
      visited.push(node);
      needVisit = [...graph[node], ...needVisit];
    }
  }

  console.log(visited);
};

DFS(graph, "A");
```
