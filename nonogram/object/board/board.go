package board

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
	"github.com/simp7/nonograminGo/nonogram/object/cell"
	"github.com/simp7/nonograminGo/nonogram/position"
)

type board struct {
	nonogram.Object
	w, h  int
	cells [][]object.Cell
}

func (b board) Copy() nonogram.Object {
	panic("shouldn't be called.")
}

func New(p position.Pos, parent nonogram.Object, w, h int) object.Board {

	b := new(board)
	b.w, b.h = w, h
	b.Object = object.New(p, parent)

	for i := 0; i < h; i++ {
		b.cells = append(b.cells, make([]object.Cell, w))
		b.initRow(i)
	}

	return b

}

func (b *board) Add(obj nonogram.Object) {

}

func (b *board) initRow(idx int) {

	for i := range b.cells[idx] {
		b.cells[i][idx] = cell.New(b.GetPos().Move(idx, i), b, cell.Empty)
	}

}
