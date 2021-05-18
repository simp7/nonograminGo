package standard

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type nonomap struct {
	Width  int
	Height int
	Bitmap [][]bool
}

/*
	nonomap is divided into 3 parts and has arguments equal or more than 3, which is separated by '/'.

	First two elements indicates Width and Height respectively.

	Rest elements indicates actual map which player has to solve.
	Each elements indicates map data of each line.
	They are designated by Bitmap, which 0 is blank and 1 is filled one.

	Since the size of int is 32bits, Width of maps can be equal or less than 32 mathematically.
	But because of display's limit, Width and Height can't be more than 25

	When it comes to player's map, 2 is checked one where player thinks that cell is blank.

	The extension of file is nm(*.nm)
*/

//Prototype returns prototype of nonogram.Map in this package.
func Prototype() nonogram.Map {
	return new(nonomap)
}

func (nm *nonomap) CopyWithBitmap(bitmap [][]bool) nonogram.Map {

	result := new(nonomap)

	result.Height = len(bitmap)
	result.Width = len(bitmap[0])
	result.Bitmap = bitmap

	return result

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
		tmp := 0

		for j := 0; j < nm.Width; j++ {

			if nm.Bitmap[i][j] {
				tmp++
				previousCell = true
			} else {
				if previousCell {
					horizontal[i] = append(horizontal[i], tmp)
					tmp = 0
				}
				previousCell = false
			}

		}

		if previousCell {
			horizontal[i] = append(horizontal[i], tmp)
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
		tmp := 0

		for j := 0; j < nm.Height; j++ {
			if nm.Bitmap[j][i] {
				tmp++
				previousCell = true
			} else {
				if previousCell {
					vertical[i] = append(vertical[i], tmp)
					tmp = 0
				}
				previousCell = false
			}
		}

		if previousCell {
			vertical[i] = append(vertical[i], tmp)
		}

		if len(vertical[i]) == 0 {
			vertical[i] = append(vertical[i], 0)
		}

	}

	return vertical

}

func (nm *nonomap) CreateProblem() nonogram.Problem {

	hData := nm.createHorizontalProblemData()
	vData := nm.createVerticalProblemData()

	hMax := getMaxLength(hData)
	vMax := getMaxLength(vData)

	return newProblem(hData, vData, hMax, vMax)

}

//This function returns Height of nonomap

func (nm *nonomap) GetHeight() int {
	return nm.Height
}

func (nm *nonomap) GetWidth() int {
	return nm.Width
}

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

func (nm *nonomap) CheckValidity() error {

	if nm.Height > nm.HeightLimit() || nm.Width > nm.WidthLimit() || nm.Height <= 0 || nm.Width <= 0 {
		return invalidMap
	}
	return nil

}

func (nm *nonomap) GetFormatter() nonogram.Formatter {
	return newFormatter()
}
