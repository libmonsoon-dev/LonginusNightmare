package codec

import "io"

type Codec interface {
	NewDecoder(r io.Reader) Decoder
	NewEncoder(w io.Writer) Encoder
	GetContentType() string
	Name() string
}

type Decoder interface {
	Decode(e interface{}) error
}

type Encoder interface {
	Encode(v interface{}) error
}
