package file

//Loader is an interface that loads data to the program from somewhere.
//Load loads data from somewhere to argument. argument should be address of wanted object.
type Loader interface {
	Load(interface{}) error
}
