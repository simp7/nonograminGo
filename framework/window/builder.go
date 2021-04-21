package window

import (
	"github.com/simp7/nonograminGo/framework"
	board2 "github.com/simp7/nonograminGo/framework/object/board"
	text2 "github.com/simp7/nonograminGo/framework/object/text"
	textField2 "github.com/simp7/nonograminGo/framework/object/textField"
	timer2 "github.com/simp7/nonograminGo/framework/object/timer"
	setting2 "github.com/simp7/nonograminGo/framework/setting"
)

type buildAction func()

type Builder interface {
	initWindow() Builder
	AddText(framework.Pos, string) Builder
	AddTexts(framework.Pos, ...string) Builder
	AddTextField(framework.Pos) Builder
	AddTimer(framework.Pos) Builder
	AddBoard(pos framework.Pos, width, height int) Builder
	GetWindow() Window
}

type builder struct {
	window  *window
	actions []buildAction
	*framework.Setting
}

func NewBuilder() Builder {
	b := new(builder)
	b.initWindow()
	return b
}

func (b *builder) initWindow() Builder {
	b.window = new(window)
	b.Setting, _ = setting2.Get()
	return b
}

func (b *builder) add(object framework.Object) {
	b.window.objects = append(b.window.objects, object)
}

func (b *builder) AddText(pos framework.Pos, content string) Builder {
	b.appendAction(func() {
		b.add(text2.New(pos, b.window, content))
	})
	return b
}

func (b *builder) AddTexts(pos framework.Pos, contents ...string) Builder {
	b.appendAction(func() {
		for i, content := range contents {
			b.AddText(framework.New(pos.X, pos.Y+i), content)
		}
	})
	return b
}

func (b *builder) AddTextField(pos framework.Pos) Builder {
	b.appendAction(func() {
		b.add(textField2.New(pos, b.window))
	})
	return b
}

func (b *builder) AddTimer(pos framework.Pos) Builder {
	b.appendAction(func() {
		b.add(timer2.New(pos, b.window))
	})
	return b
}

func (b *builder) AddBoard(pos framework.Pos, width, height int) Builder {
	b.appendAction(func() {
		b.add(board2.New(pos, b.window, width, height))
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
