package nonogram

type Map interface {
	ShouldFilled(x, y int) bool
	CreateProblem() Problem
	GetHeight() int
	GetWidth() int
	FilledTotal() int
	CheckValidity() error
	HeightLimit() int
	WidthLimit() int
	CopyWithBitmap([][]bool) Map
	Formatter() Formatter
}
