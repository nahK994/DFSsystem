package p2p

import "net"

// RPC holds any arbitary data that is being sent over the
// each transport between two nodes in the networks.
type RPC struct {
	Payload []byte
	From    net.Addr
}
