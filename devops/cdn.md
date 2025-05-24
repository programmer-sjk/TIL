# CDN (Content Delivery Network)

- **`CDN이란`** 컨텐츠들을 엔드 유저에게 빠르게 제공하기 위한 방법을 제공한다.
- **`원본 컨텐츠를 가진 서버에서 데이터를 받기에는 물리적으로 오래 걸리기 때문에`** 세계 각지에 여러 서버를 두고 원본 데이터를 캐시한다. 사용자들은 데이터를 멀리있는 원본 데이터를 가진 서버에서 받는게 아니라, 가장 가까운 서버로부터 데이터를 제공받는다.
  - 물리적으로 멀리 떨어진 서버에서 영상을 다운로드 하거나 용량이 큰 이미지를 받기에 시간이 오래 걸릴 수 있다.
- CDN을 통해 페이지 로딩 시간을 단축하고 서버의 대역폭 비용을 절감하며 컨텐츠 가용성과 보안에 이점을 얻을 수 있다.

## CloudFront

- AWS는 **`CloudFront라는`** CDN 서비스를 제공한다.
- CloudFront에는 `Edge location과 Regional edge caches`가 있다. Edge location은 사용자와 지역적으로 가까이 위치하여 빠른 데이터를 제공한다. 캐시된 데이터가 빈번히 요청되지 않아 cache에서 제거되면 Edge location은 Regional edge caches로 데이터를 요청한다. Regional edge caches은 훨씬 더 큰 용량을 가지고 있어 오래동안 캐시에 남게 된다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/cloud/cdn-edge-location.png" width="600">

- 2023년 7월을 기준으로 450개가 넘는 Edge location 서버가 있으며 13개의 Regional edge caches 서버를 두고 있다.

  <img src="https://github.com/programmer-sjk/TIL/blob/main/images/cloud/world-cloudfront.png" width="600">
