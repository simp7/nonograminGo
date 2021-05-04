package nonogram

//Problem is an interface of nonogram map's problem
//Horizontal returns horizontal ProblemUnit.
//Vertical returns vertical ProblemUnit.
type Problem interface {
	Horizontal() ProblemUnit
	Vertical() ProblemUnit
}
