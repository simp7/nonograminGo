package main

import (
	"./util/util"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

//This file deals with algoritms of whole game of nonogram.
//User's control or display should be seperated from this file.

type nonomap struct {
	width   int
	height  int
	mapdata []int
}

/*
	nonomap is devided into 3 parts and has arguments equal or more than 3, which is seperated by '/'.

	First two elements indicates width and height respectively.

	Rest elements indicates actual map which player has to solve.
	Each elements indicates map data of each line.
	is designated by bitmap, which 0 is blank and 1 is filled one.
	Because the size of int is 32bits, width of maps can't be more than 32 bits.

	When it comes to player's map, 2 is checked one where player thinks that cell is blank.

	The extention of file is nm(*.nm)
*/

func newNonomap(fileName string) *nonomap {

	var imported nonomap

	file, err := ioutil.ReadFile(fmt.Sprintf("%s.nm", fileName))
	CheckErr(err)
	elements := strings.Split(string(file), "/")
	//Extract all data from wanted file.

	imported.width, err = strconv.Atoi(elements[0])
	imported.height, err = strconv.Atoi(elements[1])
	CheckErr(err)
	//Extract map's size from file.

	for _, v := range elements[2:] {
		temp, err := strconv.Atoi(v)
		imported.mapdata = append(imported.mapdata, temp)
		CheckErr(err)
	}
	//Extract map's answer from file.

	if imported.height > 30 || imported.width > 30 {
		CheckErr(invalidMap)
	} //Check if height and width meets criteria of size.

	for _, v := range imported.mapdata {
		if float64(v) >= math.Pow(2, float64(imported.width)) {
			CheckErr(invalidMap)
		} //Check whether height matches mapdata.
	}
	if len(imported.mapdata) != imported.height {
		CheckErr(invalidMap)
	} //Check whether height matches mapdata.

	//Check validity of file.

	return &imported

}

func (nm *nonomap) compareMap(answer nonomap, int x, int y) {

	if nm.mapdata[x][y] == answer[x][y] {
		return
	} else {

	}
} //Would be called when user paint cell.

func OneGame(nm nonomap) {

	player = newNonomap("Player")
	timer := util.NewPlaytime()
	estimatedTime := timer.timeResult()
	fmt.Println(estimatedTime)

}

func checkMark() {
}
