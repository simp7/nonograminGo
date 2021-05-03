package file

//Loader is an interface that loads data to the program from somewhere.
type Loader interface {
	Load(interface{}) error
}
