package window

import (
	setting2 "github.com/simp7/nonograminGo/framework/setting"
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

	st, _ := setting2.Get()

	m.windows = make(map[View]Window)
	m.windows[MainMenu] = NewBuilder().AddTexts(st.DefaultPos, st.MainMenu()...).GetWindow()
	m.windows[Select] = NewBuilder().AddTexts(st.DefaultPos, st.GetSelectHeader()...).GetWindow()
	m.windows[Help] = NewBuilder().AddTexts(st.DefaultPos, st.GetHelp()...).GetWindow()
	m.windows[Credit] = NewBuilder().AddTexts(st.DefaultPos, st.GetCredit()...).GetWindow()

}
