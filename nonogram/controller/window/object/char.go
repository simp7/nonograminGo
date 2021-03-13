package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Char interface {
	Object
}

type char struct {
	Object
	ch rune
}

func NewChar(pos util.Pos, parent Object, ch rune) Char {
	c := new(char)
	c.Object = newObject(pos, parent)
	c.ch = ch
	return c
}
