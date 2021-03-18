package text

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
	"github.com/simp7/nonograminGo/nonogram/object/char"
	"github.com/simp7/nonograminGo/nonogram/position"
)

type text struct {
	pos     position.Pos
	content string
	chars   []object.Char
	parent  nonogram.Object
}

func New(p position.Pos, parent nonogram.Object, content string) object.Text {
	t := new(text)
	t.pos = p
	t.content = content
	t.parent = parent
	return t
}

func (t *text) initContent() {

	t.chars = make([]object.Char, 0)
	pos := t.pos

	for i, ch := range t.content {
		t.Add(char.New(pos.Move(i, 0), t, ch))
	}

}

func (t *text) GetPos() position.Pos {
	return t.pos
}

func (t *text) Move(pos position.Pos) {
	t.pos = pos
}

func (t *text) Add(o nonogram.Object) {
	t.chars = append(t.chars, o)
}

func (t *text) Parent() nonogram.Object {
	return t.parent
}

func (t *text) Child(idx int) nonogram.Object {
	return t.chars[idx]
}

func (t *text) CopyText() object.Text {
	copied := new(text)
	copied.content = t.content
	return copied
}

func (t *text) do(f func(object nonogram.Object)) {
	for _, v := range t.chars {
		f(v)
	}
}
