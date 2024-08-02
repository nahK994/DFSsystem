package p2p

import (
	"fmt"
	"net"
	"sync"
)

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
	fmt.Printf("new incoming connection %+v\n", conn)
}
