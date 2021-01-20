package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Timer interface {
	Object
}

type timer struct {
	Text
	util.Timer
}

func NewTimer(p util.Pos, parent Object) Timer {
	t := new(timer)
	t.Timer = util.StartTimer()
	t.Text = NewText(p, parent, "0:00")
	return t
}

func (t *timer) Content() <-chan string {

	c := make(chan string, 1)
	go t.Do(func(current string) {
		c <- current
	})

	return c

}
