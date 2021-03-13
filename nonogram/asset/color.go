package asset

import (
	"github.com/nsf/termbox-go"
)

type Color struct {
	Empty   termbox.Attribute
	Filled  termbox.Attribute
	Checked termbox.Attribute
	Wrong   termbox.Attribute
	Char    termbox.Attribute
}
