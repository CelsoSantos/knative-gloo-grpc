package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/CelsoSantos/knative-gloo-grpc/api"
)

var (
	ctx = context.Background()
)

type hubService struct{}

func (c *hubService) Submit(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received Submission")

	input := in.Document

	if len(input) > 0 {
		log.Printf(input)
	} else {
		log.Printf("Invalid Input")
	}

	return &pb.Response{Status: 200, Document: "Input Value: " + input}, nil
}

func main() {

	// create a listener on TCP port 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// create server
	pb.RegisterHubServiceServer(grpcServer, &hubService{})
	reflection.Register(grpcServer)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
