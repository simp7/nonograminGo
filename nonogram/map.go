package nonogram

//Map is an interface that represent map of nonogram.
type Map interface {
	ShouldFilled(x, y int) bool //ShouldFilled returns whether filling selected cell is right.
	CreateProblem() Problem     //CreateProblem returns Problem of current map.
	GetHeight() int             //GetHeight returns height of map.
	GetWidth() int              //GetWidth returns width of map.
	FilledTotal() int
	CheckValidity() error
	HeightLimit() int
	WidthLimit() int
	CopyWithBitmap([][]bool) Map
	Formatter() Formatter
}
