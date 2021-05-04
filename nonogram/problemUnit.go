package nonogram

//ProblemUnit is unit of problem.
//Get returns string array of problem.
//Max returns max length of unit.
type ProblemUnit interface {
	Get() []string
	Max() int
}
