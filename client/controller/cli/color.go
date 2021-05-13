package cli

import (
	"github.com/nsf/termbox-go"
)

//Color represents colors for display of application.
type Color struct {
	Empty   termbox.Attribute
	Filled  termbox.Attribute
	Checked termbox.Attribute
	Wrong   termbox.Attribute
	Char    termbox.Attribute
}
