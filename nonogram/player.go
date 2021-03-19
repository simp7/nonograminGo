package nonogram

type Signal uint8
type Direction uint8

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
