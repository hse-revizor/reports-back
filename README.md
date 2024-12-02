The service for getting reports in Go

# Usage

I. Starting
```
go run main.go
```

II. Generating docs
```
swag init -g ./cmd/main.go -o ./cmd/docs
```

III. Viewing docs
http://<host:port>/docs/index.html

# Development

## Stack

- Golang
- Gin
- Gin swagger (OpenAPI docs)
- Docker

## Configuring

.env.development file is used for local development. The following variables may be specified:
- DB_HOST
- DB_PORT
- DB_USER
- DB_PASSWORD
- DB_NAME

## Running locally

```
docker build -t revizor_reports_back:1.0 .
docker compose up -d
```
