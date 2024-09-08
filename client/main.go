package main

import (
	"context"
	"log"
	"time"

	pb "grpc-data-xfer/svckit-proto/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewDataTransferClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Send(ctx, &pb.SendRequest{Data: []byte("client-00000000000000000000001-data")})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Bytes Transferred: %d", res.BytesTransferred)
}
