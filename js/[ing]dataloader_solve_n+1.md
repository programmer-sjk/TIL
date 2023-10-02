# typeorm에서 dataloader를 활용한 GraphQL의 N+1 문제 해결

- GraphQL의 resolve field를 쓸 때 N+1 문제를 어떻게 해결할 수 있는지 살펴본다.
- ORM의 eager & lazy 각 장단점과 N+1 개념은 알고 있다고 가정하고 작성한다.

## typeorm과 N+1

- spring 진영의 ORM인 JPA 경우 OneToMany 관계의 기본 fetch type이 Lazy Loading이라 ORM을 쓸 때 N+1 문제를 쉽게 접할 뿐 더러 eager와 lazy에 따라 쿼리가 어떻게 수행되고 어떻게 N+1 문제를 해결하는지 아는 것이 중요하다.
- typeorm은 기본 fetch type이 eager나 lazy가 아니고 보통 find 메소드의 relations 옵션이나 쿼리 빌더의 조인을 사용하기 때문에 실무에서 N+1 문제를 접할 일이 상대적으로 드물긴 하다.
- typeorm에서 N+1은 언제 발생하는지 알아보자.

  - 아래와 같이 movie 엔티티의 연관관계가 있는 reviews에 lazy 옵션을 추가한다.

  ```ts
  @Entity()
  export class Movie {
    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    title: string;

    @OneToMany(() => Review, (review) => review.movie, { lazy: true })
    reviews: Review[];
  }
  ```

  - 테스트 상의 DB에는 movie가 3건이 저장되어 있는 상태이다. movie에 대해 reviews를 아래 코드처럼 접근해보자.

  ```ts
  async lazy() {
    const movies = await this.movieRepository.find();
    for (const movie of movies) {
      const reviews = await movie.reviews;
    }
  }
  ```

  - 아래처럼 쿼리가 수행되며 movie를 전체 조회하는 쿼리와, 각 movie의 review를 구하기 위해 movie 개수 별로 쿼리가 나가므로 N+1 문제를 확인할 수 있다.

  ![](../images/js/n+1_query.png)

## resolve field란?

## dataloader란?
