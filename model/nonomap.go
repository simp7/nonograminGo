package model

import (
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/util"
	"math"
	"strconv"
)

/*
	This file deals with algorithms of whole game of nonogram.
	User's control or display should be separated from this file.
*/

type Nonomap interface {
	ShouldFilled(x, y int) bool
	CreateProblemFormat() (hProblem, vProblem []string, hMax, vMax int)
	GetHeight() int
	GetWidth() int
	ShowBitMap() []string
	ShowProblemHorizontal() []string
	ShowProblemVertical() []string
	FilledTotal() int
	checkValidity()
}

type nonomap struct {
	Width   int
	Height  int
	MapData []int
	Bitmap  [][]bool
}

/*
	nonomap is divided into 3 parts and has arguments equal or more than 3, which is separated by '/'.

	First two elements indicates Width and Height respectively.

	Rest elements indicates actual map which player has to solve.
	Each elements indicates map data of each line.
	They are designated by Bitmap, which 0 is blank and 1 is filled one.

	Because the size of int is 32bits, Width of maps can't be more than 32 mathematically.

	But because of display's limit, Width and Height can't be more than 25

	When it comes to player's map, 2 is checked one where player thinks that cell is blank.

	The extension of file is nm(*.nm)
*/

/*
	This function compares selected row's player data and answer data so it can judge if player painted wrong cell.
	This function will be called when player paints cell(NOT when checking).
*/

func NewNonomap() Nonomap {
	return new(nonomap)
}

func (nm *nonomap) ShouldFilled(x int, y int) bool {

	return nm.Bitmap[y][x]

}

/*
	This function convert nonomap's data into problem data so player can solve with it.
	This function will be called in CreateProblemFormat().
*/

func (nm *nonomap) createProblemData() (horizontal [][]int, vertical [][]int, hMax int, vMax int) {

	horizontal = make([][]int, nm.Height)
	vertical = make([][]int, nm.Width)
	hMax = 0
	vMax = 0

	for i := 0; i < nm.Height; i++ {

		previousCell := false
		temp := 0

		for j := 0; j < nm.Width; j++ {

			if nm.Bitmap[i][j] == true {
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

	for i := 0; i < nm.Width; i++ {

		previousCell := false
		temp := 0

		for j := 0; j < nm.Height; j++ {
			if nm.Bitmap[j][i] == true {
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

func (nm *nonomap) CreateProblemFormat() (hProblem []string, vProblem []string, hMax int, vMax int) {

	hData, vData, hMax, vMax := nm.createProblemData()

	hProblem = make([]string, nm.Height)
	vProblem = make([]string, vMax)

	for i := 0; i < nm.Height; i++ {
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
		for j := 0; j < nm.Width; j++ {
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

//This function returns Height of nonomap

func (nm *nonomap) GetHeight() int {
	return nm.Height
}

//This function returns Width of nonomap

func (nm *nonomap) GetWidth() int {
	return nm.Width
}

/*
	This function generates answer Bitmap of nonomap via MapData.
	This function will be called when nonomap is initialized.
*/

func (nm *nonomap) ShowBitMap() (result []string) {
	result = make([]string, nm.Height)
	for i := 0; i < nm.Height; i++ {
		for j := 0; j < nm.Width; j++ {
			if nm.Bitmap[i][j] {
				result[i] += "1"
			} else {
				result[i] += "0"
			}
		}
	}
	return
}

func (nm *nonomap) ShowProblemHorizontal() (result []string) {
	a, _, _, _ := nm.createProblemData()

	result = make([]string, nm.Height)
	for n := range a {
		for _, v := range a[n] {
			result[n] += strconv.Itoa(v)
		}
	}

	return
}

func (nm *nonomap) ShowProblemVertical() (result []string) {
	_, b, _, _ := nm.createProblemData()

	result = make([]string, nm.Width)
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

func (nm *nonomap) FilledTotal() (total int) {

	total = 0

	for n := range nm.Bitmap {
		for _, v := range nm.Bitmap[n] {
			if v {
				total++
			}
		}
	}
	return

}

func (nm *nonomap) checkValidity() {

	setting := asset.GetSetting()
	hMax := setting.HeightMax
	wMax := setting.WidthMax

	if nm.Height > hMax || nm.Width > wMax || nm.Height <= 0 || nm.Width <= 0 {
		util.CheckErr(util.InvalidMap)
	} //Check if Height and Width meets criteria of size.

	//Extract map's answer content file.

	for _, v := range nm.MapData {
		if float64(v) >= math.Pow(2, float64(nm.Width)) {
			util.CheckErr(util.InvalidMap)
		} //Check whether Height matches MapData.
	}

	if len(nm.MapData) != nm.Height {
		util.CheckErr(util.InvalidMap)
	} //Check whether Height matches MapData.

}
