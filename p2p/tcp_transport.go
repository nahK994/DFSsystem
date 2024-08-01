package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listerer      net.Listener
	peer          map[net.Addr]Peer
	mu            sync.RWMutex
}

func NewTCPTransport(listenAdds string) *Transport {
	return &TCPTransport{
		listenAddress: listenAdds,
	}
}
