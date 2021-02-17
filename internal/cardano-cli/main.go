package main

import (
	"fmt"
	"log"
	"os/exec"

	address "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/address"
	byron "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/byron"
	genesis "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/genesis"
	governace "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/governance"
	key "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/key"
	node "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/node"
	query "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/query"
	stake_address "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/stake-address"
	stake_pool "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/stake-pool"
	text_view "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/text-view"
	transaction "github.com/eddyjlhaigh/cardano-test-framework/internal/cardano-cli/transaction"
)

// CLI encapsulates all cardano-cli commands and sub-commands
type CLI struct {
	Address      address.Address
	Byron        byron.Byron
	Genesis      genesis.Genesis
	Governance   governace.Governance
	Key          key.Key
	Node         node.Node
	StakeAddress stake_address.StakeAddress
	StakePool    stake_pool.StakePool
	Transaction  transaction.Transaction
	TextView     text_view.TextView
	Query        query.Query
}

// Version - Returns the current `cardano-cli --version` output
func (c *CLI) Version() string {
	out, err := exec.Command("cardano-cli", "--version").Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return string(out)
}

func main() {
	var cli CLI
	cli.Address.KeyGenNormal("skey", "vkey")
	fmt.Println(cli.Version())
}
