package client

//Pos is an struct that represents position.
type Pos struct {
	X, Y int
}

//Move returns another Pos that moves from p by arguments.
func (p Pos) Move(deltaX, deltaY int) Pos {
	return Pos{p.X + deltaX, p.Y + deltaY}
}
