package object

import "github.com/simp7/nonograminGo/util"

type text struct {
	Object
	content string
}

func NewText(p util.Pos, content string) Object {
	t := new(text)
	t.Object = newObject(p)
	t.content = content
	return t
}

func (t *text) Content() <-chan string {
	c := make(chan string, 1)
	c <- t.content
	return c
}
