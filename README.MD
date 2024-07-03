# Golang RESTful API

Este projeto é uma API RESTful criada em Golang utilizando os frameworks Gin e GORM, com PostgreSQL como banco de dados e Docker para containerização.

## Pré-requisitos

- Golang
- Docker
- Docker Compose

## Configuração

### 1. Clone o repositório

```bash
git clone https://github.com/carlosmoysesai/api-golang.git
cd api-golang
```

### 2. Instale as dependências
Execute os seguintes comandos para instalar o Gin, GORM e o driver do PostgreSQL:

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### 4. Inicie os containers Docker
```bash
docker-compose up -d --build
```
### 5. Iniciar api Golang
```bash
go run main.go
```
