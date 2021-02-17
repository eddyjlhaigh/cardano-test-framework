package genesis

// Genesis represents the cardano-cli genesis sub-commands
type Genesis struct{}

// Create - Create a Shelley genesis file from a genesis template and genesis/delegation/spending keys
func (g *Genesis) Create() {}

// CreateStaked - Create a staked Shelley genesis file from a genesis template and genesis/delegation/spending keys
func (g *Genesis) CreateStaked() {}

// GetVerKey - Derive the verification key from a signing key
func (g *Genesis) GetVerKey() {}

// Hash - Compute the hash of a genesis file
func (g *Genesis) Hash() {}

// InitialAddr - Get the address for an initial UTxO based on the verification key
func (g *Genesis) InitialAddr() {}

// InitialTxIn - Get the TxIn for an initial UTxO based on the verification key
func (g *Genesis) InitialTxIn() {}

// KeyGenDelegate - Create a Shelley genesis delegate key pair
func (g *Genesis) KeyGenDelegate() {}

// KeyGenGenesis - Create a Shelley genesis key pair
func (g *Genesis) KeyGenGenesis() {}

// KeyGenUtxo - Create a Shelley genesis UTxO key pair
func (g *Genesis) KeyGenUtxo() {}

// KeyHash - Print the identifier (hash) of a public key
func (g *Genesis) KeyHash() {}
