package cell

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
)

type State interface {
	nonogram.Object
}

type state struct {
	nonogram.Object
	isCursored bool
}

func newState(pos nonogram.Pos, parent nonogram.Object) State {
	s := new(state)
	s.Object = object.New(pos, parent)
	s.isCursored = false
	return s
}
