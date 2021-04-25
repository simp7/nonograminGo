package window

import (
	"github.com/simp7/nonograminGo/framework"
)

type Window interface {
	getObjects() []framework.Object
}

type window struct {
	objects []framework.Object
	framework.Object
}

func (w *window) getObjects() []framework.Object {
	return w.objects
}

func (w *window) ShowObjects() {
}
