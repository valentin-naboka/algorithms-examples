package queue

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 0)}
}

func (s *Queue) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Queue) Pop() interface{} {
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
