package main

import (
	"github.com/simp7/nonograminGo/nonogram/controller"
	"os"
)

var rd = controller.CLI()

func init() {

	switch len(os.Args) {
	case 1:
		return
	case 2:
		if os.Args[1] == "alpha" {
			rd = controller.Improved()
		}
	default:
		//util.CheckErr(util.TooManyArgs)
	}

}

func main() {

	rd.Start()

}
