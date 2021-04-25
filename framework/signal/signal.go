package signal

import (
	"github.com/simp7/nonograminGo/framework"
)

const (
	Empty framework.Signal = iota
	Fill
	Check
	Wrong
	Cursor
	CursorFilled
	CursorChecked
	CursorWrong
)
