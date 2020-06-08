package tree

type Value interface{}

type Stack struct {
	data []Value
}

func NewStack() *Stack {
	return &Stack{make([]Value, 0)}
}

func (s *Stack) Push(v Value) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() Value {
	if s.IsEmpty() {
		panic("Attempt to pop empty stack.")
	}

	stackLen := len(s.data)
	v := s.data[stackLen-1]
	s.data = s.data[0 : len(s.data)-1]
	return v
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}
