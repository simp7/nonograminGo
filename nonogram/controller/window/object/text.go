package object

import (
	"github.com/simp7/nonograminGo/util"
)

type Text interface {
	Object
	CopyText() Text
}

type text struct {
	pos     util.Pos
	content string
	chars   []Char
	parent  Object
}

func NewText(p util.Pos, parent Object, content string) Text {
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

func (t *text) GetPos() util.Pos {
	return t.pos
}

func (t *text) Move(pos util.Pos) {
	t.pos = pos
}

func (t *text) Copy() Object {
	result := new(text)
	result.pos = t.pos
	result.content = t.content
	copy(result.chars, t.chars)
	return result
}

func (t *text) Add(o Object) {
	t.chars = append(t.chars, o)
}

func (t *text) Parent() Object {
	return t.parent
}

func (t *text) Child(idx int) Object {
	return t.chars[idx]
}

func (t *text) CopyText() Text {
	copied := new(text)
	copied.content = t.content
	return copied
}

func (t *text) do(f func(object Object)) {
	for _, v := range t.chars {
		f(v)
	}
}
