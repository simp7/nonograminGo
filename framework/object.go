package framework

type Object interface {
	GetPos() Pos
	Move(Pos)
	Add(Object)
	Parent() Object
	Child(idx int) Object
}
