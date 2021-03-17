package object

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/position"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
)

type Timer interface {
	nonogram.Object
}

type timer struct {
	Text
	gadget.Stopwatch
}

func NewTimer(p position.Pos, parent nonogram.Object) Timer {
	t := new(timer)
	t.Stopwatch = stopwatch.Standard
	t.Text = NewText(p, parent, "0:00")
	return t
}

func (t *timer) Add(nonogram.Object) {
}
