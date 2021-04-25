package cell

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
)

type cell struct {
	State
	Shape
}

func New(pos framework.Pos, parent framework.Object, t Type) object2.Cell {
	c := new(cell)
	c.State = newState(pos, parent)
	c.Shape = shapeOf(t)
	return c
}
