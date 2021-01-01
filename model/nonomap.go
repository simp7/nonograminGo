package model

import (
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/util"
	"math"
	"strconv"
	"strings"
)

/*
	This file deals with algorithms of whole game of nonogram.
	User's control or display should be separated from this file.
*/

type Nonomap struct {
	width   int
	height  int
	mapData []int
	bitmap  [][]bool
}

/*
	nonomap is divided into 3 parts and has arguments equal or more than 3, which is separated by '/'.

	First two elements indicates width and height respectively.

	Rest elements indicates actual map which player has to solve.
	Each elements indicates map data of each line.
	They are designated by bitmap, which 0 is blank and 1 is filled one.

	Because the size of int is 32bits, width of maps can't be more than 32 mathematically.

	But because of display's limit, width and height can't be more than 25

	When it comes to player's map, 2 is checked one where player thinks that cell is blank.

	The extension of file is nm(*.nm)
*/

func NewNonomap(data string) *Nonomap {

	var imported Nonomap
	var err error
	setting := asset.GetSetting()

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")
	//Extract all data from wanted file.

	imported.width, err = strconv.Atoi(elements[0])
	imported.height, err = strconv.Atoi(elements[1])
	util.CheckErr(err)
	//Extract map's size from file.

	for _, v := range elements[2:] {
		temp, err := strconv.Atoi(v)
		imported.mapData = append(imported.mapData, temp)
		util.CheckErr(err)
	}
	//Extract map's answer from file.

	if imported.height > setting.HeightMax || imported.width > setting.WidthMax || imported.height <= 0 || imported.width <= 0 {
		util.CheckErr(util.InvalidMap)
	} //Check if height and width meets criteria of size.

	for _, v := range imported.mapData {
		if float64(v) >= math.Pow(2, float64(imported.width)) {
			util.CheckErr(util.InvalidMap)
		} //Check whether height matches mapData.
	}
	if len(imported.mapData) != imported.height {
		util.CheckErr(util.InvalidMap)
	} //Check whether height matches mapData.

	//Check validity of file.
	imported.bitmap = convertToBitmap(imported.width, imported.height, imported.mapData)
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

func (nm *Nonomap) createProblemData() (horizontal [][]int, vertical [][]int, hMax int, vMax int) {

	horizontal = make([][]int, nm.height)
	vertical = make([][]int, nm.width)
	hMax = 0
	vMax = 0

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

		if hMax < len(horizontal[i]) {
			hMax = len(horizontal[i])
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

		if vMax < len(vertical[i]) {
			vMax = len(vertical[i])
		}

	}

	return
}

/*
	This function trim problem data to show player problem clearly.
	This function will be called when player enter the game.
*/

func (nm *Nonomap) CreateProblemFormat() (hProblem []string, vProblem []string, hMax int, vMax int) {

	hData, vData, hMax, vMax := nm.createProblemData()

	hProblem = make([]string, nm.height)
	vProblem = make([]string, vMax)

	for i := 0; i < nm.height; i++ {
		hProblem[i] = ""
		for j := hMax; j > 0; j-- {
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

	for i := vMax; i > 0; i-- {
		vProblem[vMax-i] = ""
		for j := 0; j < nm.width; j++ {
			if i > len(vData[j]) {
				vProblem[vMax-i] += "  "
			} else {
				if vData[j][len(vData[j])-i] < 10 {
					vProblem[vMax-i] += " "
				}
				vProblem[vMax-i] += strconv.Itoa(vData[j][len(vData[j])-i])
			}
		}
	}
	hMax *= 2
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
	This function generates answer bitmap of Nonomap via mapData.
	This function will be called when Nonomap is initialized.
*/

func convertToBitmap(width int, height int, mapData []int) [][]bool {

	bitmap := make([][]bool, height)
	for n := range bitmap {
		bitmap[n] = make([]bool, width)
	}

	for n, v := range mapData {
		for i := 1; i <= width; i++ {
			bitmap[n][width-i] = v%2 == 1
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
	This function count total cells that should be filled.
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
