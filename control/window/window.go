package window

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/control"
	"github.com/simp7/nonograminGo/control/object"
)

type Window interface {
	Refresh()
}

type window struct {
	objects []object.Object
	control.Drawer
}

func (w *window) Refresh() {
	for _, object := range w.objects {
		w.Drawer.Draw(object)
	}
	termbox.Flush()
}
