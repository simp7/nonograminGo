package main

import (
	"./control"
	"./view"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	CheckErr(err)
	defer termbox.Close()

	rd := control.NewKeyStroker()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	view.ShowMenu()
	rd.ControlMenu()
}
