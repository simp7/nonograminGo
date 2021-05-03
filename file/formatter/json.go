package formatter

import (
	"bytes"
	"encoding/json"
)

type jsonFormatter struct {
	buffer bytes.Buffer
	*json.Encoder
	*json.Decoder
}

// Json returns file.Formatter that includes encoder and decoder of json.
func Json() *jsonFormatter {
	f := new(jsonFormatter)
	f.Encoder = json.NewEncoder(&f.buffer)
	f.Decoder = json.NewDecoder(&f.buffer)
	return f
}

func (f *jsonFormatter) GetRaw(from []byte) error {
	_, err := f.buffer.Write(from)
	return err
}

func (f *jsonFormatter) Content() []byte {
	return f.buffer.Bytes()
}
