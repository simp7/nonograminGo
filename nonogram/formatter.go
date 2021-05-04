package nonogram

//Formatter is just an duplication of file.Formatter.
type Formatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	GetRaw(from []byte) error
	Content() []byte
}
