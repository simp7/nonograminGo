package framework

type Pos struct {
	X, Y int
}

func New(x, y int) Pos {
	return Pos{x, y}
}

func (p Pos) Move(deltaX, deltaY int) Pos {
	return New(p.X+deltaX, p.Y+deltaY)
}

func Nil() Pos {
	return New(-1, -1)
}
