.PHONY: fmt lint test build run-echo run-grpc migrate-up migrate-down docker-up docker-down

fmt:
	go fmt ./...
	goimports -w .

lint:
	golangci-lint run ./...

test:
	go test ./... -v

build:
	go build -o bin/echo-server ./cmd/echo-server
	go build -o bin/grpc-server ./cmd/grpc-server

run-echo:
	go run ./cmd/echo-server

run-grpc:
	go run ./cmd/grpc-server

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker-compose build

clean:
	rm -rf bin/

deps:
	go mod tidy
	go mod download