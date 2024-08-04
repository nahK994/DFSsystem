package main

import (
	"dfs-system/p2p"
	"fmt"
	"log"
)

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddress: ":8000",
		HandshakeFunc: p2p.NOHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
