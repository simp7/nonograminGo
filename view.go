package main

import (
	"fmt"
	"io/ioutil"
)

// This file controls whole view of the game.

func ShowMenu() {
	fmt.Println("1. Start")
	fmt.Println("2. Create")
	fmt.Println("3. Hall of Honor")
	fmt.Println("4. Exit")
}

func ShowMapList() {
	files, err := ioutil.ReadDir("./maps")
	CheckErr(err)
	for n, file := range files {
		fmt.Printf("%d. %s\n", n, file.Name())
	}
}

func ShowInGame(nm nonomap) {
}

func ShowResult() {
}

func ShowRecord() {
}
