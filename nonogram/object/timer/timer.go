package timer

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
	"github.com/simp7/nonograminGo/nonogram/object/text"
	"github.com/simp7/nonograminGo/nonogram/position"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
)

type timer struct {
	object.Text
	gadget.Stopwatch
}

func New(p position.Pos, parent nonogram.Object) object.Timer {
	t := new(timer)
	t.Stopwatch = stopwatch.Standard
	t.Text = text.New(p, parent, "0:00")
	return t
}

func (t *timer) Add(nonogram.Object) {
}
