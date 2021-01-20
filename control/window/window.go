package window

import (
	"github.com/simp7/nonograminGo/control/window/object"
)

type Window interface {
	getObjects() []object.Object
}

type window struct {
	objects []object.Object
	object.Object
}

func (w *window) getObjects() []object.Object {
	return w.objects
}

func (w *window) ShowObjects() {
}
