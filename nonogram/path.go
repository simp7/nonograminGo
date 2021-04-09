package nonogram

type Path interface {
	String() string
	Append(...string) Path
}
