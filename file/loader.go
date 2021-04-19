package file

type Loader interface {
	Load(interface{}) error
}
