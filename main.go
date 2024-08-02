package main

import (
	"dfs-system/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":8000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
