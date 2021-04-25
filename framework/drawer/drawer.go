package drawer

import (
	"github.com/simp7/nonograminGo/framework"
)

type drawer struct {
}

func New() framework.Drawer {
	d := new(drawer)
	return d
}

func (d *drawer) Draw(object framework.Object) {

	go func() {
		for {

		}
	}()

}

func (d *drawer) Empty() {
	//util.CheckErr(termbox.Clear(d.Color.Empty, d.Color.Empty))
}
