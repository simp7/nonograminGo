package view

import (
	"../util"
	"fmt"
	"io/ioutil"
)

/*
This part controls whole view of the game.
This should be performed with control together.
*/

// This function will enumerate elements in the list

func showByList(list ...string) {
	for n := 1; n < (len(list) + 1); n++ {
		fmt.Printf("%d. %s\n", n, list[n-1])
	}
}

// This function will show main menu when program starts.

func ShowMenu() {
	list := [4]string{"Start", "Load", "Create", "Exit"}
	showByList(list)
}

//This function will show list of maps in 'maps' directory which player can play.
//This function will be called when player select 'start' in main menu

func ShowMapList() {
	files, err := ioutil.ReadDir("./maps")
	CheckErr(err)
	for n, file := range files {
		fmt.Printf("%d. %s\n", n, file.Name())
	}
}

//This function shows display of gameplay.
//This function will be called when player select map in map list or select 'load' in main menu.

func ShowInGame(nm nonomap) {
}

//This function shows result of gameplay that player finished now.
//This function will be called when player ends game.

func ShowResult() {
	fmt.Printf("----------\n\n  RESULT\n\n----------")
	fmt.Println("Map name : ")
	fmt.Println("Map size : ")
	fmt.Println("Cleared time : ")
	fmt.Println("Press any key to Continue.")
}

//This function shows display of user-creating maps.
//This function will be called when player select 'create' in main menu.
func ShowInCreate(nm nonomap) {
}
