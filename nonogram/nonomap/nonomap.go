package nonomap

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/util"
	"math"
	"strconv"
)

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

func New() nonogram.Map {
	return new(nonomap)
}

func (nm *nonomap) ShouldFilled(x int, y int) bool {

	return nm.Bitmap[y][x]

}

func getMaxLength(data [][]int) int {
	max := 0
	for _, v := range data {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func (nm *nonomap) createHorizontalProblemData() [][]int {

	horizontal := make([][]int, nm.Height)

	for i := 0; i < nm.Height; i++ {

		previousCell := false
		temp := 0

		for j := 0; j < nm.Width; j++ {

			if nm.Bitmap[i][j] {
				temp++
				previousCell = true
			} else {
				if previousCell {
					horizontal[i] = append(horizontal[i], temp)
					temp = 0
				}
				previousCell = false
			}

		}

		if previousCell {
			horizontal[i] = append(horizontal[i], temp)
		}

		if len(horizontal[i]) == 0 {
			horizontal[i] = append(horizontal[i], 0)
		}

	}

	return horizontal

}

func (nm *nonomap) createVerticalProblemData() [][]int {

	vertical := make([][]int, nm.Width)

	for i := 0; i < nm.Width; i++ {

		previousCell := false
		temp := 0

		for j := 0; j < nm.Height; j++ {
			if nm.Bitmap[j][i] {
				temp++
				previousCell = true
			} else {
				if previousCell {
					vertical[i] = append(vertical[i], temp)
					temp = 0
				}
				previousCell = false
			}
		}

		if previousCell {
			vertical[i] = append(vertical[i], temp)
		}

		if len(vertical[i]) == 0 {
			vertical[i] = append(vertical[i], 0)
		}

	}

	return vertical

}

/*
	This function trim problem data to show player problem clearly.
	This function will be called when player enter the game.
*/

func (nm *nonomap) CreateProblemFormat() (hProblem []string, vProblem []string, hMax int, vMax int) {

	hData := nm.createHorizontalProblemData()
	vData := nm.createVerticalProblemData()

	hMax = getMaxLength(hData)
	vMax = getMaxLength(vData)

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

func (nm *nonomap) BitmapToStrings() []string {
	result := make([]string, nm.Height)
	for y := 0; y < nm.Height; y++ {
		result[y] = nm.rowToString(y)
	}
	return result
}

func (nm *nonomap) rowToString(y int) (result string) {
	for x := 0; x < nm.Width; x++ {
		result += nm.cellToString(x, y)
	}
	return
}

func (nm *nonomap) cellToString(x, y int) string {
	if nm.Bitmap[y][x] {
		return "1"
	}
	return "0"
}

func (nm *nonomap) ShowProblemHorizontal() (result []string) {

	d := nm.createHorizontalProblemData()

	result = make([]string, nm.Height)
	for n := range d {
		for _, v := range d[n] {
			result[n] += strconv.Itoa(v)
		}
	}

	return

}

func (nm *nonomap) ShowProblemVertical() (result []string) {

	d := nm.createVerticalProblemData()

	result = make([]string, nm.Width)
	for n := range d {
		for _, v := range d[n] {
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
		total += nm.countRow(n)
	}

	return

}

func (nm *nonomap) countRow(y int) int {
	result := 0
	for _, v := range nm.Bitmap[y] {
		if v {
			result++
		}
	}
	return result
}

func (nm *nonomap) HeightLimit() int {
	return 30
}

func (nm *nonomap) WidthLimit() int {
	return 30
}

func (nm *nonomap) CheckValidity() {
	util.CheckErr(nm.checkSize())
	util.CheckErr(nm.checkWidth())
	util.CheckErr(nm.checkHeight())
}

func (nm *nonomap) checkSize() (err error) {
	if nm.Height > nm.HeightLimit() || nm.Width > nm.WidthLimit() || nm.Height <= 0 || nm.Width <= 0 {
		err = util.InvalidMap
	}
	return
}

func (nm *nonomap) checkWidth() (err error) {
	for _, v := range nm.MapData {
		if float64(v) >= math.Pow(2, float64(nm.Width)) {
			err = util.InvalidMap
		}
	}
	return
}

func (nm *nonomap) checkHeight() (err error) {
	if len(nm.MapData) != nm.Height {
		err = util.InvalidMap
	}
	return
}
