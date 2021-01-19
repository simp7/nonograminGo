package window

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/control/window/object"
)

type Drawer interface {
	Draw(object object.Object)
}

type drawer struct {
}

func NewDrawer() Drawer {
	d := new(drawer)
	return d
}

func (d *drawer) Draw(object object.Object) {
	pos := object.GetPos()
	fg, bg := object.GetAttribute()
	go func() {
		for {
			content, ok := <-object.Content()
			if !ok {
				return
			}
			for i, c := range content {
				termbox.SetCell(pos.X+i, pos.Y, c, fg, bg)
			}
		}
	}()

}
