package file

type Path interface {
	String() string
	Append(...string) Path
}
