package list

import (
	"reflect"
	"testing"

	"github.com/algorithms-examples/testutil"
)

func testList(t *testing.T, expected []interface{}, actual *SignlyLinkedList) {
	if actual.GetHead() == nil {
		testutil.PrintCaller(t, 2)
		t.Errorf("List is empty")
	}

	idx := 0
	for current := actual.GetHead(); current != nil; current = current.MoveNext() {
		actualValue := current.GetValue()
		if !reflect.DeepEqual(actualValue, expected[idx]) {
			testutil.PrintCaller(t, 2)
			t.Errorf("Expected: %d, Actual: %d\n", expected[idx], actualValue)
		}
		idx++
	}
}

func TestInsertAfterEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := len(data) - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}

	testList(t, data, list)
}

func TestPushFron(t *testing.T) {
	list := NewSinglyLinkedList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	currNode := list.GetHead()
	for _, v := range data {
		currNode = list.InsertAfter(currNode, v)
	}

	testList(t, data, list)
}

func TestInitList(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := NewSinglyLinkedList(data...)

	expected := 1
	for current := list.GetHead(); current != nil; current = current.MoveNext() {
		actual := testutil.ToInt(current.GetValue())
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
		expected++
	}
}

func TestRemoveAfter(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := NewSinglyLinkedList(data...)

	for current := list.GetHead(); list.GetLength() != 1; {
		list.RemoveAfter(current)
		testList(t, append(data[0:1], data[2:len(data)]...), list)
	}
	list.RemoveAfter(nil)
	if nil == list.GetHead() {
		t.Errorf("Expected: %v, Actual: %v", nil, list.GetHead())
	}
}
