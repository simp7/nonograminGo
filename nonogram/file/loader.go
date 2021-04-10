package file

type Loader interface {
	Load(interface{})
	LoadDefault(interface{})
}
