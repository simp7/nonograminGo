package cell

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
)

type State interface {
	framework.Object
}

type state struct {
	framework.Object
	isCursored bool
}

func newState(pos framework.Pos, parent framework.Object) State {
	s := new(state)
	s.Object = object2.New(pos, parent)
	s.isCursored = false
	return s
}
