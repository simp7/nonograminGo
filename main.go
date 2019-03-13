package main

import (
	"./control"
	"./view"
)

func main() {

	rd := control.NewKeyReader()

	rd.Control()
	view.ShowMenu()

}
