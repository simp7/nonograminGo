package object

import (
	"github.com/simp7/nonograminGo/nonogram"
	pos2 "github.com/simp7/nonograminGo/nonogram/position"
)

type Text interface {
	nonogram.Object
	CopyText() Text
}

type text struct {
	pos     pos2.Pos
	content string
	chars   []Char
	parent  nonogram.Object
}

func NewText(p pos2.Pos, parent nonogram.Object, content string) Text {
	t := new(text)
	t.pos = p
	t.content = content
	t.parent = parent
	return t
}

func (t *text) initContent() {

	t.chars = make([]Char, 0)
	pos := t.pos

	for i, ch := range t.content {
		t.Add(NewChar(pos.Move(i, 0), t, ch))
	}

}

func (t *text) GetPos() pos2.Pos {
	return t.pos
}

func (t *text) Move(pos pos2.Pos) {
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

func (t *text) CopyText() Text {
	copied := new(text)
	copied.content = t.content
	return copied
}

func (t *text) do(f func(object nonogram.Object)) {
	for _, v := range t.chars {
		f(v)
	}
}
