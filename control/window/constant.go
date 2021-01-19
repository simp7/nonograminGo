package window

import (
	"github.com/simp7/nonograminGo/asset"
	"sync"
)

type View uint8

const (
	MainMenu View = iota
	Select
	Help
	Credit
)

var instance Constant
var once sync.Once

type Constant interface {
	Get(View) Window
}

type constant struct {
	windows map[View]Window
}

func GetManager() Constant {
	once.Do(func() {
		m := new(constant)
		m.initialize()
		instance = m
	})
	return instance
}

func (m *constant) Get(v View) Window {
	return m.windows[v]
}

func (m *constant) initialize() {
	setting := asset.GetSetting()

	m.windows = make(map[View]Window)
	m.windows[MainMenu] = NewBuilder().AddTexts(setting.DefaultPos, setting.MainMenu()).GetWindow()
	m.windows[Select] = NewBuilder().AddTexts(setting.DefaultPos, setting.GetSelectHeader()).GetWindow()
}
