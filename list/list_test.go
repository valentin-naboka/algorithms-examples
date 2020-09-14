package list

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/algorithms-examples/testutil"
)

func printListStateIfFailed(t *testing.T, list *UnrolledForwardList) {
	if t.Failed() {
		var result string
		for node := list.head; node != nil; node = node.next {
			result += fmt.Sprintf("%v", node.values)
		}
		splitter := strings.Repeat("-", len(result))
		t.Log(splitter)
		t.Log(result)
		t.Log(splitter)
	}
}

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

func checkValue(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		testutil.PrintCaller(t, 2)
		t.Errorf("Expected: %d, Actual: %d\n", expected, actual)
	}
}

// Singly linked list tests
// ========================
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

// Unrolled forward list tests
// ============================
func testUnrolledList(t *testing.T, expected []interface{}, actual *UnrolledForwardList) {
	if actual.GetBegin() == nil {
		testutil.PrintCaller(t, 2)
		t.Errorf("List is empty")
	}

	idx := 0
	for it := actual.GetBegin(); it != nil; {
		actualValue := it.GetValue()
		if !reflect.DeepEqual(actualValue, expected[idx]) {
			testutil.PrintCaller(t, 2)
			t.Errorf("Expected: %d, Actual: %d\n", expected[idx], actualValue)
		}
		it = it.MoveNext()
		idx++
	}
	printListStateIfFailed(t, actual)
}

func TestMoveToUnrolledList(t *testing.T) {
	list := NewUnrolledForwardList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i := len(data) - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}

	for i := 1; i < len(data); i++ {
		it := list.GetBegin().MoveTo(i)
		checkValue(t, data[i], it.GetValue())
	}

	it := list.GetBegin().MoveTo(2)
	it = it.MoveTo(13)
	checkValue(t, data[15], it.GetValue())

	it = list.GetBegin().MoveTo(1)
	it = it.MoveTo(2)
	checkValue(t, data[3], it.GetValue())
}

func TestPushFrontUnrolledList(t *testing.T) {
	list := NewUnrolledForwardList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := len(data) - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}

	testUnrolledList(t, data, list)
}

func TestInsertAfterUnrolledList(t *testing.T) {
	list := NewUnrolledForwardList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list.PushFront(data[0])

	it := list.GetBegin()
	for i := 1; i < len(data); i++ {
		list.InsertAfter(it, data[i])
		it.MoveNext()
	}

	testUnrolledList(t, data, list)
}

func newFullUnrolledForwardList(values ...interface{}) *UnrolledForwardList {
	lenght := len(values)
	if lenght == 0 {
		return &UnrolledForwardList{nil, 0}
	}

	head := newNode(nil)
	var node *node = nil

	for i := 0; i < lenght; i += maxChunkSize {
		if node == nil {
			node = head
		} else {
			node.next = newNode(nil)
			node = node.next
		}
		node.values = append(node.values, values[i:i+maxChunkSize]...)
	}

	if lenght%maxChunkSize != 0 {
		node.values = append(node.values, values[(lenght/maxChunkSize)*maxChunkSize:]...)
	}
	return &UnrolledForwardList{head, lenght}
}

func TestInsertAfterInMiddleUnrolledList(t *testing.T) {
	list := NewUnrolledForwardList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7}
	list.PushFront(data[0])

	it := list.GetBegin()
	for i := 1; i < len(data); i++ {
		list.InsertAfter(it, data[i])
		it.MoveNext()
	}

	it = list.GetBegin()
	list.InsertAfter(it, 12)

	testUnrolledList(t, []interface{}{1, 12, 2, 3, 4, 5, 6, 7}, list)

	it = list.GetBegin().MoveTo(2)
	list.InsertAfter(it, 23)

	testUnrolledList(t, []interface{}{1, 12, 2, 23, 3, 4, 5, 6, 7}, list)

	list = newFullUnrolledForwardList(1, 12, 2, 3, 4, 5, 6, 7)
	it = list.GetBegin().MoveTo(4)
	list.InsertAfter(it, 45)

	testUnrolledList(t, []interface{}{1, 12, 2, 3, 4, 45, 5, 6, 7}, list)
}

func TestPopFrontUnrolledList(t *testing.T) {
	list := NewUnrolledForwardList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	dataLen := len(data)
	for i := dataLen - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}

	for i := 1; i < dataLen; i++ {
		value := list.PopFront()
		checkValue(t, i, value)
		testUnrolledList(t, data[i:dataLen], list)
	}
}

func TestRemoveAfterUnrolledList(t *testing.T) {
	list := NewUnrolledForwardList()

	//TODO: copy-paste init list
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := len(data) - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}

	for it := list.GetBegin(); list.GetLength() != 1; {
		list.RemoveAfter(it)

		copy(data[1:], data[2:])
		data = data[:len(data)-1]

		testUnrolledList(t, data, list)
	}
}

//TODO: add cases to test internal storage representation
