package stack

import (
	"testing"

	"github.com/algorithms-examples/testutil"
)

type data []int

func newStack(d data) *Stack {
	stack := NewStack()
	for _, v := range d {
		stack.Push(v)
	}
	return stack
}

func TestPush(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stack := newStack(input)
	for i := len(input) - 1; !stack.IsEmpty(); i-- {
		actual := testutil.ToInt(stack.Pop())
		expected := input[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}

func TestPushPop(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stack := NewStack()
	for i, v := range input {
		if i != 0 && i%3 == 0 {
			actual := testutil.ToInt(stack.Pop())
			expected := input[i-1]
			if actual != expected {
				t.Errorf("Expected: %d, Actual: %d", expected, actual)
			}
		}
		stack.Push(v)
	}

	for i := len(input) - 1; !stack.IsEmpty(); i-- {
		actual := testutil.ToInt(stack.Pop())
		expected := input[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
		if i != 0 && i%3 == 0 {
			i--
		}
	}
}
