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

type TCPTransportOptions struct {
	ListenAddress string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	listerer net.Listener
	peer     map[net.Addr]Peer
	mu       sync.RWMutex
	rpcch    chan RPC
	TCPTransportOptions
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
		rpcch:               make(chan RPC),
	}
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close implements the Peer interface.
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// Consume implements the transport interface, which will return read-only channel
// for reading the incoming messages received from another peer in the network.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listerer, err = net.Listen("tcp", t.ListenAddress)
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

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	// Read loop
	rpc := RPC{}
	// buf := make([]byte, 200)
	for {
		if err := t.Decoder.Decoder(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
		// fmt.Printf("message: %+v\n", rpc)
		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}
}
