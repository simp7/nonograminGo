package nonogram

type Problem interface {
	Horizontal() ProblemUnit
	Vertical() ProblemUnit
}
