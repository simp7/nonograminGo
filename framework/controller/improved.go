package controller

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/framework"
	setting2 "github.com/simp7/nonograminGo/framework/setting"
	window2 "github.com/simp7/nonograminGo/framework/window"
)

type improved struct {
	*framework.Setting
	windows     window2.Stack
	eventChan   chan termbox.Event
	endChan     chan struct{}
	event       termbox.Event
	winConstant window2.Constant
}

func Improved() framework.Controller {
	c := new(improved)
	c.windows = window2.NewStack()
	c.Setting, _ = setting2.Get()
	c.eventChan = make(chan termbox.Event)
	c.endChan = make(chan struct{})
	c.winConstant = window2.GetManager()
	return c
}

func (c *improved) Start() {
	termbox.Init()
	go func() {
		select {
		case c.eventChan <- termbox.PollEvent():
		case <-c.endChan:
			c.terminate()
			return
		}
	}()
	c.openWindow(window2.MainMenu)
}

func (c *improved) openWindow(v window2.View) {
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
