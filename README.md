# Health Check API

API desenvolvida em Go para realizar verificações básicas de disponibilidade e funcionamento de uma aplicação.

O projeto foi desenvolvido com foco no aprendizado e aplicação prática de conceitos de desenvolvimento backend, Docker, CI/CD e DevSecOps.

## Tecnologias

* Go 1.25
* Gin
* Docker
* GitLab CI/CD
* GitHub
* golangci-lint
* gosec

## Funcionalidades

* Endpoint para verificação de saúde da aplicação
* Testes automatizados
* Análise estática de código
* Verificação de possíveis vulnerabilidades
* Build automatizado de imagem Docker
* Pipeline de CI/CD

## Pipeline CI/CD

O projeto utiliza GitLab CI/CD para automatizar etapas de validação e build.

A pipeline é dividida nas seguintes etapas:

```text
Test
  ↓
Lint
  ↓
Security
  ↓
Build
```

### Test

Executa os testes automatizados da aplicação:

```bash
go test ./...
```

### Lint

Utiliza o `golangci-lint` para analisar o código e identificar problemas relacionados à qualidade, possíveis erros e padrões de desenvolvimento.

### Security

Utiliza o `gosec` para realizar uma análise estática do código Go e identificar possíveis problemas de segurança.

### Build

Realiza o build da imagem Docker da aplicação após a conclusão das etapas anteriores.

## Execução local

### Pré-requisitos

Para executar o projeto localmente, é necessário ter instalado:

* Go 1.25 ou superior
* Docker, caso queira executar a aplicação em um container

### Clonando o projeto

```bash
git clone https://github.com/Gzcryy/health-check-api.git
cd health-check-api
```

### Instalando as dependências

```bash
go mod download
```

### Executando a aplicação

```bash
go run .
```

A aplicação será iniciada na porta `8080`.

## Executando com Docker

Para criar a imagem:

```bash
docker build -t health-check-api .
```

Para executar o container:

```bash
docker run -p 8080:8080 health-check-api
```

## Testes

Os testes podem ser executados utilizando:

```bash
go test ./...
```

## Análise de código

Para executar o `golangci-lint` utilizando Docker:

```bash
docker run --rm \
  -v "$PWD:/app" \
  -w /app \
  golangci/golangci-lint:v2.13.2-alpine \
  golangci-lint run
```

## Análise de segurança

Para executar o `gosec`:

```bash
docker run --rm \
  -v "$PWD:/app" \
  -w /app \
  securego/gosec:latest \
  gosec ./...
```

## Estrutura do projeto

```text
health-check-api/
├── .gitlab-ci.yml
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── main_test.go
└── README.md
```

## Objetivo

O objetivo deste projeto é aplicar, de forma prática, conceitos de desenvolvimento em Go e integração com ferramentas utilizadas em ambientes de DevOps e DevSecOps.

Além da implementação da API, o projeto busca demonstrar a utilização de testes automatizados, análise de código, verificação de segurança, containerização e integração contínua.

## Autor

Guilherme Matheus

GitHub: https://github.com/Gzcryy

GitLab: https://gitlab.com/Gzcryy/health-check-api

## Licença

Este projeto foi desenvolvido para fins de estudo e portfólio.
