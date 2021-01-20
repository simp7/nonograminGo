package util

type Pos struct {
	X, Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func (p Pos) Move(deltaX, deltaY int) Pos {
	return NewPos(p.X+deltaX, p.Y+deltaY)
}

func NilPos() Pos {
	return NewPos(-1, -1)
}
