package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
	"google.golang.org/grpc"
)

func main() {
	// Command line flags
	var portflag int

	flag.IntVar(&portflag, "p", 50005, "Specify port number for server.")
	flag.Parse()

	targetURL := fmt.Sprintf("%s%d", ":", portflag)

	// dial server
	conn, err := grpc.Dial(targetURL, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewNodeServiceClient(conn)

	req := service.Request{Body: "Test"}
	if res, err := client.GetNodeVersion(context.Background(), &req); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	if res, err := client.OfflineKeyGen(context.Background(), &req); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
