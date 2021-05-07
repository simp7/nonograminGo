package file

//Saver is an interface that saves data from the program to somewhere.
//Save saves data from argument to Saver.
type Saver interface {
	Save(interface{}) error
}
