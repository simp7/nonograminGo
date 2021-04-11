package nonogram

/*
	This file deals with algorithms of whole game of nonogram.
	User's control or display should be separated from this file.
*/

type Map interface {
	ShouldFilled(x, y int) bool
	CreateProblemFormat() (hProblem, vProblem []string, hMax, vMax int)
	GetHeight() int
	GetWidth() int
	BitmapToStrings() []string
	ShowProblemHorizontal() []string
	ShowProblemVertical() []string
	FilledTotal() int
	CheckValidity()
	HeightLimit() int
	WidthLimit() int
	Builder() MapBuilder
}
