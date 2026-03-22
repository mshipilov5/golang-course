.PHONY: proto swag build docker-up

proto:
	protoc -I api/proto \
		--go_out=api/proto --go_opt=paths=source_relative \
		--go-grpc_out=api/proto --go-grpc_opt=paths=source_relative \
		collectorv1/collector.proto

swag:
	swag init -g cmd/gateway/main.go -o docs --parseDependency --parseInternal

build:
	go build -o bin/collector ./cmd/collector
	go build -o bin/gateway ./cmd/gateway

docker-up:
	docker compose up --build
