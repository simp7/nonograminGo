package cli

type Signal uint8

const (
	Empty Signal = iota
	Fill
	Check
	Wrong
	Cursor
	CursorFilled
	CursorChecked
	CursorWrong
)
