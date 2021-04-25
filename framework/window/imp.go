package window

import (
	"github.com/simp7/nonograminGo/framework"
)

type Imp interface {
	Draw(object framework.Object)
}

type imp struct {
}

func NewImp() Imp {
	d := new(imp)
	return d
}

func (d *imp) Draw(object framework.Object) {
	go func() {
		for {
		}
	}()

}
