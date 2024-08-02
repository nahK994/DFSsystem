package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOptions{
		ListenAddress: ":8000",
		HandshakeFunc: NOHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	listenAddr := ":8000"
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.TCPTransportOptions.ListenAddress, listenAddr)
	assert.Nil(t, tr.ListenAndAccept())
	select {}
}
