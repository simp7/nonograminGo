package util

import (
	"bytes"
	"encoding/gob"
)

type FileFormatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
}

type fileFormatter struct {
	buffer bytes.Buffer
	*gob.Encoder
	*gob.Decoder
}

func NewFileFormatter() FileFormatter {
	f := new(fileFormatter)
	f.Encoder = gob.NewEncoder(&f.buffer)
	f.Decoder = gob.NewDecoder(&f.buffer)
	return f
}
