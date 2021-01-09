package window

type Stack interface {
	Push(Window)
	Pop()
	Top() Window
	Size() int
}

type stack struct {
	windows []Window
}

func NewStack() Stack {
	s := new(stack)
	s.windows = make([]Window, 0)
	return s
}

func (s *stack) Push(window Window) {
	s.windows = append(s.windows, window)
}

func (s *stack) Pop() {
	s.windows = s.windows[:s.Size()-1]
}

func (s *stack) Top() Window {
	return s.windows[s.Size()-1]
}

func (s *stack) Size() int {
	return len(s.windows)
}
