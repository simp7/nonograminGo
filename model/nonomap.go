package model

import (
	"nonograminGo/asset"
	"nonograminGo/util"
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

	if imported.height > asset.NumberHeightMax || imported.width > asset.NumberWidthMax || imported.height <= 0 || imported.width <= 0 {
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

func (nm *Nonomap) CompareValidity(x int, y int) bool {

	return nm.bitmap[y][x]

}

/*
	This function convert nonomap's data into problem data so player can solve with it.
	This function will be called in CreateProblemFormat().
*/

func (nm *Nonomap) createProblemData() (horizontal [][]int, vertical [][]int, hmax int, vmax int) {

	horizontal = make([][]int, nm.height)
	vertical = make([][]int, nm.width)
	hmax = 0
	vmax = 0

	for i := 0; i < nm.height; i++ {

		previousCell := false
		temp := 0

		for j := 0; j < nm.width; j++ {

			if nm.bitmap[i][j] == true {
				temp++
				previousCell = true
			} else {
				if previousCell == true {
					horizontal[i] = append(horizontal[i], temp)
					temp = 0
				}
				previousCell = false
			}

		}

		if previousCell == true {
			horizontal[i] = append(horizontal[i], temp)
		} else if len(horizontal[i]) == 0 {
			horizontal[i] = append(horizontal[i], 0)
		}

		if hmax < len(horizontal[i]) {
			hmax = len(horizontal[i])
		}

	}

	for i := 0; i < nm.width; i++ {

		previousCell := false
		temp := 0

		for j := 0; j < nm.height; j++ {
			if nm.bitmap[j][i] == true {
				temp++
				previousCell = true
			} else {
				if previousCell == true {
					vertical[i] = append(vertical[i], temp)
					temp = 0
				}
				previousCell = false
			}
		}

		if previousCell == true {
			vertical[i] = append(vertical[i], temp)
		} else if len(vertical[i]) == 0 {
			vertical[i] = append(vertical[i], 0)
		}

		if vmax < len(vertical[i]) {
			vmax = len(vertical[i])
		}

	}

	return
}

/*
	This function trim problem data to show player problem clearly.
	This function will be called when player enter the game.
*/

func (nm *Nonomap) CreateProblemFormat() (hProblem []string, vProblem []string, hmax int, vmax int) {

	hData, vData, hmax, vmax := nm.createProblemData()

	hProblem = make([]string, nm.height)
	vProblem = make([]string, vmax)

	for i := 0; i < nm.height; i++ {
		hProblem[i] = ""
		for j := hmax; j > 0; j-- {
			if len(hData[i]) < j {
				hProblem[i] += "  "
			} else {
				if hData[i][len(hData[i])-j] < 10 {
					hProblem[i] += " "
				}
				hProblem[i] += strconv.Itoa(hData[i][len(hData[i])-j])
			}
		}
	}

	for i := vmax; i > 0; i-- {
		vProblem[vmax-i] = ""
		for j := 0; j < nm.width; j++ {
			if i > len(vData[j]) {
				vProblem[vmax-i] += "  "
			} else {
				if vData[j][len(vData[j])-i] < 10 {
					vProblem[vmax-i] += " "
				}
				vProblem[vmax-i] += strconv.Itoa(vData[j][len(vData[j])-i])
			}
		}
	}
	hmax *= 2
	return

}

//This function returns height of nonomap

func (nm *Nonomap) GetHeight() int {
	return nm.height
}

//This function returns width of nonomap

func (nm *Nonomap) GetWidth() int {
	return nm.width
}

/*
	This function generates answer bitmap of Nonomap via mapdata.
	This function will be called when Nonomap is initialized.
*/

func convertToBitmap(width int, height int, mapdata []int) [][]bool {

	bitmap := make([][]bool, height)
	for n := range bitmap {
		bitmap[n] = make([]bool, width)
	}

	for n, v := range mapdata {
		for i := 1; i <= width; i++ {
			bitmap[n][width-i] = (v%2 == 1)
			v = v / 2
		}
	}

	return bitmap

}

func (nm *Nonomap) ShowBitMap() (result []string) {
	result = make([]string, nm.height)
	for i := 0; i < nm.height; i++ {
		for j := 0; j < nm.width; j++ {
			if nm.bitmap[i][j] {
				result[i] += "1"
			} else {
				result[i] += "0"
			}
		}
	}
	return
}

func (nm *Nonomap) ShowProblemHorizontal() (result []string) {
	a, _, _, _ := nm.createProblemData()

	result = make([]string, nm.height)
	for n := range a {
		for _, v := range a[n] {
			result[n] += strconv.Itoa(v)
		}
	}

	return
}

func (nm *Nonomap) ShowProblemVertical() (result []string) {
	_, b, _, _ := nm.createProblemData()

	result = make([]string, nm.width)
	for n := range b {
		for _, v := range b[n] {
			result[n] += strconv.Itoa(v)
		}
	}
	return
}

/*
	This function count total cells that sould be filled.
	The result will be used when judging whether player complete the map.
	This function will be called when player enter the game.
*/

func (nm *Nonomap) TotalCells() (total int) {

	total = 0

	for n := range nm.bitmap {
		for _, v := range nm.bitmap[n] {
			if v {
				total++
			}
		}
	}
	return

}
