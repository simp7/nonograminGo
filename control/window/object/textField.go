package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/util"
)

type TextField interface {
	Object
	Activate()
	Deactivate()
}

type textField struct {
	Object
	isActive bool
	content  chan string
}

func NewTextField(p util.Pos, fg, bg termbox.Attribute) TextField {
	t := new(textField)
	t.Object = newObject(p, fg, bg)
	t.isActive = false
	return t
}

func (t *textField) Content() <-chan string {
	return t.content
}

func (t *textField) Activate() {
	t.isActive = true
}

func (t *textField) Deactivate() {
	t.isActive = false
}
