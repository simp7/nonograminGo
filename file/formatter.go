package file

//Formatter is an interface that encodes or decodes data into specified format.
type Formatter interface {
	Encode(interface{}) error //Encode is function that saves objects from argument to Formatter with specific format.
	Decode(interface{}) error //Decode is function that loads objects from Formatter to argument with specific format. argument should be address of wanted object.
	GetRaw(from []byte) error //GetRaw is function that loads raw data to Formatter
	Content() []byte          //Content is function that returns raw data in Formatter.
}
