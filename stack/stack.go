package stack

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0)}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("Attempt to pop empty stack.")
	}

	stackLen := len(s.data)
	v := s.data[stackLen-1]
	s.data = s.data[0 : stackLen-1]
	return v
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}
