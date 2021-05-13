package file

//Saver is an interface that saves data from the program to somewhere.
type Saver interface {
	Save(interface{}) error //Save saves data from argument to Saver.
}
