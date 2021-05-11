package nonogram

//ProblemUnit is unit of problem.
type ProblemUnit interface {
	Get() []string //Get returns string array of problem.
	Max() int      //Max returns max length of unit.
}
