package object

import (
	"github.com/simp7/nonograminGo/util"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
)

type Timer interface {
	Object
}

type timer struct {
	Text
	gadget.Stopwatch
}

func NewTimer(p util.Pos, parent Object) Timer {
	t := new(timer)
	t.Stopwatch = stopwatch.Standard
	t.Text = NewText(p, parent, "0:00")
	return t
}

func (t *timer) Content() <-chan string {

	c := make(chan string, 1)
	go t.Stopwatch.Add(func(current string) {
		c <- current
	})

	return c

}

func (t *timer) Add(obj Object) {
}
