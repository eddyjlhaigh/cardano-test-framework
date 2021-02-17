package node

// Node represents the cardano-cli node sub-commands
type Node struct{}

// IssueOpCert - Issue a node operational certificate
func (n *Node) IssueOpCert() {}

// KeyGen - Create a key pair for a node operator's offline key and a new certificate issue counter
func (n *Node) KeyGen() {}

// KeyGenKES - Create a key pair for a node KES operational key
func (n *Node) KeyGenKES() {}

// KeyGenVRF - Create a key pair for a node VRF operational key
func (n *Node) KeyGenVRF() {}

// KeyHashVRF - Print hash of a node's operational VRF key
func (n *Node) KeyHashVRF() {}

// NewCounter - Create a new certificate issue counter
func (n *Node) NewCounter() {}
