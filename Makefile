generate-proto:
	make proto
	make grpc

proto:
	protoc --proto_path=internal ./internal/handler/grpc/*.proto --go_out=.

grpc:
	protoc --proto_path=internal ./internal/handler/grpc/*.proto --go_out=. --go-grpc_out=.

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/grpc_api/main.go

run:
	docker run

docker-build:
	docker build -t grpc-example-with-go .

docker-run:
	docker run -d --name grpc-example-with-go -p 50051:50051 grpc-example-with-go
