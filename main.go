package main

import (
	"fmt"
	"github.com/simp7/nonograminGo/client/controller/cli"
	"github.com/simp7/nonograminGo/file/formatter"
	"github.com/simp7/nonograminGo/file/localstorage"
	"github.com/simp7/nonograminGo/nonogram/standard"
)

func main() {

	fs, err := localstorage.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	rd := cli.Controller(fs, formatter.Json(), standard.Map())
	rd.Start()

}
