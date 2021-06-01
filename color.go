package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonogram/setting"
	"image/color"
)

//Color is an struct that includes color of cells.
type Color struct {
	Empty   termbox.Attribute
	Char    termbox.Attribute
	Filled  termbox.Attribute
	Checked termbox.Attribute
	Wrong   termbox.Attribute
}

func adapt(target color.RGBA) termbox.Attribute {
	return termbox.RGBToAttribute(target.R, target.G, target.B)
}

func AdaptColor(target setting.Color) Color {
	return Color{adapt(target.Empty), adapt(target.Char), adapt(target.Filled), adapt(target.Checked), adapt(target.Wrong)}
}
