package cli

import (
	"fmt"
	"log"
	"os/exec"
)

// CLI encapsulates all cardano-cli commands and sub-commands
type CLI struct {
	Address      Address
	Byron        Byron
	Genesis      Genesis
	Governance   Governance
	Key          Key
	Node         Node
	StakeAddress StakeAddress
	StakePool    StakePool
	Transaction  Transaction
	TextView     TextView
	Query        Query
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
