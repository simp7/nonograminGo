package standard

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type problem struct {
	horizontal unit
	vertical   unit
}

func newProblem(hProblem []string, vProblem []string, hMax int, vMax int) problem {
	return problem{newUnit(hProblem, hMax), newUnit(vProblem, vMax)}
}

func (p problem) Horizontal() nonogram.ProblemUnit {
	return p.horizontal
}

func (p problem) Vertical() nonogram.ProblemUnit {
	return p.vertical
}
