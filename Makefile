.PHONY: proto server client

proto:
	buf generate

server:
	CGO_ENABLED=0 GOOS=linux go build -o ./server ./server

client:
	CGO_ENABLED=0 GOOS=linux go build -o ./client ./client

docker-server:
	docker build -f Dockerfile.server -t grpc-server .

docker-client:
	docker build -f Dockerfile.client -t grpc-client .

run-server:
	docker run -p 50051:50051 grpc-server

run-client:
	docker run grpc-client
