package main

import (
	"io"
	"log"

	node "github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// dial server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := node.NewNodeServiceClient(conn)
	in := &node.Request{Body: "1"}
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Body)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}
