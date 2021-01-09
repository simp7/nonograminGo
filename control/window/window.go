package window

import "github.com/simp7/nonograminGo/control/object"

type Window interface {
}

type window struct {
	objects []object.Object
}

func (w *window) Refresh() {
}
