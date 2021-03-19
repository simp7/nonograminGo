package window

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type Window interface {
	getObjects() []nonogram.Object
}

type window struct {
	objects []nonogram.Object
	nonogram.Object
}

func (w *window) getObjects() []nonogram.Object {
	return w.objects
}

func (w *window) ShowObjects() {
}
