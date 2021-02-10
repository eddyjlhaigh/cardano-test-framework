package impl

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
	node "github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
)

type NodeServiceGrpcImpl struct {
}

func NewNodeServiceGrpcImpl() *NodeServiceGrpcImpl {
	return &NodeServiceGrpcImpl{}
}

func (serviceImpl *NodeServiceGrpcImpl) GetNodeVersion(ctx context.Context, in *node.Request) (*service.Response, error) {
	out, err := exec.Command("cardano-cli", "--version").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}

func (serviceImpl *NodeServiceGrpcImpl) OfflineKeyGen(ctx context.Context, in *node.Request) (*service.Response, error) {
	os.Chdir("./tmp")
	out, err := exec.Command("cardano-cli", "node", "key-gen",
		"--cold-verification-key-file", "cold.vkey",
		"--cold-signing-key-file", "cold.skey",
		"--operational-certificate-issue-counter-file", "cold.counter").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}
