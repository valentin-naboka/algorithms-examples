package list

import (
	"testing"

	"github.com/algorithms-examples/testutil"
)

func TestInsertAfter(t *testing.T) {
	list := NewSinglyLinkedListWithHead(InsertAfter(nil, 1))
	current := InsertAfter(list.Head(), 2)
	for i := 3; i < 11; i++ {
		current = InsertAfter(current, i)
	}

	expected := 1
	for current := list.Head(); current != nil; current = current.Next {
		actual := testutil.ToInt(current.Value)
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
		expected++
	}
}
