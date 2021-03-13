package window

import (
	"github.com/simp7/nonograminGo/nonogram/asset"
	"github.com/simp7/nonograminGo/nonogram/controller/window/object"
	"github.com/simp7/nonograminGo/util"
)

type buildAction func()

type Builder interface {
	initWindow() Builder
	AddText(util.Pos, string) Builder
	AddTexts(util.Pos, []string) Builder
	AddTextField(util.Pos) Builder
	AddTimer(util.Pos) Builder
	AddBoard(pos util.Pos, width, height int) Builder
	GetWindow() Window
}

type builder struct {
	window  *window
	actions []buildAction
	*asset.Setting
}

func NewBuilder() Builder {
	b := new(builder)
	b.initWindow()
	return b
}

func (b *builder) initWindow() Builder {
	b.window = new(window)
	b.Setting = asset.GetSetting()
	return b
}

func (b *builder) add(object object.Object) {
	b.window.objects = append(b.window.objects, object)
}

func (b *builder) AddText(pos util.Pos, content string) Builder {
	b.appendAction(func() {
		b.add(object.NewText(pos, b.window, content))
	})
	return b
}

func (b *builder) AddTexts(pos util.Pos, contents []string) Builder {
	b.appendAction(func() {
		for i, content := range contents {
			b.AddText(util.NewPos(pos.X, pos.Y+i), content)
		}
	})
	return b
}

func (b *builder) AddTextField(pos util.Pos) Builder {
	b.appendAction(func() {
		b.add(object.NewTextField(pos, b.window))
	})
	return b
}

func (b *builder) AddTimer(pos util.Pos) Builder {
	b.appendAction(func() {
		b.add(object.NewTimer(pos, b.window))
	})
	return b
}

func (b *builder) AddBoard(pos util.Pos, width, height int) Builder {
	b.appendAction(func() {
		b.add(object.NewBoard(pos, b.window, width, height))
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
