package nonogram

import (
	"github.com/simp7/nonograminGo/nonogram/position"
)

type Object interface {
	GetPos() position.Pos
	Move(position.Pos)
	Add(Object)
	Parent() Object
	Child(idx int) Object
}
