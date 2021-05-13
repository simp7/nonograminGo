package nonogram

//Map is an interface that represent map of nonogram.
type Map interface {
	ShouldFilled(x, y int) bool  //ShouldFilled returns whether filling selected cell is right.
	CreateProblem() Problem      //CreateProblem returns Problem of current map.
	GetHeight() int              //GetHeight returns height of map.
	GetWidth() int               //GetWidth returns width of map.
	FilledTotal() int            //FilledTotal returns amount of cells to fill.
	CheckValidity() error        //CheckValidity returns whether the map is valid.
	HeightLimit() int            //HeightLimit returns limit of map's height.
	WidthLimit() int             //WidthLimit returns limit of map's width.
	CopyWithBitmap([][]bool) Map //CopyWithBitmap returns map by argument.
	GetFormatter() Formatter     //GetFormatter returns Formatter of map.
}
