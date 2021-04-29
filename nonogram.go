package main

import (
	"fmt"
	"github.com/simp7/nonograminGo/client/controller/cli"
	"github.com/simp7/nonograminGo/file/formatter"
	"github.com/simp7/nonograminGo/file/localStorage"
	"github.com/simp7/nonograminGo/nonogram/standard"
)

func main() {

	fs, err := localStorage.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	rd := cli.Controller(fs, formatter.Json(), standard.Map())
	rd.Start()

}
