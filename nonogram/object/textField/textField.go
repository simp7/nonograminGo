package textField

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object"
	"strconv"
)

type textField struct {
	nonogram.Object
	isActive bool
	content  chan string
}

func New(p nonogram.Pos, parent nonogram.Object) object.TextField {
	t := new(textField)
	t.Object = object.New(p, parent)
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

func (t *textField) GetInt() int {
	result, err := strconv.Atoi(<-t.content)
	errs.Check(err)
	return result
}
