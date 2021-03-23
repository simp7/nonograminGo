package nonogram

type Drawer interface {
	Draw(Object)
	Empty()
}
