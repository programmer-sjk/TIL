# ECS (Elastic Container Service)

- ECS란 AWS에서 제공하는 **`컨테이너 오케스트레이션이다`**.
- **`오케스트레이션은`** 여러 객체(컨테이너, 프로세스, 서버)를 효율적으로 배포,관리해줄 수 있는 서비스이다.
  - 다수의 객체를 배포하거나 오토스케일링, 로드 밸런싱, 자동 재시작, 모니터링 등을 효율적으로 할 수 있는 방법들을 제공한다.
- **`ECS Fargate는`** serverless 환경을 제공하며 손쉽게 컨테이너들을 실행하는 환경을 제공한다.
- ECS는 Docker를 기반으로 하고 `EKS(Elastic Kubernetes Service)`는 쿠버네티스를 기본으로 한다.

## ECS vs EKS, EC2 vs fargate 선택

- EKS는 k8s에 대한 팀 이해도가 없으면 학습 곡선이 높다는 단점이 존재.
- ECS에 EC2를 사용할 경우 EC2에 대한 보안 패치를 직접적으로 관리해야 한다는 단점이 존재.
- ECS에 fargate를 사용할 경우 인스턴스에 대한 관리와 배포, 운영은 AWS 쪽에 맡김