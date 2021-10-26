package main

import (
	"github.com/gdamore/tcell/termbox"
)

//Color is an struct that includes color of cells.
type Color struct {
	Empty   termbox.Attribute
	Char    termbox.Attribute
	Filled  termbox.Attribute
	Checked termbox.Attribute
	Wrong   termbox.Attribute
}

// Light returns an instance of Color
func Light() Color {
	return Color{termbox.ColorWhite, termbox.ColorBlack, termbox.ColorBlack, termbox.ColorCyan, termbox.ColorRed}
}

// Dark returns an instance of Color
func Dark() Color {
	return Color{termbox.ColorBlack, termbox.ColorGreen, termbox.ColorWhite, termbox.ColorCyan, termbox.ColorRed}
}
