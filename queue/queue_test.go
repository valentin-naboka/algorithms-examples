package queue

import (
	"testing"

	"github.com/algorithms-examples/testutil"
)

type data []int

func newQueue(d data) *Queue {
	stack := NewQueue()
	for _, v := range d {
		stack.Push(v)
	}
	return stack
}

func TestPush(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stack := newQueue(input)
	for i := 0; !stack.IsEmpty(); i++ {
		actual := testutil.ToInt(stack.Pop())
		expected := input[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}

func TestPushPop(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	queue := NewQueue()
	expectedIdx := 0
	for i, v := range input {
		if i != 0 && i%3 == 0 {
			actual := testutil.ToInt(queue.Pop())
			expected := input[expectedIdx]
			expectedIdx++
			if actual != expected {
				t.Errorf("Expected: %d, Actual: %d", expected, actual)
			}
		}
		queue.Push(v)
	}

	for i := 2; !queue.IsEmpty(); i++ {
		actual := testutil.ToInt(queue.Pop())
		expected := input[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}
