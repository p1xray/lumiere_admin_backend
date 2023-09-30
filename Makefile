.PHONY: dc build run test lint swagger

dc:
	docker-compose up --remove-orphans -- build

build:
	go build -v -o app.exe cmd/lumiere_admin_backend/main.go

run:
	go run cmd/lumiere_admin_backend/main.go

test:
	go test -v -race -timeout 30s ./...

lint:
	golangci-lint run

swagger:
	swag init --pd -g cmd/lumiere_admin_backend/main.go -ot "go,json" -p "camelcase"