package controller

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/asset"
	"github.com/simp7/nonograminGo/nonogram/controller/window"
	"github.com/simp7/nonograminGo/util"
)

type improvedController struct {
	*asset.Setting
	windows     window.Stack
	eventChan   chan termbox.Event
	endChan     chan struct{}
	event       termbox.Event
	winConstant window.Constant
}

func Improved() nonogram.Controller {
	c := new(improvedController)
	c.windows = window.NewStack()
	c.Setting = asset.GetSetting()
	c.eventChan = make(chan termbox.Event)
	c.endChan = make(chan struct{})
	c.winConstant = window.GetManager()
	return c
}

func (c *improvedController) Start() {
	util.CheckErr(termbox.Init())
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

func (c *improvedController) openWindow(v window.View) {
	c.windows.Push(c.winConstant.Get(v))
	c.showWindow()
}

func (c *improvedController) closeWindow() {
	c.windows.Pop()
	if c.windows.Size() == 0 {
		close(c.endChan)
	} else {
		c.showWindow()
	}
}

func (c *improvedController) showWindow() {
	for {
		c.windows.Top()
	}
}

func (c *improvedController) terminate() {
	close(c.eventChan)
	termbox.Close()
}
