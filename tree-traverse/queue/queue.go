package queue

type Value interface{}

type Queue struct {
	data []Value
}

func NewQueue() *Queue {
	return &Queue{make([]Value, 0)}
}

func (s *Queue) Put(v Value) {
	s.data = append(s.data, v)
}

func (s *Queue) Pop() Value {
	if s.IsEmpty() {
		panic("Attempt to pop empty queue.")
	}

	queueLen := len(s.data)
	v := s.data[0]
	s.data = s.data[1:queueLen]
	return v
}

func (s Queue) IsEmpty() bool {
	return len(s.data) == 0
}
