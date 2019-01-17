package main

import (
	"fmt"
	"io/ioutil"
)

func showMenu() {
	fmt.Println("1. Start")
	fmt.Println("2. Random")
	fmt.Println("3. Hall of Honor")
	fmt.Println("4. Exit")
}

func showMapList() {
	files, err := ioutil.ReadDir("./maps")
	CheckErr(err)
	for n, file := range files {
		fmt.Printf("%d. %s\n", n, file.Name())
	}
}

func showInGame() {
}

func showResult() {
}

func showRecord() {
}
