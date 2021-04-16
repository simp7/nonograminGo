package main

import (
	"errors"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram/controller"
	"os"
)

var (
	rd          = controller.CLI()
	tooManyArgs = errors.New("argument should be less than 2")
)

func init() {

	switch len(os.Args) {
	case 1:
		return
	case 2:
		if os.Args[1] == "alpha" {
			rd = controller.Improved()
		}
	default:
		errs.Check(tooManyArgs)
	}

}

func main() {
	rd.Start()
}
