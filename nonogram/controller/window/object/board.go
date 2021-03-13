package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Board interface {
	Object
}

type board struct {
	Object
	w, h  int
	cells [][]Cell
}

func (b board) Copy() Object {
	panic("shouldn't be called.")
}

func NewBoard(p util.Pos, parent Object, w, h int) Board {

	b := new(board)
	b.w, b.h = w, h
	b.Object = newObject(p, parent)

	for i := 0; i < h; i++ {
		b.cells = append(b.cells, make([]Cell, w))
		b.initRow(i)
	}

	return b

}

func (b *board) Add(obj Object) {

}

func (b *board) initRow(idx int) {

	for i := range b.cells[idx] {
		b.cells[i][idx] = GetCell(Empty)
	}

}
