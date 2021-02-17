package stakepool

// StakePool represents the cardano-cli stake-pool sub-commands
type StakePool struct{}

// DeregistrationCertificate - Create a stake pool deregistration certificate
func (s *StakePool) DeregistrationCertificate() {}

// ID - Build pool id from the offline key
func (s *StakePool) ID() {}

// MetadataHash - Print the hash of pool metadata
func (s *StakePool) MetadataHash() {}

// RegistrationCertificate - Create a stake pool registration certificate
func (s *StakePool) RegistrationCertificate() {}
