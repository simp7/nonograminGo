package nonogram

/*
	This file deals with algorithms of whole game of nonogram.
	User's control or display should be separated from this file.
*/

type Map interface {
	ShouldFilled(x, y int) bool
	CreateProblem() Problem
	GetHeight() int
	GetWidth() int
	FilledTotal() int
	CheckValidity() error
	HeightLimit() int
	WidthLimit() int
	Formatter() Formatter
}
