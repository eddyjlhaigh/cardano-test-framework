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

func getNodeConfigs() {
	os.Chdir("./tmp/node_configs")
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-config.json").Run()
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-byron-genesis.json").Run()
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-shelley-genesis.json").Run()
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-topology.json").Run()
}

func (serviceImpl *NodeServiceGrpcImpl) RunNode(ctx context.Context, in *node.Request) (*service.Response, error) {
	getNodeConfigs()
	err := exec.Command("cardano-node", "run",
		"--topology", "mainnet-topology.json",
		"--database-path", "db",
		"--socket-path", "db/node.socket",
		"--port", "3011",
		"--config", "mainnet-config.json").Start()

	if err != nil {
		log.Fatal(err)
		return &service.Response{Body: err.Error()}, nil
	} else {
		return &service.Response{Body: "Node Running"}, nil
	}
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
