package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Object interface {
	GetPos() util.Pos
}

type object struct {
	pos util.Pos
}

func NewObject(p util.Pos) Object {
	obj := new(object)
	obj.pos = p
	return obj
}

func (obj *object) Draw() {

}

func (obj *object) GetPos() util.Pos {
	return obj.pos
}
