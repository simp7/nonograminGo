package control

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/control/window"
)

type improvedController struct {
	*asset.Setting
	windows   window.Stack
	eventChan chan termbox.Event
	endChan   chan struct{}
	event     termbox.Event
	view      map[View]window.Window
}

func NewImprovedController() Controller {
	c := new(improvedController)
	c.Setting = asset.GetSetting()
	c.eventChan = make(chan termbox.Event)
	c.endChan = make(chan struct{})
	return c
}

func (c *improvedController) Start() {
	c.openWindow(MainMenu)
	go func() {
		select {
		case c.eventChan <- termbox.PollEvent():
		case <-c.endChan:
			c.terminate()
			return
		}
	}()
}

func (c *improvedController) openWindow(v View) {
	c.windows.Push(c.view[v])
}

func (c *improvedController) closeWindow() {
	c.windows.Pop()
	if c.windows.Size() == 0 {
		close(c.endChan)
	}
}

func (c *improvedController) showWindow() {

}

func (c *improvedController) terminate() {
	close(c.eventChan)
	termbox.Close()
}
