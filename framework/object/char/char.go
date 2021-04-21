package char

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
)

type char struct {
	framework.Object
	ch rune
}

func New(pos framework.Pos, parent framework.Object, ch rune) object2.Char {
	c := new(char)
	c.Object = object2.New(pos, parent)
	c.ch = ch
	return c
}
