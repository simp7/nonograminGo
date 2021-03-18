package object

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/position"
)

type object struct {
	pos    position.Pos
	parent nonogram.Object
}

func New(p position.Pos, parent nonogram.Object) nonogram.Object {
	obj := new(object)
	obj.pos = p
	obj.parent = parent
	return obj
}

func (obj *object) GetPos() position.Pos {
	return obj.pos
}

func (obj *object) Move(p position.Pos) {
	obj.pos = p
}
func (obj *object) Add(object nonogram.Object) {
	panic("Implement this")
}

func (obj *object) Parent() nonogram.Object {
	return obj.parent
}

func (obj *object) Child(idx int) nonogram.Object {
	panic("Implement this")
}
