package file

//Loader is an interface that loads data to the program from somewhere.
type Loader interface {
	Load(interface{}) error //Load loads data from Loader to argument. argument should be address of wanted object.
}
