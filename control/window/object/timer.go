package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/util"
)

type Timer interface {
	Object
}

type timer struct {
	Object
	util.Timer
}

func NewTimer(p util.Pos, fg, bg termbox.Attribute) Timer {
	t := new(timer)
	t.Timer = util.StartTimer()
	t.Object = newObject(p, fg, bg)
	return t
}

func (t *timer) Content() <-chan string {

	c := make(chan string, 1)
	go t.Do(func(current string) {
		c <- current
	})

	return c

}
