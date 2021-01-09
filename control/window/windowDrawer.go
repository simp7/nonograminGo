package window

import (
	"github.com/simp7/nonograminGo/asset"
)

type Drawer interface {
}

type drawer struct {
	*asset.Setting
	window Window
}

func NewDrawer(setting *asset.Setting, w Window) Drawer {
	d := new(drawer)
	d.Setting = setting
	d.window = w
	return d
}
