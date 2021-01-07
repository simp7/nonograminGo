package control

import (
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/control/window"
)

type improvedController struct {
	*asset.Setting
	windows window.Stack
}

func NewImprovedController() Controller {
	c := new(improvedController)
	c.Setting = asset.GetSetting()
	return c
}

func (c *improvedController) Start() {
}
