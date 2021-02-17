package key

// Key represents the cardano-cli key sub-commands
type Key struct{}

// ConvertByronGenesisVkey - Convert a Base64-encoded Byron genesis verification key to a Shelley genesis verification key
func (k *Key) ConvertByronGenesisVkey() {}

// ConvertByronKey - Convert a Byron payment, genesis or genesis delegate key (signing or verification) to a corresponding Shelley-format key.
func (k *Key) ConvertByronKey() {}

// ConvertCardanoAddressKey -  Convert a cardano-address extended signing key to a corresponding Shelley-format key.
func (k *Key) ConvertCardanoAddressKey() {}

// ConvertItnBip32Key - Convert an Incentivized Testnet (ITN) BIP32 (Ed25519Bip32) signing key to a corresponding Shelley stake signing key
func (k *Key) ConvertItnBip32Key() {}

// ConvertItnExtendedKey - Convert an Incentivized Testnet (ITN) extended (Ed25519Extended) signing key to a corresponding Shelley stake signing key
func (k *Key) ConvertItnExtendedKey() {}

// ConvertItnKey - Convert an Incentivized Testnet (ITN) non-extended (Ed25519) signing or verification key to a corresponding Shelley stake key
func (k *Key) ConvertItnKey() {}

// NonExtendedKey - Get a non-extended verification key from an extended verification key. This supports all extended key types.
func (k *Key) NonExtendedKey() {}

// VerificationKey - Get a verification key from a signing key. This supports all key types.
func (k *Key) VerificationKey() {}
