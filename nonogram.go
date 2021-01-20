package main

import (
	"github.com/simp7/nonograminGo/control"
	"github.com/simp7/nonograminGo/util"
	"os"
)

var rd control.Controller = control.NewCliController()

func init() {

	switch len(os.Args) {
	case 1:
		return
	case 2:
		if os.Args[1] == "alpha" {
			rd = control.NewImprovedController()
		}
	default:
		util.CheckErr(util.TooManyArgs)
	}

}

func main() {

	rd.Start()

}
