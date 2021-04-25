package text

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
	char2 "github.com/simp7/nonograminGo/framework/object/char"
)

type text struct {
	pos     framework.Pos
	content string
	chars   []object2.Char
	parent  framework.Object
}

func New(p framework.Pos, parent framework.Object, content string) object2.Text {
	t := new(text)
	t.pos = p
	t.content = content
	t.parent = parent
	return t
}

func (t *text) initContent() {

	t.chars = make([]object2.Char, 0)
	pos := t.pos

	for i, ch := range t.content {
		t.Add(char2.New(pos.Move(i, 0), t, ch))
	}

}

func (t *text) GetPos() framework.Pos {
	return t.pos
}

func (t *text) Move(pos framework.Pos) {
	t.pos = pos
}

func (t *text) Add(o framework.Object) {
	t.chars = append(t.chars, o)
}

func (t *text) Parent() framework.Object {
	return t.parent
}

func (t *text) Child(idx int) framework.Object {
	return t.chars[idx]
}

func (t *text) CopyText() object2.Text {
	copied := new(text)
	copied.content = t.content
	return copied
}

func (t *text) do(f func(object framework.Object)) {
	for _, v := range t.chars {
		f(v)
	}
}
