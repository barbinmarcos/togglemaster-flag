# README — `togglemaster-flag`

```md id="w3p2jk"
# ToggleMaster Flag Service

Microsserviço responsável pelo gerenciamento de Feature Flags.

## Responsabilidades

- Criação de flags
- Ativação/desativação de funcionalidades
- Versionamento de flags
- Controle de rollout

## Stack

- Golang
- Docker
- Kubernetes
- GitHub Actions
- Amazon ECR
- Amazon EKS

## Execução local

```bash
go run cmd/main.go
Endpoint de Health Check
GET /health
Docker

Build:

docker build -t togglemaster-flag .

Run:

docker run -p 8081:8081 togglemaster-flag
CI/CD

Pipeline DevSecOps com:

Build
Unit Tests
GolangCI-Lint
Gosec
Trivy
Docker Build
Push para Amazon ECR
Deploy

Deploy automatizado via GitOps utilizando ArgoCD.


