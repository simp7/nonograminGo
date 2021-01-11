package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/util"
)

type text struct {
	Object
	content string
}

func NewText(p util.Pos, fg, bg termbox.Attribute, content string) Object {
	t := new(text)
	t.Object = newObject(p, fg, bg)
	t.content = content
	return t
}

func (t *text) Content() <-chan string {
	c := make(chan string, 1)
	c <- t.content
	return c
}
