package char

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
	"github.com/simp7/nonograminGo/nonogram/position"
)

type char struct {
	nonogram.Object
	ch rune
}

func New(pos position.Pos, parent nonogram.Object, ch rune) object.Char {
	c := new(char)
	c.Object = object.New(pos, parent)
	c.ch = ch
	return c
}
