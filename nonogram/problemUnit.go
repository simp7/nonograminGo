package nonogram

type ProblemUnit interface {
	Get() []string
	Max() int
}
