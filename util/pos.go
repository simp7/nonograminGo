package util

type Pos struct {
	X, Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func (p Pos) Add(another Pos) Pos {
	return NewPos(p.X+another.X, p.Y+another.Y)
}

func NilPos() Pos {
	return NewPos(-1, -1)
}
