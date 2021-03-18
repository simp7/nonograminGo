package nonogram

type Signal uint8
type Direction uint8

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

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Player interface {
	SetMap(Signal)
	SetCursor(Signal)
	RealPos() (x, y int)
	GetMapSignal() Signal
	SetMapSignal(Signal)
	Toggle(Signal)
	Move(Direction)
	FinishCreating() [][]bool
}
