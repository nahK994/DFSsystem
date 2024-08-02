package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct {
	conn net.Conn

	// if we dial a connection => outbound = true
	// if we accept and retrieve a connection => outbound = false
	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listerer      net.Listener
	peer          map[net.Addr]Peer
	mu            sync.RWMutex
}

func NewTCPTransport(listenAdds string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAdds,
	}
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listerer, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listerer.Accept()
		if err != nil {
			fmt.Printf("TCP accept error %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("new incoming connection %+v\n", peer)
}
