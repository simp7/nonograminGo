package asset

import (
	"github.com/nsf/termbox-go"
)

type Color interface {
	Empty() termbox.Attribute
	Filled() termbox.Attribute
	Checked() termbox.Attribute
	Wrong() termbox.Attribute
	Text() termbox.Attribute
}

type color struct {
	EmptyColor   termbox.Attribute
	FilledColor  termbox.Attribute
	CheckedColor termbox.Attribute
	WrongColor   termbox.Attribute
	TextColor    termbox.Attribute
}

func (c *color) Empty() termbox.Attribute {
	return c.EmptyColor
}

func (c *color) Filled() termbox.Attribute {
	return c.FilledColor
}

func (c *color) Checked() termbox.Attribute {
	return c.CheckedColor
}

func (c *color) Wrong() termbox.Attribute {
	return c.WrongColor
}

func (c *color) Text() termbox.Attribute {
	return c.TextColor
}
