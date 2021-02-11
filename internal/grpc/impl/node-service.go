package impl

import (
	"context"
	"log"
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
	} else {
		return &service.Response{Body: "Node Running"}, nil
	}
}

// GetNodeVersion returns the current cardano-cli version
func (serviceImpl *NodeServiceGrpcImpl) GetNodeVersion(ctx context.Context, in *node.Request) (*service.Response, error) {

	out, err := exec.Command("cardano-cli", "--version").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}

// GetPaymentAddress returns a stake address
// Requires: `stake.vkey`, `payment.vkey` file
// Returns: `payment.addr` file
func (serviceImpl *NodeServiceGrpcImpl) GetPaymentAddress(ctx context.Context, in *node.Request) (*service.Response, error) {

	out, err := exec.Command("cardano-cli", "address", "build",
		"--payment-verification-key-file", "payment.vkey",
		"--stake-verification-key-file", "stake.vkey",
		"--out-file", "payment.addr",
		"--mainnet").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}

// GetStakeAddress returns a stake address
// Requires: `stake.vkey` file
// Returns: `stake.addr` file
func (serviceImpl *NodeServiceGrpcImpl) GetStakeAddress(ctx context.Context, in *node.Request) (*service.Response, error) {

	out, err := exec.Command("cardano-cli", "stake-address", "build",
		"--stake-verification-key-file", "stake.vkey",
		"--out-file", "stake.addr",
		"--mainnet").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}

// GetAddressBalance returns the balance of an address
// Requires: A `payment.addr` file
// Returns: The balance of that file
func (serviceImpl *NodeServiceGrpcImpl) GetAddressBalance(ctx context.Context, in *node.Request) (*service.Response, error) {

	out, err := exec.Command("cardano-cli", "query", "utxo",
		"--address", "$(cat payment.addr)",
		"--mainnet").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}

// GeneratePaymentKeyPair creates a verification and signing key
// Requires: Nothing
// Returns: `payment.vkey`, `payment.skey`
func (serviceImpl *NodeServiceGrpcImpl) GeneratePaymentKeyPair(ctx context.Context, in *node.Request) (*service.Response, error) {
	out, err := exec.Command("cardano-cli", "address", "key-gen",
		"--verification-key-file", "payment.vkey",
		"--signing-key-file", "payment.skey").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}

// GenerateStakeKeyPair creates a verification and signing key
// Requires: Nothing
// Returns: `stake.vkey`, `stake.skey`
func (serviceImpl *NodeServiceGrpcImpl) GenerateStakeKeyPair(ctx context.Context, in *node.Request) (*service.Response, error) {
	out, err := exec.Command("cardano-cli", "stake-address", "key-gen",
		"--verification-key-file", "stake.vkey",
		"--signing-key-file", "stake.skey").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return &service.Response{Body: string(out)}, nil
}
