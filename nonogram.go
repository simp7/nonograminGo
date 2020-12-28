package main

import "github.com/simp7/nonograminGo/control"

func main() {

	var rd control.Controller
	rd = control.NewKeyReader()

	rd.Start()

}
