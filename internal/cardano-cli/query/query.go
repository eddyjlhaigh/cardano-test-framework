package query

import (
	"io/ioutil"
	"log"
	"os/exec"
)

// Query represents the cardano-cli query sub-commands
type Query struct{}

// LedgerState - Dump the current ledger state of the node (Ledger.NewEpochState -- advanced command)
func (q *Query) LedgerState() {}

// ProtocolParameters - Get the node's current protocol parameters
func (q *Query) ProtocolParameters() {}

// ProtocolState - Dump the current protocol state of the node (Ledger.ChainDepState -- advanced command)
func (q *Query) ProtocolState() {}

// StakeAddressInfo - Get the current delegations and reward accounts
func (q *Query) StakeAddressInfo() {}

// StakeDistribution - Get the node's current aggregated stake distribution
func (q *Query) StakeDistribution() {}

// Tip - Get the node's current tip (slot no, hash, block no)
func (q *Query) Tip() {}

// Utxo - Get the node's current UTxO with the option of filtering by address(es)
func (q *Query) Utxo(
	addressFile string,
) {
	bytes, err := ioutil.ReadFile(addressFile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command(
		"cardano-cli", "query", "utxo",
		"--address", string(bytes),
		"--mainnet",
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
