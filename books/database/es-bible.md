# 엘라스틱서치 바이블

- [책 링크](https://www.yes24.com/Product/Goods/119719070)

## 1. ES 소개

- 데이터를 시각화하는 **키바나**와 ES에 색인할 데이터를 수집하고 변환하는 **로그 스태시**. 이를 합쳐서 **ELK 스택**이라 부른다.
- ES의 기본 컨셉
  - **검색 엔진**: ES는 기본적으로 검색 엔진으로 역색인을 사용하여 검색 속도가 빠르고 형태소 분석, 전문 검색이 가능하다.
  - **분산처리**: 데이터를 여러 노드에 분산 저장하며 검색이나 집계 작업을 수행할 수 있다.
  - **고가용성**: 클러스터를 구성하는 일부 노드에 장애가 발생해도 복제를 이용해 중단 없이 서비스를 지속할 수 있다.
  - **수평적 확장성**
  - **JSON 기반 REST API 제공**: ES는 JSON 형태의 문서를 저장, 색인, 검색하고 REST API를 사용한다.
  - **데이터 안정성**: 데이터 색인 요청후 200 OK 응답을 받았드면 그 데이터는 디스크에 저장됨을 보장한다.
  - **준실시간 검색**: ES가 역색인을 구성하고 검색이 가능해지기 까지 시간이 걸린다. 기본 설정은 1초 정도 걸리는데 이런 특성을 이해하고 있어야 한다.
  - **트랜잭션 지원되지 않음**: RDBMS와 달리 트랜잭션을 지원하지 않는다.
  - **조인 지원하지 않음**: ES는 기본적으로 조인을 염두에 두고 설계되지 않았다.

## 2. ES 기본 동작과 구조

### ES 기본 동작

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

- ES는 쿼리 전용 DSL을 제공한다. `_search`를 붙여 GET 메서드를 사용한다.

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

- 위에 대한 결과는 아래와 같은데 문서가 2개 검색되었고 `_score`에서 유사도 점수를 확인할 수 있다. 전통적인 RDBMS와는 동작방식이 상당히 다른걸 알 수 있다.

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

### ES 구조

- ES 기본 구조와 용어를 살펴보자
  - **문서**: ES가 저장하고 색인하는 json 문서
  - **인덱스**: 문서를 모아 놓은 단위
  - **샤드**: 인덱스는 그 내용이 여러 샤드로 분산 저장된다. 원본 역할을 담당하는 주 샤드와 복제본 샤드가 있다.
  - **\_id**: 인덱스 내 문서에 부여되는 고유한 구분자.
  - **노드**: ES 프로세스 하나가 노드를 구성한다.
    - **데이터 노드**: 샤드를 보유하고 샤드에 읽기/쓰기 작업을 수행
    - **마스터 노드**: 클러스터를 관리하는 역할을 하는 노드
  - **클러스터**: ES 노드가 여러개 모여 하나의 클러스터를 구성

### ES 내부 구조와 루씬

- ES는 문서를 색인하고 검색하는 `아파치 루씬`을 코어 라이브러리로 사용한다.
- **루씬 flush**
  - 문서의 색인 요청이 들어오면 루씬은 역색인을 생성한다. 최초 생성은 메모리 버퍼에 들어가며 주기적으로 디스크에 flush한다.
  - 루씬은 파일을 연 시점에 색인이 완료된 문서만 검색할 수 있고, 이후 변경사항이 발생하고 검색하고 싶으면 파일을 새로 열어야 한다.
  - 루씬이 **파일을 열고 변경사항이 적용된 새로운 인덱스를 얻는데** 이를 ES에서 **`refresh`** 라고 부른다.
- **루씬 commit**
  - 루씬의 **`flush는`** OS Page Cache에 데이터를 넘겨주는 것을 보장하지만 디스크에 파일이 기록되는 것을 보장하지는 않는다.
  - 따라서 루씬은 주기적으로 commit을 통해 Page Cache와 디스크의 싱크를 맞추며 ES의 flush 작업은 내부적으로 루씬의 commit 과정을 거친다.
  - ES의 **flush는 refresh 보다 훨씬 비용이 드는 작업**이다. 따라서 refresh와 마찬가지로 적절한 주기로 수행된다.
- **세그먼트**
  - 디스크에 기록된 파일들이 모이면 **세그먼트**라는 단위가 된다. 이 세그먼트가 루씬의 검색 대상이다.
  - 새로운 문서가 들어오면 새 세그먼트가 생성되고 문서가 삭제되면 삭제 플래그만 표시해둔다. 불변인 세그먼트를 무작정 늘릴 수 없기에 루씬은 중간에 적당히 세그먼트의 병합을 수행한다. 이 과정에서 삭제 플래그가 표시된 데이터를 실제로 삭제하기도 한다.
  - 세그먼트 병합은 비싼 작업이지만 병합 후에 검색 성능의 향상을 기대할 수 있다. 다만 명시적인 세그먼트 병합은 추가 데이터 색인이 없는게 보장될 때 수행되어야 한다.
- **루씬 인덱스와 ES 인덱스**
  - 여러 세그먼트가 모이면 하나의 루씬 인덱스가 되어 검색이 가능하다. ES 샤드는 루씬의 인덱스 하나를 래핑한 개념이다.
  - ES 샤드가 여러 개 모이면 ES 인덱스가 된다. ES에서는 여러 샤드에 있는 문서를 모두 검색할 수 있다.
- **translog**
  - ES에 색인된 문서는 루씬 commit까지 완료되어야 디스크에 안전하게 기록되지만 모든 요청에 대해 루신 commit을 하는 작업은 비싸다. 그렇다고 주기적으로 commit하면 장애가 발생할 때 데이터 유실의 우려가 있기 때문에 ES 샤드는 모든 작업마다 translog에 로그를 기록한다.
  - translog는 색인, 삭제 작업이 수행된 직후에 기록되고 ES는 장애가 발생한 뒤 샤드 복구 단계에서 translog를 읽어 유실된 데이터를 복구한다.
  - 이 translog가 커지면 샤드 복구에 시간이 오래 걸리기 때문에 translog의 크기를 적절히 유지해 줄 필요가 있다. 참고로 ES flush는 루씬의 commit을 수행하고 새로운 translog를 만드는 작업이다. ES flush는 백그라운드에서 주기적으로 수행되며 translog의 크기를 적절한 수준으로 유지한다.

## 3. 인덱스 설계

### 3.1 인덱스 설정

- 인덱스 설정 조회

  ```elixir
    GET my_index/_settings

    // 응답
    {
    "my_index": {
      "settings": {
        "index": {
          "routing": {
            "allocation": {
              "include": {
                "_tier_preference": "data_content"
              }
            }
          },
          "number_of_shards": "1",
          "provided_name": "my_index",
          "creation_date": "1697937794587",
          "number_of_replicas": "1",
          "uuid": "AtAfbz1tROSfK4310S17AQ",
          "version": {
            "created": "8040299"
          }
        }
      }
    }
  }
  ```

- 존재하지 않는 인덱스에 문서 색인 요청을 하면 ES는 인덱스를 자동으로 생성한다. 자동 생성된 인덱스가 어떤 기본값을 가지는지 알아보자.

  - number_of_shards
    - 인덱스가 데이터를 몇 개의 샤드로 쪼갤 것인지 지정하는 값이다. 한 번 지정하면 reindex 동작을 통해 인덱스를 통째로 색인하는 작업을 수행하지 않는 한 바꿀수 없다.
    - 샤드 개수를 어떻게 지정하느냐는 ES 전체의 성능에도 큰 영향을 미친다. 클러스터에 샤드 숫자가 너무 많으면 색인 성능이 감소한다. 반대로 인덱스당 샤드 숫자를 적게 지정하면 샤드 하나의 크기가 너무 커진다. 샤드 하나의 크기가 크면 샤드 복구에 너무 많은 시간이 소요된다.
  - number_of_replicas

    - 주 샤드 하나당 복제본 샤드를 몇 개 둘 것인지 설장한다. 인덱스 생성 후에도 동적으로 변경이 가능하다.
    - 아래와 같이 0으로 지정하면 복제본 샤드 없이 주 샤드만 둔다.

    ```elixir
      PUT my_index/_setting
      {
        "index.number_of_replicas": 0
      }
    ```

  - refresh_interval

    - ES가 인덱스를 대상으로 refresh를 얼마나 자주 수행할지 지정한다. 인덱스에 색인된 문서는 refresh 되어야 검색 대상이 되기 때문에 중요한 설정이다. 아래와 같이 값을 지정할 수 있다.

    ```elixir
      PUT my_index/_setting
      {
        "index.refresh_interval": "1s"
      }
    ```

    - 값을 -1로 지정하면 refresh를 수행하지 않는다. 기본 값은 1초 마다 refresh를 수행하며 30초 이상 검색 쿼리가 들어오지 않으면 검색 쿼리가 들어올 때까지 refresh를 수행하지 않는다. 이 대기시간은 `index.search.idel.after` 설정으로 변경이 가능하며 `index.refresh_interval` 값을 null로 업데이트 하면 인덱스를 refresh_interval 값을 설정하지 않은 상태로 업데이트 할 수 있다.

- 위에서 사용한 my_index는 문서 색인 요청을 통해 자동으로 생성한 인덱스다. 이렇게 생성되면 인덱스 설정이 모두 기본값으로 지정되기 때문에 실제 운영 환경에선 적절하지 않다. 인덱스를 수동으로 생성/삭제 하는 방법을 알아보자.

  ```elixir
    // 인덱스 수동 생성
    PUT my_index
    {
      "settings": {
        "number_of_shards": 2,
        "number_of_replicas": 2
      }
    }

    // 인덱스 삭제
    DELETE my_index
  ```

### 3.2 매핑과 필드 타입

- 매핑은 문서가 인덱스에 어떻게 색인되고 저장되는지 정의하는 부분이다.
- 아래와 같이 문서가 색인될 때 기존에 매핑 정보를 가지고 있지 않던 새로운 필드가 들어오면 ES는 자동으로 문서의 필드 타입을 지정해서 매핑 정보를 생성한다.

  ```elixir
    // 문서 색인
    PUT my_index2/_doc/1
    {
      "title": "hello world",
      "views": 1234,
      "public": true,
      "point": 4.5,
      "created": "2019-01-17T14:05:01.234Z"
    }

    // 응답
    {
    "my_index2": {
      "aliases": {},
      "mappings": {
        "properties": {
          "created": {
            "type": "date"
          },
          "point": {
            "type": "float"
          },
          "public": {
            "type": "boolean"
          },
          "title": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "views": {
            "type": "long"
          }
        }
      },
      "settings": {...}
    }
  }
  ```

#### 동적 매핑 vs 명시적 매핑

- ES가 자동 생성하는 매핑을 동적 매핑이라고 하고 사용자가 직접 매핑을 지정해 주는 방법을 명시적 매핑이라고 부른다.
- 아래와 같이 인덱스를 생성할 때 직접 매핑 정보를 지정할 수 있다.

  ```elixir
  mapping_test
  {
    "mappings": {
      "properties": {
        "createdDate": {
          "type": "date",
          "format": "strict_date_time || epoch_millis"
        },
        "keywordString": {
          "type": "keyword"
        },
        "textString": {
          "type": "text"
        }
      }
    },
    "settings": {
      "number_of_replicas": 1,
      "number_of_shards": 1
    }
  }
  ```

- 중요한 점은 필드 타입을 포함한 매핑설정은 한 번 지정되면 변경이 불가능하다는 점이다. 따라서 서비스 설계와 데이터 설계를 할 때는 매우 신중해야 한다. 서비스 운영 환경에서 대용량의 데이터를 처리해야 할 때는 명시적 매핑을 지정해서 인덱스를 운영해야 한다. 매핑을 어떻게 지정하냐에 따라 성능의 차이도 크다. 동적 매핑은 유연한 운영을 가능하게 해 주지만 그럼에도 명시적으로 매핑을 지정하는 것이 좋다.
- 이미 인덱스가 생성된 경우에도 매핑 정보를 아래와 같이 추가할 수 있다.

  ```elixir
    PUT mapping_test/_mapping
    {
      "properties": {
        "longValue": {
          "type": "long"
        }
      }
    }
  ```
