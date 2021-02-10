package main

import (
	"fmt"
	"log"
	"net"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/impl"
	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	netListener := getNetListener(50005)
	grpcServer := grpc.NewServer()

	nodeServiceImpl := impl.NewNodeServiceGrpcImpl()
	service.RegisterNodeServiceServer(grpcServer, nodeServiceImpl)

	log.Println("start server")
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
