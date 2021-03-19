package signal

import "github.com/simp7/nonograminGo/nonogram"

const (
	Empty nonogram.Signal = iota
	Fill
	Check
	Wrong
	Cursor
	CursorFilled
	CursorChecked
	CursorWrong
)
