package window

import (
	"github.com/simp7/nonograminGo/nonogram/controller/window/object"
)

type Imp interface {
	Draw(object object.Object)
}

type imp struct {
}

func NewImp() Imp {
	d := new(imp)
	return d
}

func (d *imp) Draw(object object.Object) {
	go func() {
		for {
		}
	}()

}
