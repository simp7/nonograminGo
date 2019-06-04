package model

import (
	"../util"
	"math"
	"strconv"
	"strings"
)

/*
	This file deals with algoritms of whole game of nonogram.
	User's control or display should be seperated from this file.
*/

type Nonomap struct {
	width   int
	height  int
	mapdata []int
	bitmap  [][]bool
}

/*
	nonomap is devided into 3 parts and has arguments equal or more than 3, which is seperated by '/'.

	First two elements indicates width and height respectively.

	Rest elements indicates actual map which player has to solve.
	Each elements indicates map data of each line.
	They are designated by bitmap, which 0 is blank and 1 is filled one.

	Because the size of int is 32bits, width of maps can't be more than 32 mathmatically.

	But because of display's limit, width and height can't be more than 25

	When it comes to player's map, 2 is checked one where player thinks that cell is blank.

	The extention of file is nm(*.nm)
*/

func NewNonomap(data string) *Nonomap {

	var imported Nonomap
	var err error

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")
	//Extract all data from wanted file.

	imported.width, err = strconv.Atoi(elements[0])
	imported.height, err = strconv.Atoi(elements[1])
	util.CheckErr(err)
	//Extract map's size from file.

	for _, v := range elements[2:] {
		temp, err := strconv.Atoi(v)
		imported.mapdata = append(imported.mapdata, temp)
		util.CheckErr(err)
	}
	//Extract map's answer from file.

	if imported.height > 25 || imported.width > 25 || imported.height < 0 || imported.width < 0 {
		util.CheckErr(util.InvalidMap)
	} //Check if height and width meets criteria of size.

	for _, v := range imported.mapdata {
		if float64(v) >= math.Pow(2, float64(imported.width)) {
			util.CheckErr(util.InvalidMap)
		} //Check whether height matches mapdata.
	}
	if len(imported.mapdata) != imported.height {
		util.CheckErr(util.InvalidMap)
	} //Check whether height matches mapdata.

	//Check validity of file.
	imported.bitmap = convertToBitmap(imported.width, imported.height, imported.mapdata)
	return &imported

}

/*
	This function compares selected row's player data and answer data so it can judge if player painted wrong cell.
	This function will be called when player paints cell(NOT when checking).
*/

func (nm *Nonomap) CompareMap(answer Nonomap, x int) bool {

	if nm.mapdata[x] == answer.mapdata[x] {
		return true
	} else {
		return false
	}

}

func (nm *Nonomap) EmptyMap() *Nonomap {
	empty := NewNonomap("1/1/0")
	*empty = *nm
	for n := range empty.mapdata {
		empty.mapdata[n] = 0
	}
	return empty
}

/*
func (nm *Nonomap) CreateProblem() (horizontal [][]int, vertical [][]int) {

}
*/
func convertToBitmap(width int, height int, mapdata []int) [][]bool {

	bitmap := make([][]bool, height)
	for n := range bitmap {
		bitmap[n] = make([]bool, width)
	}

	for n, v := range mapdata {
		for i := 1; i <= width; i++ {
			bitmap[width-i][n] = (v%2 == 1)
			v = v / 2
		}
	}

	return bitmap

}
