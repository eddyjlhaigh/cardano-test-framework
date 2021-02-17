package stakeaddress

import (
	"log"
	"os/exec"
)

// StakeAddress represents the cardano-cli stake-address sub-commands
type StakeAddress struct{}

// Build - Build a stake address
func (s *StakeAddress) Build(
	stakeVerificationKeyFile string,
	outFile string,
) {
	_, err := exec.Command(
		"cardano-cli", "stake-address", "build",
		"--stake-verification-key-file", stakeVerificationKeyFile,
		"--out-file", outFile,
		"--mainnet",
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// DelegationCertificate - Create a stake address delegation certificate
func (s *StakeAddress) DelegationCertificate() {}

// DeregistrationCertificate - Create a stake address deregistration certificate
func (s *StakeAddress) DeregistrationCertificate() {}

// KeyHash - Print the hash of a stake address key
func (s *StakeAddress) KeyHash() {}

// KeyGen - Create a stake address key pair
func (s *StakeAddress) KeyGen(
	verificationKeyFileName string,
	signingKeyFileName string,
) {
	_, err := exec.Command(
		"cardano-cli", "stake-address", "key-gen",
		"--verification-key-file", verificationKeyFileName,
		"--signing-key-file", signingKeyFileName,
	).Output()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// RegistrationCertificate - Create a stake address registration certificate
func (s *StakeAddress) RegistrationCertificate() {}
