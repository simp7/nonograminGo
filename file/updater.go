package file

//Updater is an interface that browse files from somewhere to another.
type Updater interface {
	Update() //Update updates files.
}
