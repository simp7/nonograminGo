package controller

import (
	"github.com/simp7/nonograminGo/nonogram/controller/window/object"
)

type Drawer interface {
	Draw(object.Object)
	Empty()
}

type drawer struct {
}

func NewDrawer() Drawer {
	d := new(drawer)
	return d
}

func (d *drawer) Draw(target object.Object) {

	go func() {
		for {

		}
	}()

}

func (d *drawer) Empty() {
	//util.CheckErr(termbox.Clear(d.Color.Empty, d.Color.Empty))
}
