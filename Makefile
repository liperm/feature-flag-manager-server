run:
	go run cmd/ff-manager-server/*

compile-proto:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=api/ --go-grpc_out=api/

up:
	docker compose up -d
