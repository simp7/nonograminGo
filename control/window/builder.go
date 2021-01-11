package window

import (
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/control"
	"github.com/simp7/nonograminGo/control/object"
	"github.com/simp7/nonograminGo/util"
)

type buildAction func()

type Builder interface {
	InitWindow() Builder
	AddText(util.Pos, string) Builder
	AddTexts(util.Pos, []string) Builder
	AddTextField(util.Pos) Builder
	AddTimer(util.Pos) Builder
	AddCell(pos util.Pos, width, height int) Builder
	GetWindow() Window
}

type builder struct {
	window  *window
	actions []buildAction
	*asset.Setting
}

func NewBuilder() Builder {
	b := new(builder)
	return b
}

func (b *builder) InitWindow() Builder {
	b.window = new(window)
	b.window.Drawer = control.NewDrawer()
	b.Setting = asset.GetSetting()
	return b
}

func (b *builder) AddText(pos util.Pos, content string) Builder {
	b.appendAction(func() {
		b.window.objects = append(b.window.objects, object.NewText(pos, b.Char, b.Empty, content))
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
		b.window.objects = append(b.window.objects, object.NewTextField(pos, b.Char, b.Empty))
	})
	return b
}

func (b *builder) AddTimer(pos util.Pos) Builder {
	b.appendAction(func() {
		b.window.objects = append(b.window.objects, object.NewTimer(pos, b.Char, b.Empty))
	})
	return b
}

func (b *builder) AddCell(pos util.Pos, width, height int) Builder {
	b.appendAction(func() {

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
