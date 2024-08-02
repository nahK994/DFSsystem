package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decoder(io.Reader, any) error
}

type GOBDecoder struct{}

func (d *GOBDecoder) Decoder(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}
