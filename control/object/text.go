package object

import "github.com/simp7/nonograminGo/util"

type Text interface {
	Object
}

type text struct {
	Object
	content string
}

func NewText(p util.Pos, content string) Text {
	t := new(text)
	t.Object = NewObject(p)
	t.content = content
	return t
}

func (t *text) String() string {
	return t.content
}

type texts struct {
	Object
	contents []Text
}

func NewTexts(pos util.Pos, contents []string) Text {

	t := new(texts)
	t.Object = NewObject(pos)
	t.contents = make([]Text, len(contents))
	for i := range contents {
		t.contents[i] = NewText(util.NewPos(pos.X, pos.Y+i), contents[i])
	}

	return t

}
