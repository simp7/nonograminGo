package char

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
)

type char struct {
	nonogram.Object
	ch rune
}

func New(pos nonogram.Pos, parent nonogram.Object, ch rune) object.Char {
	c := new(char)
	c.Object = object.New(pos, parent)
	c.ch = ch
	return c
}
