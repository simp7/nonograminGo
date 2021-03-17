package object

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/position"
)

type Char interface {
	nonogram.Object
}

type char struct {
	nonogram.Object
	ch rune
}

func NewChar(pos position.Pos, parent nonogram.Object, ch rune) Char {
	c := new(char)
	c.Object = New(pos, parent)
	c.ch = ch
	return c
}
