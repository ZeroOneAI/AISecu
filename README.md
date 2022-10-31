# AI.Secu
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)
***

## Features

- AI.Secu Features
  - Scan an image by automatically detecting the image being pushed to the image repository
  - Continuous tracking of updated vulnerabilities
  - Point of Distinction

    | Other existing services                                                                                              | AI.Secu                                                                               | Current Support |
    |---------------------------------------------------------------------------------------|---------|-----------------|
    | Services supported by existing registry only support scanning of images uploaded to that registry                    | Manage vulnerabilities in one place for images stored in multiple different registers | X (only docker hub for now, but expandable) |
    | Most services offer only one-time scans of images (need to scan again to see updated vulnerabilities after scanning) | Once scanned, vulnerabilities can be tracked continuously                             | O |
    | SBOM format output not supported for components in the image                                                         | SBOM format output support for components in the image                                | X |
    | Unable to receive results that have already been scanned into the image as input | Can receive result (CycloneDx format) that have already been scanned as input         | X |

## Installation
- ```shell
  git clone https://github.com/ZeroOneAI/AISecu.git
  cd AISecu/build/
  kubectl create ns secu
  make deploy
    ```
  
## Project Timeline

| Date | Action |
|------|--------|
| 22.05.26 | Inception Meeting |
| 22.05.31 | Concept & Feature Design |
| 22.06.30 | Initial Contept Implementation | 
| 22.07.30 | Minimum Viable Product Release for Alpha Tester |


## Roadmap
- SBOM format output for image components
- Processing inputs in SBOM format
- Image Permission Control
- webhook functionality for CI/CD tools
  - It's not just that the image is updated and automatically distributed, but if you check the security vulnerabilities and give them permission manually, the deployment will proceed automatically afterwards
  - Provide a Audit Log of who allowed the image
- Queue image scan
- RBAC User Permission Control

## Authors and acknowledgment
- JunHo Song
    - Initiator
- MyeongSuk Yoon
    - Initiator
    - Developer

## More Documentation
- [Public](docs/external/index.md)

## Get in Touch
For issues, suggentions, contributors and supportors 
- `ai.secu@zeroone.ai`
