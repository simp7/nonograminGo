package main

import (
	"fmt"
	"github.com/simp7/nonograminGo/client/controller/cli"
	config2 "github.com/simp7/nonograminGo/config"
)

func main() {
	setting, err := config2.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := cli.Controller(setting)
	rd.Start()
}
