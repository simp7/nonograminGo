package file

//Formatter is an interface that encodes or decodes data into specified format.
//Encode is function that saves objects from argument to Formatter with specific format.
//Decode is function that loads objects from Formatter to argument with specific format. argument should be address of wanted object.
//GetRaw is function that loads raw data to Formatter
//Content is function that returns raw data in Formatter.
type Formatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	GetRaw(from []byte) error
	Content() []byte
}
