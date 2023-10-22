# 엘라스틱서치 바이블

- [책 링크](https://www.yes24.com/Product/Goods/119719070)

## 1. ES 소개

- 데이터를 시각화하는 키바나와 ES에 색인할 데이터를 수집하고 변환하는 로그 스태시. 이를 합쳐서 ELK 스택이라 부른다.
- ES의 기본 컨셉
  - 검색 엔진: ES는 기본적으로 검색 엔진으로 역색인을 사용하여 검색 속도가 빠르고 형태소 분석, 전문 검색이 가능하다.
  - 분산처리: 데이터를 여러 노드에 분산 저장하며 검색이나 집계 작업을 수행할 수 있다.
  - 고가용성: 클러스터를 구성하는 일부 노드에 장애가 발생해도 복제를 이용해 중단 없이 서비스를 지속할 수 있다.
  - 수평적 확장성
  - JSON 기반 REST API 제공: ES는 JSON 형태의 문서를 저장, 색인, 검색하고 REST API를 사용한다.
  - 데이터 안정성: 데이터 색인 요청후 200 OK 응답을 받았드면 그 데이터는 디스크에 저장됨을 보장한다.
  - 준실시간 검색: ES가 역색인을 구성하고 검색이 가능해지기 까지 시간이 걸린다. 기본 설정은 1초 정도 걸리는데 이런 특성을 이해하고 있어야 한다.
  - 트랜잭션 지원되지 않음: RDBMS와 달리 트랜잭션을 지원하지 않는다.
  - 조인 지원하지 않음: ES는 기본적으로 조인을 염두에 두고 설계되지 않았다.

## 2. ES 기본 동작과 구조

- kibana의 [dev tools]에 들어가 간단한 API를 확인해본다.

  ```Elixir
  # 인덱스 이름이 my_index에 1번 문서를 색인
  PUT my_index/_doc/1
  {
    "title": "제목",
    "views": 10
  }

  # _id 지정없이 인덱스 색인 (ES가 _id를 자동 생성)
  POST my_index/_doc
  {
    "title": "제목",
    "views": 10
  }

  # 1번 문서 조회
  GET my_index/_doc/1

  # 문서 업데이트
  POST my_index/_update/1
  {
    "doc": {
      "title": "업데이트 된 제목"
    }
  }

  # 문서 삭제
  DELETE my_index/_doc/1
  ```

- ES는 쿼리 전용 DSL을 제공한다. _search를 붙여 GET 메서드를 사용한다.

  ```elixir
    GET my_index/_search
    {
      "query": {
        "match": {
          "title": "world"
        }
      }
    }
  ```

- 위에 대한 결과는 아래와 같은데 문서가 2개 검색되었고 _score에서 유사도 점수를 확인할 수 있다. 전통적인 RDBMS와는 동작방식이 상당히 다른걸 알 수 있다.

  ```elixir
    {
      "took": 12,
      "timed_out": false,
      "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
      },
      "hits": {
        "total": {
          "value": 2,
          "relation": "eq"
        },
        "max_score": 0.13353139,
        "hits": [
          {
            "_index": "my_index",
            "_id": "O3D7VIsBd1YI2lqiQb6F",
            "_score": 0.13353139,
            "_source": {
              "title": "hello world",
              "views": 1234,
              "public": true
            }
          },
          {
            "_index": "my_index",
            "_id": "1",
            "_score": 0.13353139,
            "_source": {
              "title": "high world",
              "views": 1234,
              "public": true
            }
          }
        ]
      }
    }
  ```
