package file

type Storage interface {
	Save(interface{}) error
	Load(interface{}) error
}
