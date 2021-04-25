package formatter

import (
	"bytes"
	"encoding/json"
)

type fileFormatter struct {
	buffer bytes.Buffer
	*json.Encoder
	*json.Decoder
}

func Json() *fileFormatter {
	f := new(fileFormatter)
	f.Encoder = json.NewEncoder(&f.buffer)
	f.Decoder = json.NewDecoder(&f.buffer)
	return f
}

func (f *fileFormatter) GetRaw(from []byte) error {
	_, err := f.buffer.Write(from)
	return err
}

func (f *fileFormatter) Content() []byte {
	return f.buffer.Bytes()
}
