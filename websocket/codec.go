package websocket

import (
	"io"

	"github.com/libmonsoon-dev/LonginusNightmare/codec"
)

type Codec interface {
	NewDecoder(r io.Reader) codec.Decoder
	NewEncoder(w io.Writer) codec.Encoder
	Name() string
}
