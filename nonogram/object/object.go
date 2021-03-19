package object

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type object struct {
	pos    nonogram.Pos
	parent nonogram.Object
}

func New(p nonogram.Pos, parent nonogram.Object) nonogram.Object {
	obj := new(object)
	obj.pos = p
	obj.parent = parent
	return obj
}

func (obj *object) GetPos() nonogram.Pos {
	return obj.pos
}

func (obj *object) Move(p nonogram.Pos) {
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
