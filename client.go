package main

import (
	"log"

	"github.com/eddyjlhaigh/cardano-test-framework/node"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	n := node.NewNodeServiceClient(conn)

	message := node.GetVersionMessage{
		Body: "Hello from the client!",
	}

	response, err := n.PrintDir(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)

}
