package p2p

// Message holds any arbitary data that is being sent over the
// each transport between two nodes in the networks.
type Message struct {
	Payload []byte
}
