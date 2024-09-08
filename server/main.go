package main

import (
	"context"
	"log"
	"net"

	pb "grpc-data-xfer/svckit-proto/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDataTransferServer
}

func (s *server) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	// Process the Send logic here
	return &pb.SendResponse{BytesTransferred: int64(len(req.Data))}, nil
}

func (s *server) SendAll(ctx context.Context, req *pb.SendAllRequest) (*pb.SendResponse, error) {
	var totalBytes int64
	for _, data := range req.Data {
		totalBytes += int64(len(data))
	}
	return &pb.SendResponse{BytesTransferred: totalBytes}, nil
}

func (s *server) Receive(ctx context.Context, req *pb.ReceiveRequest) (*pb.ReceiveResponse, error) {
	data := make([]byte, req.Size)
	// Simulate receiving data here
	return &pb.ReceiveResponse{Data: data}, nil
}

func (s *server) ReceiveAll(ctx context.Context, req *pb.ReceiveAllRequest) (*pb.ReceiveResponse, error) {
	var data []byte
	for _, size := range req.Sizes {
		data = append(data, make([]byte, size)...)
	}
	return &pb.ReceiveResponse{Data: data}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataTransferServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
