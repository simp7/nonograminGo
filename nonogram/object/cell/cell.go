package cell

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
	"github.com/simp7/nonograminGo/nonogram/position"
)

type cell struct {
	State
	Shape
}

func New(pos position.Pos, parent nonogram.Object, t Type) object.Cell {
	c := new(cell)
	c.State = newState(pos, parent)
	c.Shape = shapeOf(t)
	return c
}
