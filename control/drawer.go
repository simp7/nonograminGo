package control

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/control/window/object"
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

	p := target.GetPos()
	fg, bg := target.GetAttribute()

	go func() {
		for {

			text, ok := <-target.Content()
			if !ok {
				return
			}

			for i, character := range []rune(text) {
				termbox.SetCell(p.X+i, p.Y, character, fg, bg)
			}

		}
	}()

}

func (d *drawer) Empty() {
	//util.CheckErr(termbox.Clear(d.Color.Empty, d.Color.Empty))
}
