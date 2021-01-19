package window

import (
	"github.com/simp7/nonograminGo/control/window/object"
)

type Window interface {
	getObjects() []object.Object
}

type window struct {
	texts      []object.Text
	textFields []object.TextField
	timer      object.Timer
	board      object.Board
	objects    []object.Object
}

func (w *window) getObjects() []object.Object {
	return w.objects
}

func (w *window) ShowObjects() {
	objects := w.getObjects()
	objects[1].GetAttribute()
}
