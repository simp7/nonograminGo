package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	CheckErr(err)
	defer termbox.Close()

	rd := NewKeyStroker()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	ShowMenu()
	rd.ControlMenu()
}
