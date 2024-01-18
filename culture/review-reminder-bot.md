# 코드 리뷰 리마인더 봇

## 리뷰 리마인더의 필요성

- MSA로 코드가 여러 Repository에서 관리되는 환경에서 모든 PR을 리뷰해야 하는 환경에 있다.
- 리뷰를 한 적이 없는 PR이라면 [review-requested](https://github.com/pulls/review-requested) 페이지에서 확인이 가능하지만 comment를 달았던 PR은 이 페이지에서 확인되지 않는다.
- 매일 출근해서 리뷰를 하는데 두 단계로 나뉘어진다.
  - [review-requested](https://github.com/pulls/review-requested) 페이지에서 나에게 요청은 왔으나 리뷰를 한 적 없는 PR을 리뷰한다.
  - github 각 Repository에서 pulls 페이지를 들어가 comment를 단 적이 있는 PR을 찾아 리뷰한다.
- 매일 하는 리뷰인데 이 과정들이 너무 귀찮아서 매일 아침 리뷰 목록을 알려주는 봇을 만들기로 결심했다.
