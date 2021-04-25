package cell

import "sync"

var shapes map[Type]Shape
var once sync.Once

type Shape interface {
	Right() rune
	Left() rune
}

type shape struct {
	data [2]rune
}

func initShapes() {

	shapes = make(map[Type]Shape)

	blank := "  "
	check := "><"
	cursor := "()"

	shapes[Empty] = newShape(blank)
	shapes[Fill] = newShape(blank)
	shapes[Check] = newShape(check)
	shapes[Wrong] = newShape(check)
	shapes[Cursor] = newShape(cursor)
	shapes[CursorFilled] = newShape(cursor)
	shapes[CursorChecked] = newShape(cursor)
	shapes[CursorWrong] = newShape(cursor)

}

func shapeOf(t Type) Shape {
	once.Do(func() {
		initShapes()
	})
	return shapes[t]
}

func newShape(str string) Shape {
	s := new(shape)
	data := []rune(str)
	s.data[0] = data[0]
	s.data[1] = data[1]
	return s
}

func (s *shape) Right() rune {
	return s.data[0]
}

func (s *shape) Left() rune {
	return s.data[1]
}
