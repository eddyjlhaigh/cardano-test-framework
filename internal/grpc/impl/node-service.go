package impl

import (
	"context"
	"log"
	"os/exec"

	cli "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli"
	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
	node "github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
)

type NodeServiceGrpcImpl struct {
}

func NewNodeServiceGrpcImpl() *NodeServiceGrpcImpl {
	return &NodeServiceGrpcImpl{}
}

func getNodeConfigs() {
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-config.json").Run()
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-byron-genesis.json").Run()
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-shelley-genesis.json").Run()
	exec.Command("wget", "https://hydra.iohk.io/build/5416636/download/1/mainnet-topology.json").Run()
}

// RunNode launches an instance of cardano-node with the mainnet configuration files
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
	}

	return &service.Response{Body: "Node Running"}, nil
}

// GetNodeVersion returns the current cardano-cli version
func (serviceImpl *NodeServiceGrpcImpl) GetNodeVersion(ctx context.Context, in *node.Request) (*service.Response, error) {
	var cli cli.CLI
	nodeVersion := cli.Version()
	return &service.Response{Body: string(nodeVersion)}, nil
}
