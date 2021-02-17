package transaction

// Transaction represents the cardano-cli transaction sub-commands
type Transaction struct{}

// Assemble - Assemble a tx body and witness(es) to form a transaction
func (t *Transaction) Assemble() {}

// BuildRaw - Build a transaction (low-level, inconvenient)
func (t *Transaction) BuildRaw() {}

// CalculateMinFee - Calculate the minimum fee for a transaction
func (t *Transaction) CalculateMinFee() {}

// PolicyID - Calculate the PolicyId from the monetary policy script
func (t *Transaction) PolicyID() {}

// Sign - Sign a transaction
func (t *Transaction) Sign() {}

// Submit -  Submit a transaction to the local node whose Unix domain socket is obtained from the CARDANO_NODE_SOCKET_PATH enviromnent variable.
func (t *Transaction) Submit() {}

// TxID - Print a transaction identifier
func (t *Transaction) TxID() {}

// Witness - Create a transaction witness
func (t *Transaction) Witness() {}
