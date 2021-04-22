package main

import (
	"fmt"
	config2 "github.com/simp7/nonograminGo/config"
	"github.com/simp7/nonograminGo/framework/controller/cli"
)

func main() {
	setting, err := config2.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := cli.CLI(setting)
	rd.Start()
}
