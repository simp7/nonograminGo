package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/util"
)

type Board interface {
	Object
}

type board struct {
	pos   util.Pos
	w, h  int
	cells [][]Cell
}

func (b *board) GetPos() util.Pos {
	return b.pos
}

func (b board) Content() <-chan string {
	panic("implement me")
}

func (b board) GetAttribute() (foreground termbox.Attribute, background termbox.Attribute) {
	panic("shouldn't be called.")
}

func (b board) Move(pos util.Pos) {
	b.pos = pos
}

func (b board) Copy() Object {
	panic("shouldn't be called.")
}

func NewBoard(p util.Pos, w, h int) Board {

	b := new(board)
	b.w, b.h = w, h
	b.pos = p

	for i := 0; i < h; i++ {
		b.cells = append(b.cells, make([]Cell, w))
	}

	return b

}
