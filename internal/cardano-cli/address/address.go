package address

import (
	"log"
	"os/exec"
)

// Address represents the cardano-cli address sub-commands
type Address struct{}

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
	outFile string,
) {
	output, err := exec.Command(
		"cardano-cli", "address", "build",
		"--payment-verification-key-file", paymentVerificationKeyFile,
		"--stake-verification-key-file", stakeVerificationKeyFile,
		"--out-file", outFile,
		"--mainnet",
	).Output()
	if err != nil {
		log.Printf("error: %v\n%v", err, string(output))
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
