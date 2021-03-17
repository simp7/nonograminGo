package window

import (
	"github.com/simp7/nonograminGo/nonogram/setting"
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
	st := setting.Get()

	m.windows = make(map[View]Window)
	m.windows[MainMenu] = NewBuilder().AddTexts(st.DefaultPos, st.MainMenu()).GetWindow()
	m.windows[Select] = NewBuilder().AddTexts(st.DefaultPos, st.GetSelectHeader()).GetWindow()
}
