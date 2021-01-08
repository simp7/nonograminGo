package control

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/control/window"
)

type improvedController struct {
	*asset.Setting
	windows     window.Stack
	eventChan   chan termbox.Event
	endChan     chan struct{}
	currentView View
	event       termbox.Event
}

func NewImprovedController() Controller {
	c := new(improvedController)
	c.Setting = asset.GetSetting()
	return c
}

func (c *improvedController) Start() {
}
