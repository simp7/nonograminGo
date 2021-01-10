package control

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/control/object"
	"github.com/simp7/nonograminGo/util"
)

type Drawer interface {
	Draw(object.Object)
	Empty()
}

type drawer struct {
	asset.Color
}

func NewDrawer(setting *asset.Setting) Drawer {
	d := new(drawer)
	d.Color = setting.Color
	return d
}

func (d *drawer) Draw(target object.Object) {

	p := target.GetPos()

	go func() {
		for {

			text, ok := <-target.Content()
			if !ok {
				return
			}

			for i, character := range []rune(text) {
				termbox.SetCell(p.X+i, p.Y, character, d.Color.Char, d.Color.Empty)
			}

		}
	}()

}

func (d *drawer) Empty() {
	util.CheckErr(termbox.Clear(d.Color.Empty, d.Color.Empty))
}
