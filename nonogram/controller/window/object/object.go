package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Object interface {
	GetPos() util.Pos
	Move(util.Pos)
	Copy() Object
	Add(Object)
	Parent() Object
	Child(idx int) Object
	do(func(Object))
}

type object struct {
	pos    util.Pos
	parent Object
}

func newObject(p util.Pos, parent Object) Object {
	obj := new(object)
	obj.pos = p
	obj.parent = parent
	return obj
}

func (obj *object) GetPos() util.Pos {
	return obj.pos
}

func (obj *object) Move(p util.Pos) {
	obj.pos = p
}

func (obj *object) Copy() Object {
	return newObject(obj.pos, obj.parent)
}

func (obj *object) Add(object Object) {
	panic("Implement this")
}

func (obj *object) Parent() Object {
	return obj.parent
}

func (obj *object) Child(idx int) Object {
	panic("Implement this")
}

func (obj *object) do(f func(object Object)) {
	panic("Implement this")
}
