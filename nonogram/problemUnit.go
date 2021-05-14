package nonogram

//ProblemUnit is unit of problem.
type ProblemUnit interface {
	Get(idx int) []int //Get returns problem of selected index.
	Max() int          //Max returns max length of unit.
}
