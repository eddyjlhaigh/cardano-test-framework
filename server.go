package main

import (
	"log"
	"net"

	"github.com/eddyjlhaigh/cardano-test-framework/node"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := node.Server{}

	grpcServer := grpc.NewServer()

	node.RegisterNodeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
