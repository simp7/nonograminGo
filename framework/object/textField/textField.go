package textField

import (
	"github.com/simp7/nonograminGo/framework"
	object2 "github.com/simp7/nonograminGo/framework/object"
	"strconv"
)

type textField struct {
	framework.Object
	isActive bool
	content  chan string
}

func New(p framework.Pos, parent framework.Object) object2.TextField {
	t := new(textField)
	t.Object = object2.New(p, parent)
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

func (t *textField) GetString() string {
	return <-t.content
}

func (t *textField) GetInt() (int, error) {
	return strconv.Atoi(<-t.content)
}
