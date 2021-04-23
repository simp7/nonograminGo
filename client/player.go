package client

import "github.com/simp7/nonograminGo/nonogram"

type Player interface {
	SetCell(Signal)
	SetCursor(Signal)
	RealPos() (x, y int)
	GetMapSignal() Signal
	SetMapSignal(Signal)
	Toggle(Signal)
	Move(Direction)
	FinishCreating() nonogram.Map
}
