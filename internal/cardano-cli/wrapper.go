package main

import (
	"fmt"
	"log"
	"os/exec"
)

// Address represents the cardano-cli address sub-commands
type Address struct{}

// Byron represents the cardano-cli Byron sub-commands
type Byron struct{}

// Genesis represents the cardano-cli genesis sub-commands
type Genesis struct{}

// Governance represents the cardano-cli governance sub-commands
type Governance struct{}

// Key represents the cardano-cli key sub-commands
type Key struct{}

// Node represents the cardano-cli node sub-commands
type Node struct{}

// StakeAddress represents the cardano-cli stake-address sub-commands
type StakeAddress struct{}

// StakePool represents the cardano-cli stake-pool sub-commands
type StakePool struct{}

// Transaction represents the cardano-cli transaction sub-commands
type Transaction struct{}

// TextView represents the cardano-cli text-view sub-commands
type TextView struct{}

// Query represents the cardano-cli query sub-commands
type Query struct{}

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

/****************************************************************
*	ADDRESS FUNCTIONS		                        *
****************************************************************/

// KeyGenNormal creates a normal address key pair
func (a *Address) KeyGenNormal(
	verificationKeyFileName string,
	signingKeyFileName string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "key-gen", "--normal-key",
		"--verification-key-file", verificationKeyFileName,
		"--signing-key-file", signingKeyFileName,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// KeyGenExtended creates an extended address key pair
func (a *Address) KeyGenExtended(
	verificationKeyFileName string,
	signingKeyFileName string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "key-gen", "--extended-key",
		"--verification-key-file", verificationKeyFileName,
		"--signing-key-file", signingKeyFileName,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// KeyGenByron creates a Byron address key pair
func (a *Address) KeyGenByron(
	verificationKeyFileName string,
	signingKeyFileName string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "key-gen", "--byron-key",
		"--verification-key-file", verificationKeyFileName,
		"--signing-key-file", signingKeyFileName,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// KeyHash - Print the hash of an address key
func (a *Address) KeyHash(
	paymentVerificationKeyFile string,
	outFile string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "key-hash",
		"--payment-verification-key-file", paymentVerificationKeyFile,
		"--out-file", outFile,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// Build - Build a Shelley payment address, with optional delegation to a stake address
func (a *Address) Build(
	paymentVerificationKeyFile string,
	stakeVerificationKeyFile string,
	magicId string,
	outFile string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "build",
		"--payment-verification-key-file", paymentVerificationKeyFile,
		"--stake-verification-key-file", stakeVerificationKeyFile,
		"--mainnet", magicId,
		"--out-file", outFile,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// BuildScript - Build a Shelley script address
func (a *Address) BuildScript(
	scriptFile string,
	magicId string,
	outFile string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "build-script",
		"--script-file", scriptFile,
		"--mainnet", magicId,
		"--out-file", outFile,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// Info - Print information about an address
func (a *Address) Info(
	address string,
	outFile string,
) {
	_, err := exec.Command(
		"cardano-cli", "address", "info",
		"--address", address,
		"--out-file", outFile,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

/****************************************************************
*	BYRON FUNCTIONS						*
****************************************************************/

/****************************************************************
*	GENESIS FUNCTIONS					*
****************************************************************/

/****************************************************************
*	GOVERNANCE FUNCTIONS					*
****************************************************************/

/****************************************************************
*	KEY FUNCTIONS						*
****************************************************************/

/****************************************************************
*	NODE FUNCTIONS						*
****************************************************************/

/****************************************************************
*	STAKE ADDRESS FUNCTIONS					*
****************************************************************/

/****************************************************************
*	STAKE POOL FUNCTIONS					*
****************************************************************/

/****************************************************************
*	TRANSACTION FUNCTIONS					*
****************************************************************/

/****************************************************************
*	TEXT VIEW FUNCTIONS					*
****************************************************************/

/****************************************************************
*	QUERY FUNCTIONS						*
****************************************************************/

// Version returns the `cardano-cli --version` output as a string
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
