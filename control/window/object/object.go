package object

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/util"
)

type Object interface {
	GetPos() util.Pos
	Content() <-chan string
	GetAttribute() (foreground termbox.Attribute, background termbox.Attribute)
	Move(util.Pos)
	Copy() Object
}

type object struct {
	pos util.Pos
	bg  termbox.Attribute
	fg  termbox.Attribute
}

func newObject(p util.Pos, fg, bg termbox.Attribute) Object {
	obj := new(object)
	obj.pos = p
	obj.fg = fg
	obj.bg = bg
	return obj
}

func (obj *object) GetPos() util.Pos {
	return obj.pos
}

func (obj *object) Content() <-chan string {
	return nil
}

func (obj *object) GetAttribute() (termbox.Attribute, termbox.Attribute) {
	return obj.fg, obj.bg
}

func (obj *object) Move(p util.Pos) {
	obj.pos = p
}

func (obj *object) Copy() Object {
	return newObject(obj.pos, obj.fg, obj.bg)
}
