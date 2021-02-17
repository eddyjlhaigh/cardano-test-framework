package stakeaddress

// StakeAddress represents the cardano-cli stake-address sub-commands
type StakeAddress struct{}

// Build - Build a stake address
func (s *StakeAddress) Build() {}

// DelegationCertificate - Create a stake address delegation certificate
func (s *StakeAddress) DelegationCertificate() {}

// DeregistrationCertificate - Create a stake address deregistration certificate
func (s *StakeAddress) DeregistrationCertificate() {}

// KeyHash - Print the hash of a stake address key
func (s *StakeAddress) KeyHash() {}

// KeyGen - Create a stake address key pair
func (s *StakeAddress) KeyGen() {}

// RegistrationCertificate - Create a stake address registration certificate
func (s *StakeAddress) RegistrationCertificate() {}
