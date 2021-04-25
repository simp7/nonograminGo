package timer

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
	text2 "github.com/simp7/nonograminGo/framework/object/text"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
)

type timer struct {
	object2.Text
	gadget.Stopwatch
}

func New(p framework.Pos, parent framework.Object) object2.Timer {
	t := new(timer)
	t.Stopwatch = stopwatch.Standard
	t.Text = text2.New(p, parent, "0:00")
	return t
}

func (t *timer) Add(framework.Object) {
}
