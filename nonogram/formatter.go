package nonogram

type Formatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	GetRaw(from []byte)
	Content() []byte
}
