package object

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/position"
	"github.com/simp7/nonograminGo/util"
	"strconv"
)

type TextField interface {
	nonogram.Object
	Activate()
	Deactivate()
}

type textField struct {
	nonogram.Object
	isActive bool
	content  chan string
}

func NewTextField(p position.Pos, parent nonogram.Object) TextField {
	t := new(textField)
	t.Object = New(p, parent)
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
	util.CheckErr(err)
	return result
}
