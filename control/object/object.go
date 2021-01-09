package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Object interface {
	GetPos() util.Pos
	Content() <-chan string
}

type object struct {
	pos util.Pos
}

func newObject(p util.Pos) Object {
	obj := new(object)
	obj.pos = p
	return obj
}

func (obj *object) GetPos() util.Pos {
	return obj.pos
}

func (obj *object) Content() <-chan string {
	return nil
}
