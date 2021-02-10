package main

import (
	"context"
	"fmt"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
	"google.golang.org/grpc"
)

func main() {
	// dial server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewNodeServiceClient(conn)

	req := service.Request{Body: "Test"}
	if res, err := client.PrintDir(context.Background(), &req); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
