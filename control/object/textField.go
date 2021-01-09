package object

import "github.com/simp7/nonograminGo/util"

type TextField interface {
	Object
	Activate()
}

type textField struct {
	object
	isActive bool
	content  chan string
}

func NewTextField(p util.Pos) Object {
	t := new(textField)
	t.pos = p
	t.isActive = false
	return t
}

func (t *textField) Content() <-chan string {
	return t.content
}

func (t *textField) Activate() {
	t.isActive = true
}
