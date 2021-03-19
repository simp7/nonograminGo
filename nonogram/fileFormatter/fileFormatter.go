package fileFormatter

import (
	"bytes"
	"encoding/json"
	"github.com/simp7/nonograminGo/nonogram"
)

type fileFormatter struct {
	buffer bytes.Buffer
	*json.Encoder
	*json.Decoder
}

func New() nonogram.FileFormatter {
	f := new(fileFormatter)
	f.Encoder = json.NewEncoder(&f.buffer)
	f.Decoder = json.NewDecoder(&f.buffer)
	return f
}

func (f *fileFormatter) GetRaw(from []byte) {
	f.buffer.Write(from)
}

func (f *fileFormatter) Content() []byte {
	return f.buffer.Bytes()
}
