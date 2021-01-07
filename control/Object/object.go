package Object

import "github.com/simp7/nonograminGo/util"

type Object interface {
	Draw()
	GetPos() util.Pos
	GetRealPos() util.Pos
}

type object struct {
	pos      util.Pos
	children []Object
	parent   Object
}

func NewObject(parent Object) Object {
	obj := new(object)
	obj.parent = parent
	return obj
}

func (obj *object) Draw() {

}

func (obj *object) GetPos() util.Pos {
	return obj.pos
}

func (obj *object) GetRealPos() util.Pos {
	if obj.parent == nil {
		return obj.pos
	}
	return obj.pos.Add(obj.parent.GetRealPos())
}
