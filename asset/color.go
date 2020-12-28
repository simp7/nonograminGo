package asset

import (
	"github.com/nsf/termbox-go"
)

type Color struct {
	Empty   termbox.Attribute
	Filled  termbox.Attribute
	Checked termbox.Attribute
	Wrong   termbox.Attribute
	Text    termbox.Attribute
}

func defaultColor() Color {

	var c Color

	c.Empty = termbox.ColorWhite
	c.Filled = termbox.ColorBlack
	c.Checked = termbox.ColorBlue
	c.Wrong = termbox.ColorRed | termbox.AttrBold
	c.Text = termbox.ColorBlack

	return c

}
