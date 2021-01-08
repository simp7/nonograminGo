package window

import (
	"github.com/simp7/nonograminGo/control/object"
	"github.com/simp7/nonograminGo/util"
)

type buildAction func()

type Builder interface {
	InitWindow() Builder
	AddText(util.Pos, string) Builder
	AddTexts(util.Pos, []string) Builder
	GetWindow() Window
}

type builder struct {
	window  *window
	actions []buildAction
}

func NewBuilder() Builder {
	b := new(builder)
	return b
}

func (b *builder) InitWindow() Builder {
	b.window = new(window)
	return b
}

func (b *builder) AddText(pos util.Pos, content string) Builder {
	b.appendAction(func() {
		b.window.objects = append(b.window.objects, object.NewText(pos, content))
	})
	return b
}

func (b *builder) AddTexts(pos util.Pos, contents []string) Builder {
	b.appendAction(func() {
		b.window.objects = append(b.window.objects, object.NewTexts(pos, contents))
	})
	return b
}

func (b *builder) GetWindow() Window {
	b.build()
	return b.window
}

func (b *builder) appendAction(action buildAction) {
	b.actions = append(b.actions, action)
}

func (b *builder) build() {
	for _, action := range b.actions {
		action()
	}
}
