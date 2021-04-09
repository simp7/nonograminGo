package controller

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/setting"
	"github.com/simp7/nonograminGo/nonogram/window"
)

type improved struct {
	*nonogram.Setting
	windows     window.Stack
	eventChan   chan termbox.Event
	endChan     chan struct{}
	event       termbox.Event
	winConstant window.Constant
}

func Improved() nonogram.Controller {
	c := new(improved)
	c.windows = window.NewStack()
	c.Setting = setting.Get()
	c.eventChan = make(chan termbox.Event)
	c.endChan = make(chan struct{})
	c.winConstant = window.GetManager()
	return c
}

func (c *improved) Start() {
	errs.Check(termbox.Init())
	go func() {
		select {
		case c.eventChan <- termbox.PollEvent():
		case <-c.endChan:
			c.terminate()
			return
		}
	}()
	c.openWindow(window.MainMenu)
}

func (c *improved) openWindow(v window.View) {
	c.windows.Push(c.winConstant.Get(v))
	c.showWindow()
}

func (c *improved) closeWindow() {
	c.windows.Pop()
	if c.windows.Size() == 0 {
		close(c.endChan)
	} else {
		c.showWindow()
	}
}

func (c *improved) showWindow() {
	for {
		c.windows.Top()
	}
}

func (c *improved) terminate() {
	close(c.eventChan)
	termbox.Close()
}
