# GRPC Go Data Xfer
Go 1.23 gRPC Client and Server Data Transfer Service using Buf Build gRPC Generator


## Installing BufBuild.Buf Kit

Prerequisites to install `BufBuild.Buf Kit`.
```shell
# Install the Go protocol buffer compiler plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install the gRPC plugin for Go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Finally install the `buf` generator CLI.
```shell
go install github.com/bufbuild/buf/cmd/buf@latest
````


## Running the Project

1. Generate Protobufs

Run `make proto` to generate the protobuf files using `buf`.

2. Build the Server and Client

Run `make server` and `make client`to build the exe targets.

3. Build the Docker Images

Run `make docker-server` and `make docker-client`to build the Docker images.

4. Run the Server

Run `make run-server` to start the server.

5. Run the Client

Run `make run-client` to start and run the client and execute gRPC requests.
