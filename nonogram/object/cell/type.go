package cell

type Type int8

const (
	Empty Type = iota
	Fill
	Check
	Wrong
	Cursor
	CursorFilled
	CursorChecked
	CursorWrong
)
