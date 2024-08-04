package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decoder(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (d GOBDecoder) Decoder(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct{}

func (d DefaultDecoder) Decoder(r io.Reader, rpc *RPC) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	rpc.Payload = buf[:n]
	return nil
}
