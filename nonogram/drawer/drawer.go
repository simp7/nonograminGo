package drawer

import "github.com/simp7/nonograminGo/nonogram"

type drawer struct {
}

func New() nonogram.Drawer {
	d := new(drawer)
	return d
}

func (d *drawer) Draw(object nonogram.Object) {

	go func() {
		for {

		}
	}()

}

func (d *drawer) Empty() {
	//util.CheckErr(termbox.Clear(d.Color.Empty, d.Color.Empty))
}
