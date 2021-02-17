package byron

// Genesis represents the cardano-cli byron genesis sub-commands
type Genesis struct{}

// Governance represents the cardano-cli byron governance sub-commands
type Governance struct{}

// Key represents the cardano-cli byron key sub-commands
type Key struct{}

// Miscellaneous represents the cardano-cli byron miscellaneous sub-commands
type Miscellaneous struct{}

// Query represents the cardano-cli byron query sub-commands
type Query struct{}

// Transaction represents the cardano-cli byron transaction sub-commands
type Transaction struct{}

// Byron represents the cardano-cli Byron sub-commands
type Byron struct {
	Genesis       Genesis
	Governance    Governance
	Key           Key
	Miscellaneous Miscellaneous
	Query         Query
	Transaction   Transaction
}

// Genesis - Create Genesis
func (g *Genesis) Genesis() {
}

// PrintGenesisHash - Compute hash of a genesis file
func (g *Genesis) PrintGenesisHash() {
}

// CreateUpdateProposal - Create an update proposal
func (g *Governance) CreateUpdateProposal() {}

// CreateProposalVote - Create an update proposal vote
func (g *Governance) CreateProposalVote() {}

// SubmitUpdateProposal - Submit an update proposal
func (g *Governance) SubmitUpdateProposal() {}

// SubmitProposalVote - Submit a proposal vote
func (g *Governance) SubmitProposalVote() {}

// KeyGen - Generate a signing key
func (k *Key) KeyGen() {}

// MigrateDelegateKeyFrom - Migrate a delegate key from an older version.
func (k *Key) MigrateDelegateKeyFrom() {}

// SigningKeyAddress - Print address of a signing key.
func (k *Key) SigningKeyAddress() {}

// SigningKeyPublic - Pretty-print a signing key's verification key (not a secret)
func (k *Key) SigningKeyPublic() {}

// ToVerification - Extract a verification key in its base64 form
func (k *Key) ToVerification() {}

// ValidateCbor - Validate a CBOR blockchain object
func (m *Miscellaneous) ValidateCbor() {}

// PrettyPrintCbor - Pretty print a CBOR file
func (m *Miscellaneous) PrettyPrintCbor() {}

// GetTip - Get the tip of your local node's blockchain
func (q *Query) GetTip() {}

// SubmitTx - Submit a raw, signed transaction, in its on-wire representation.
func (t *Transaction) SubmitTx() {}

// IssueGenesisUtxoExpenditure - Write a file with a signed transaction, spending genesis UTxO.
func (t *Transaction) IssueGenesisUtxoExpenditure() {}

// IssueUtxoExpenditure - Write a file with a signed transaction, spending normal UTxO.
func (t *Transaction) IssueUtxoExpenditure() {}

// TxID - Print the txid of a raw, signed transaction.
func (t *Transaction) TxID() {}
