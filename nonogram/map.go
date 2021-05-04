package nonogram

//Map is an interface that represent map of nonogram.
//ShouldFilled returns whether filling selected cell is right.
//Problem returns Problem of current map.
//GetHeight returns height of map.
//GetWidth returns width of map.
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
