package nonogram

type FileFormatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	GetRaw(from []byte)
	Content() []byte
}
