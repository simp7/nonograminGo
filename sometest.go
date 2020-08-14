package main

import (
	"nonograminGo/model"
	"fmt"
)

func main() {
	doTest("3/2/5/2")
	doTest("3/3/7/5/2")
	doTest("4/4/13/7/14/11")
}

func doTest(data string) {
	test := model.NewNonomap(data)
	for _, v := range test.ShowBitMap() {
		fmt.Println(v)
	}
	for _, v := range test.ShowProblemHorizontal() {
		fmt.Println(v)
	}
	for _, v := range test.ShowProblemVertical() {
		fmt.Println(v)
	}
}
