package asset

type Figure struct {
	NameMax   int
	WidthMax  int
	HeightMax int
	DefaultX  int
	DefaultY  int
}

func defaultFigure() Figure {

	var f Figure

	f.NameMax = 30
	f.WidthMax = 30
	f.HeightMax = 30
	f.DefaultX = 5
	f.DefaultY = 5

	return f

}
