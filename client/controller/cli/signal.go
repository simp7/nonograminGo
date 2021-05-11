package cli

//Signal represents state of nonogram cell.
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
