package main

import (
	"dfs-system/p2p"
	"log"
)

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddress: ":8000",
		HandshakeFunc: p2p.NOHandshakeFunc,
		Decoder:       &p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
