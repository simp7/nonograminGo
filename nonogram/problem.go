package nonogram

//Problem is an interface of nonogram map's problem.
type Problem interface {
	Horizontal() ProblemUnit //Horizontal returns horizontal ProblemUnit.
	Vertical() ProblemUnit   //Vertical returns vertical ProblemUnit.
}
