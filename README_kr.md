# AI.Secu

***

## 특징
- AI.Secu 기능
  - 레포지토리에 이미지 푸쉬 시 자동으로 이미지 스캔
  - 취약점 업데이트 되는 것을 지속적으로 트랙킹
- 차별점
  - 기존 레지스트리들과의 차별점

    | 기존의 레지스트리에서 이미 지원 중인 기능                        | AI.Secu 만의 차별점                                                           | 현재 지원 여부 |
    |--------------------------------------------------------------------------|-----------------------------------------------------------------------------------|-----------------------|
    | 기존 레지스트리에서 해당 레지스트리에 업로드 된 이미지 스캐닝 및 트랙킹 기능 지원 | 서로 다른 여러 레지스트리에 저장된 이미지들에 대한 트랙킹을 한 곳에서 관리                               | O |
    | 이미지에 대한 일회성 스캔 (스캔 이후 업데이트되는 취약점에 대해서 확인 하기 위해서는 다시 스캔을 해야 됨) | 한 번의 스캔을 통해 지속적으로 트랙킹 가능                                                 | O |
    | 이미지의 구성 요소에 대한  sbom 형식 출력 X                    | 이미지의 구성 요소에 대한 sbom 형식(CycloneDx) 출력 지원                    | X |
    | 이미 스캔된 정보를 입력으로 받을 수 없음                                       | 이미 스캔된 정보(CycloneDx 형식)를 입력으로 받을 수 있음 (image를 직접 주지 않아도 보안 사항을 관리할 수 있음) | X |

## 설치 방법
- ```
  git clone https://github.com/ZeroOneAI/AISecu.git
  cd AISecu/build/
  kubectl create ns secu
  make deploy
  ```

## Project 타임라인

| 날짜       | 내용            |
|----------|---------------|
| 22.05.26 | KickOff       |
| 22.05.31 | V10 구상 중      |
| 22.06.30 | V10 최소 기능 구현  |

## Roadmap
- 이미지 구성 요소에 대한 SBOM 형식 출력
- SBOM 형식의 인풋 처리
- 이미지 권한 허가 통제 (w. [Notary](https://github.com/notaryproject/notary))
- 권한 허가에 대한 ci/cd 툴들에 대한 web hook 기능
  - 단순히 이미지가 업데이트 되었다고 자동 배포되는 것이 아니라 보안적인 취약점들 확인 후 수동으로 허가를 해주면 이후 자동으로 배포 진행을 위함
  - 누가 해당 이미지를 허용했는 지 Audit Log 제공
- OAuth 및 다양한 API 제공

## 더 생각해볼 것
- trivy DB 다운로드 반복 하는 부분 없애기
  - 공통된 스토리지 공유한다해도 동시에 다운받을 시 문제 생길 가능성 있음
- 이미지 스캔 큐잉
- RBAC 권한 관리

## Authors and acknowledgment
- JunHo Song
  - Initiator
- MyeongSuk Yoon
  - Initiator
  - Developer


## More Documentation
- [자료 더보기](docs/external/index.md)

## Support
문제 발생 시 아래 이메일로 연락 혹은 이슈 발행 부탁드립니다.
- `ai.secu@zeroone.ai`
