package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/asset"
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

var cells map[CellState]Cell
var once sync.Once

type Cell interface {
	Object
	CopyCell() Cell
}

type cell struct {
	Text
}

func (c *cell) CopyCell() Cell {
	copied := new(cell)
	copied.Text = c.CopyText()
	return copied
}

func GetCell(p util.Pos, state CellState) Cell {
	once.Do(func() {
		cells = make(map[CellState]Cell)
		setting := asset.GetSetting()
		cells[Empty] = newCell(setting.Empty, setting.Empty, "  ")
		cells[Check] = newCell(setting.Checked, setting.Empty, "><")
		cells[Fill] = newCell(setting.Filled, setting.Filled, "  ")
		cells[Wrong] = newCell(setting.Wrong, setting.Empty, "><")
		cells[Cursor] = newCell(setting.Char, setting.Empty, "()")
		cells[CursorFilled] = newCell(setting.Empty, setting.Filled, "()")
		cells[CursorChecked] = newCell(setting.Checked, setting.Empty, "()")
		cells[CursorWrong] = newCell(setting.Wrong, setting.Empty, "()")
	})
	selected := cells[state]
	result := selected.CopyCell()
	result.Move(p)
	return result
}

func newCell(fg, bg termbox.Attribute, shape string) Cell {
	c := new(cell)
	c.Text = NewText(util.NilPos(), fg, bg, shape)
	return c
}
