package util

import (
	"bytes"
	"encoding/json"
)

type FileFormatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	GetRaw(from string)
	Content() []byte
}

type fileFormatter struct {
	buffer bytes.Buffer
	*json.Encoder
	*json.Decoder
}

func NewFileFormatter() FileFormatter {
	f := new(fileFormatter)
	f.Encoder = json.NewEncoder(&f.buffer)
	f.Decoder = json.NewDecoder(&f.buffer)
	return f
}

func (f *fileFormatter) GetRaw(from string) {
	f.buffer.WriteString(from)
}

func (f *fileFormatter) Content() []byte {
	return f.buffer.Bytes()
}
