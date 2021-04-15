package problem

import "github.com/simp7/nonograminGo/nonogram"

type problem struct {
	horizontal unit
	vertical   unit
}

func New(hProblem []string, vProblem []string, hMax int, vMax int) problem {
	return problem{NewUnit(hProblem, hMax), NewUnit(vProblem, vMax)}
}

func (p problem) Horizontal() nonogram.ProblemUnit {
	return p.horizontal
}

func (p problem) Vertical() nonogram.ProblemUnit {
	return p.vertical
}

//func horizontal(bitmap [][]bool) [][]int {
//
//	height := len(bitmap)
//	width := len(bitmap[0])
//	result := make([][]int, height)
//
//	for i := 0; i < height; i++ {
//
//		previousCell := false
//		temp := 0
//
//		for j := 0; j < width; j++ {
//
//			if bitmap[i][j] {
//				temp++
//				previousCell = true
//			} else {
//				if previousCell {
//					result[i] = append(result[i], temp)
//					temp = 0
//				}
//				previousCell = false
//			}
//
//		}
//
//		result[i] = processRest(result[i], previousCell, temp)
//
//	}
//
//	return result
//
//}
//
//func vertical(bitmap [][]bool) [][]int {
//
//	height := len(bitmap)
//	width := len(bitmap[0])
//	result := make([][]int, width)
//
//	for i := 0; i < width; i++ {
//
//		previousCell := false
//		temp := 0
//
//		for j := 0; j < height; j++ {
//			if bitmap[j][i] {
//				temp++
//				previousCell = true
//			} else {
//				if previousCell {
//					result[i] = append(result[i], temp)
//					temp = 0
//				}
//				previousCell = false
//			}
//		}
//
//		result[i] = processRest(result[i], previousCell, temp)
//
//	}
//
//	return result
//
//}
//
//func processRest(data []int, previous bool, last int) (result []int) {
//
//	result = make([]int, len(data))
//	copy(result, data)
//
//	if previous {
//		result = append(result, last)
//	}
//
//	if len(result) == 0 {
//		result = append(result, 0)
//	}
//
//	return
//
//}
//
//func (p problem) Horizontal() (hProblem []string, hMax int) {
//
//	hProblem = make([]string, len(p.horizontal.height))
//
//	for i := 0; i < nm.Height; i++ {
//		hProblem[i] = ""
//		for j := hMax; j > 0; j-- {
//			if len(hData[i]) < j {
//				hProblem[i] += "  "
//			} else {
//				if hData[i][len(hData[i])-j] < 10 {
//					hProblem[i] += " "
//				}
//				hProblem[i] += strconv.Itoa(hData[i][len(hData[i])-j])
//			}
//		}
//	}
//}
//
//func (p problem) Vertical() (vProblem []string, vMax int) {
//
//}
//
//func (p problem) Format() (hProblem []string, vProblem []string, hMax int, vMax int) {
//	vProblem = make([]string, vMax)
//
//
//
//	for i := vMax; i > 0; i-- {
//		vProblem[vMax-i] = ""
//		for j := 0; j < nm.Width; j++ {
//			if i > len(vData[j]) {
//				vProblem[vMax-i] += "  "
//			} else {
//				if vData[j][len(vData[j])-i] < 10 {
//					vProblem[vMax-i] += " "
//				}
//				vProblem[vMax-i] += strconv.Itoa(vData[j][len(vData[j])-i])
//			}
//		}
//	}
//	hMax *= 2
//	return
//
//}
