package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/impl"
	node "github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"

	"google.golang.org/grpc"
)

type server struct{}

func (s server) FetchResponse(in *node.Request, srv node.NodeService_FetchResponseServer) error {

	log.Printf("fetch response for id : %d", in.Body)

	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			//time sleep to simulate server process time
			time.Sleep(time.Duration(count) * time.Second)
			resp := node.Response{Body: fmt.Sprintf("Request #%d For Id:%d", count, in.Body)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc := grpc.NewServer()
	s := impl.NewNodeServiceGrpcImpl()
	node.RegisterNodeServiceServer(grpc, s)

	log.Println("start server")
	// and start...
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
