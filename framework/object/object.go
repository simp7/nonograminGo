package object

import (
	"github.com/simp7/nonograminGo/framework"
)

type object struct {
	pos    framework.Pos
	parent framework.Object
}

func New(p framework.Pos, parent framework.Object) framework.Object {
	obj := new(object)
	obj.pos = p
	obj.parent = parent
	return obj
}

func (obj *object) GetPos() framework.Pos {
	return obj.pos
}

func (obj *object) Move(p framework.Pos) {
	obj.pos = p
}
func (obj *object) Add(object framework.Object) {
	panic("Implement this")
}

func (obj *object) Parent() framework.Object {
	return obj.parent
}

func (obj *object) Child(idx int) framework.Object {
	panic("Implement this")
}
