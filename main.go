package main

import (
	"./control"
	"./view"
)

func main() {

	rd := control.NewKeyReader()

	go rd.Control()
	view.ShowMenu()
	rd.ControlMenu()

}
