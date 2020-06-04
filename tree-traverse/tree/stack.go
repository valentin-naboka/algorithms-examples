package tree

// type Value interface{

// }

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

	v := s.data[0]
	s.data = s.data[1:len(s.data)]
	return v
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}
