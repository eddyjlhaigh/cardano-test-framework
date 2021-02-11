package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/impl"
	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// Command line flags
	var portflag int

	flag.IntVar(&portflag, "p", 50005, "Specify port number for server.")
	flag.Parse()

	netListener := getNetListener(portflag)

	// Set-up and register gRPC Server & Services
	grpcServer := grpc.NewServer()

	nodeServiceImpl := impl.NewNodeServiceGrpcImpl()
	service.RegisterNodeServiceServer(grpcServer, nodeServiceImpl)

	// Run server and register
	log.Println("start server")
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func getNetListener(port int) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
