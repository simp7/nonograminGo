package cli

import (
	"github.com/nsf/termbox-go"
)

//Color is an struct that includes color of cells.
type Color struct {
	Empty   termbox.Attribute
	Filled  termbox.Attribute
	Checked termbox.Attribute
	Wrong   termbox.Attribute
	Char    termbox.Attribute
}
