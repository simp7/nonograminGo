package cli

type Pos struct {
	X, Y int
}

func (p Pos) Move(deltaX, deltaY int) Pos {
	return Pos{p.X + deltaX, p.Y + deltaY}
}
