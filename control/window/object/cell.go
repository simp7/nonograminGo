package object

import (
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

var cells map[CellState]Cell
var once sync.Once

type Cell interface {
	Object
	CopyCell() Cell
}

type cell struct {
	Object
	state CellState
}

func (c *cell) CopyCell() Cell {
	copied := new(cell)
	copied.state = c.state
	return copied
}

func GetCell(state CellState) Cell {
	c := new(cell)
	c.state = state
	return c
}
