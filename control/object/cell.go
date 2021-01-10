package object

import (
	"github.com/simp7/nonograminGo/util"
	"sync"
)

type CellState int8

const (
	Empty CellState = iota
	Fill
	Check
	Wrong
	Cursor
	CursorFilled
	CursorChecked
	CursorWrong
)

var cells map[CellState]Object
var once sync.Once

type cell struct {
	Object
}

func GetCell(state CellState) Object {
	once.Do(func() {
		cells = make(map[CellState]Object)
	})
	return cells[state]
}

func NewCell(p util.Pos) Object {
	c := new(cell)
	c.Object = newObject(p)
	return c
}
