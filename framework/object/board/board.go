package board

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
	cell2 "github.com/simp7/nonograminGo/framework/object/cell"
)

type board struct {
	framework.Object
	w, h  int
	cells [][]object2.Cell
}

func (b board) Copy() framework.Object {
	panic("shouldn't be called.")
}

func New(p framework.Pos, parent framework.Object, w, h int) object2.Board {

	b := new(board)
	b.w, b.h = w, h
	b.Object = object2.New(p, parent)

	for i := 0; i < h; i++ {
		b.cells = append(b.cells, make([]object2.Cell, w))
		b.initRow(i)
	}

	return b

}

func (b *board) Add(obj framework.Object) {

}

func (b *board) initRow(idx int) {

	for i := range b.cells[idx] {
		b.cells[i][idx] = cell2.New(b.GetPos().Move(idx, i), b, cell2.Empty)
	}

}
