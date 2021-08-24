package codec

import (
	"encoding/json"
	"io"
)

const jsonContentType = "application/json"

type jsonCodec struct{}

func (jsonCodec) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}

func (jsonCodec) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (jsonCodec) GetContentType() string {
	return jsonContentType
}

func (jsonCodec) Name() string {
	return "json"
}

func NewJSON() Codec {
	return jsonCodec{}
}
