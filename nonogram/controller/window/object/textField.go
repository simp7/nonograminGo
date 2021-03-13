package object

import (
	"github.com/simp7/nonograminGo/util"
	"strconv"
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

func NewTextField(p util.Pos, parent Object) TextField {
	t := new(textField)
	t.Object = newObject(p, parent)
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
