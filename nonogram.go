package main

import (
	"github.com/simp7/nonograminGo/control"
	"github.com/simp7/nonograminGo/util"
	"os"
)

var rd control.Controller = control.NewCliController()

func init() {

	if len(os.Args) > 2 {
		util.CheckErr(util.TooManyArgs)
	} else if os.Args[1] == "alpha" {
		rd = control.NewImprovedController()
	}

}

func main() {

	rd.Start()

}
