package window

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/object/board"
	"github.com/simp7/nonograminGo/nonogram/object/text"
	"github.com/simp7/nonograminGo/nonogram/object/textField"
	"github.com/simp7/nonograminGo/nonogram/object/timer"
	"github.com/simp7/nonograminGo/nonogram/position"
	"github.com/simp7/nonograminGo/nonogram/setting"
)

type buildAction func()

type Builder interface {
	initWindow() Builder
	AddText(position.Pos, string) Builder
	AddTexts(position.Pos, []string) Builder
	AddTextField(position.Pos) Builder
	AddTimer(position.Pos) Builder
	AddBoard(pos position.Pos, width, height int) Builder
	GetWindow() Window
}

type builder struct {
	window  *window
	actions []buildAction
	*setting.Setting
}

func NewBuilder() Builder {
	b := new(builder)
	b.initWindow()
	return b
}

func (b *builder) initWindow() Builder {
	b.window = new(window)
	b.Setting = setting.Get()
	return b
}

func (b *builder) add(object nonogram.Object) {
	b.window.objects = append(b.window.objects, object)
}

func (b *builder) AddText(pos position.Pos, content string) Builder {
	b.appendAction(func() {
		b.add(text.New(pos, b.window, content))
	})
	return b
}

func (b *builder) AddTexts(pos position.Pos, contents []string) Builder {
	b.appendAction(func() {
		for i, content := range contents {
			b.AddText(position.New(pos.X, pos.Y+i), content)
		}
	})
	return b
}

func (b *builder) AddTextField(pos position.Pos) Builder {
	b.appendAction(func() {
		b.add(textField.New(pos, b.window))
	})
	return b
}

func (b *builder) AddTimer(pos position.Pos) Builder {
	b.appendAction(func() {
		b.add(timer.New(pos, b.window))
	})
	return b
}

func (b *builder) AddBoard(pos position.Pos, width, height int) Builder {
	b.appendAction(func() {
		b.add(board.New(pos, b.window, width, height))
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
