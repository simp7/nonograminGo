package window

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type Imp interface {
	Draw(object nonogram.Object)
}

type imp struct {
}

func NewImp() Imp {
	d := new(imp)
	return d
}

func (d *imp) Draw(object nonogram.Object) {
	go func() {
		for {
		}
	}()

}
