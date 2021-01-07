package window

type Stack interface {
	push(Window)
	pop()
	front() Window
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

func (s *stack) push(window Window) {
	s.windows = append(s.windows, window)
}

func (s *stack) pop() {
	s.windows = s.windows[:s.Size()-1]
}

func (s *stack) front() Window {
	return s.windows[s.Size()-1]
}

func (s *stack) Size() int {
	return len(s.windows)
}
