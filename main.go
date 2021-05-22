package main

import (
	"fmt"
	"github.com/simp7/nonogram/core"
	"github.com/simp7/nonogram/file/formatter"
	"github.com/simp7/nonogram/file/local"
	"github.com/simp7/nonogram/unit/standard"
	"github.com/simp7/nonograminGo/client"
)

func main() {

	fs, err := local.System()
	if err != nil {
		fmt.Println(err)
		return
	}
	coreData := core.New(fs, standard.Prototype(), formatter.Json(), formatter.Json())

	rd := client.Controller(coreData)
	rd.Start()

}
