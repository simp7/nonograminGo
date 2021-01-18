package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/util"
)

type Text interface {
	Object
	CopyText() Text
}

type text struct {
	Object
	content string
}

func NewText(p util.Pos, fg, bg termbox.Attribute, content string) Text {
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

func (t *text) CopyText() Text {
	copied := new(text)
	copied.Object = t.Copy()
	copied.content = t.content
	return copied
}
