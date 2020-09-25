package list

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/algorithms-examples/testutil"
)

func checkValue(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		testutil.PrintCaller(t, 2)
		t.Errorf("Expected: %d, Actual: %d\n", expected, actual)
	}
}

// Singly linked list tests
// ========================
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

//====> InsertAfter tests
func TestInsertAfterEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := len(data) - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}

	testList(t, data, list)
}

//====> PushFront tests
func TestPushFront(t *testing.T) {
	list := NewSinglyLinkedList()

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	currNode := list.GetHead()
	for _, v := range data {
		currNode = list.InsertAfter(currNode, v)
	}

	testList(t, data, list)
}

//====> Init list tests
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

//====> RemoveAfter tests
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
func pushFront(list *UnrolledForwardList, data ...interface{}) {
	for i := 0; i < len(data); i++ {
		list.PushFront(data[i])
	}
}

func createUnrolledForwardList(data []interface{}) *UnrolledForwardList {
	list := NewUnrolledForwardList()
	for i := len(data) - 1; i >= 0; i-- {
		list.PushFront(data[i])
	}
	return list
}

func printListStateIfFailed(t *testing.T, list *UnrolledForwardList) {
	if t.Failed() {
		var result string
		for node := list.head; node != nil; node = node.next {
			result += fmt.Sprintf("%v", node.values)
		}

		headerName := " actual list "
		resultLen := len(result)
		headerNameLen := len(headerName)
		if resultLen <= headerNameLen {
			resultLen = headerNameLen + 6
		}
		headerFiller := strings.Repeat("-", (resultLen-len(headerName))/2)
		header := headerFiller + headerName + headerFiller
		t.Log(header)
		t.Log(result)
		footer := strings.Repeat("-", resultLen)
		t.Log(footer)
	}
}

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

func testUnrolledForwardListInternals(t *testing.T, actual *UnrolledForwardList, expected ...[]interface{}) {
	if actual.GetBegin() == nil {
		testutil.PrintCaller(t, 2)
		t.Errorf("list is empty")
	}

	expectedLen := 0
	node := actual.head
	for _, expectedValues := range expected {
		if node == nil {
			t.Error("the end of the list reached early ")
			break
		}
		expectedLen += len(expectedValues)
		if !reflect.DeepEqual(expectedValues, node.values) {
			testutil.PrintCaller(t, 2)
			t.Errorf("Expected: %d, Actual: %d\n", expectedValues, node.values)
		}
		node = node.next
	}
	if node != nil {
		t.Error("the actual list still has untested values")
	}

	if expectedLen != actual.GetLength() {
		testutil.PrintCaller(t, 2)
		t.Errorf("the actual length is unexpected - expected: %d, actual: %d", expected, actual.GetLength())
	}
	printListStateIfFailed(t, actual)
}

//====> MoveTo tests
func TestMoveToUnrolledList(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	list := createUnrolledForwardList(data)

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

//====> PushFront tests
func TestPushFrontUnrolledList(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := createUnrolledForwardList(data)

	testUnrolledList(t, data, list)
}

func TestPushFrontUnrolledListInternals(t *testing.T) {
	data := []interface{}{4, 3, 2, 1}
	list := createUnrolledForwardList(data)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{4, 3, 2, 1})

	pushFront(list, 5, 6, 7, 8)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{8, 7, 6, 5, 4, 3, 2, 1})

	pushFront(list, 9)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{9, 8, 7, 6, 5},
		[]interface{}{4, 3, 2, 1})

	pushFront(list, 10, 11, 12)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{12, 11, 10, 9, 8, 7, 6, 5},
		[]interface{}{4, 3, 2, 1})

	pushFront(list, 13, 14, 15, 16)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{16, 15, 14, 13, 12, 11, 10, 9},
		[]interface{}{8, 7, 6, 5},
		[]interface{}{4, 3, 2, 1})
}

//====> InsertAfter tests
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
	length := len(values)
	if length == 0 {
		return &UnrolledForwardList{nil, 0}
	}

	head := newNode(nil)
	var node *node = nil

	for i := 0; i < length; i += maxChunkSize {
		if node == nil {
			node = head
		} else {
			node.next = newNode(nil)
			node = node.next
		}
		node.values = append(node.values, values[i:i+maxChunkSize]...)
	}

	if length%maxChunkSize != 0 {
		node.values = append(node.values, values[(length/maxChunkSize)*maxChunkSize:]...)
	}
	return &UnrolledForwardList{head, length}
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

func insertAfter(list *UnrolledForwardList, it *Iterator, data ...interface{}) {
	for _, v := range data {
		list.InsertAfter(it, v)
		it.MoveNext()
	}
}

func TestInsertAfterUnrolledListInternals(t *testing.T) {
	list := NewUnrolledForwardList()
	list.PushFront(1)
	it := list.GetBegin()

	insertAfter(list, it, 2, 3, 4, 5, 6, 7, 8)
	testUnrolledForwardListInternals(t, list, []interface{}{1, 2, 3, 4, 5, 6, 7, 8})

	list.InsertAfter(it, 9)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4},
		[]interface{}{5, 6, 7, 8, 9})

	it = list.GetBegin().MoveTo(3)
	insertAfter(list, it, 4, 4, 4, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 4, 4, 4, 4},
		[]interface{}{5, 6, 7, 8, 9})

	it = list.GetBegin().MoveTo(9)
	insertAfter(list, it, 6, 6, 6)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 4, 4, 4, 4},
		[]interface{}{5, 6, 6, 6, 6, 7, 8, 9})

	it = list.GetBegin().MoveTo(8)
	insertAfter(list, it, 5, 5)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 4, 4, 4, 4},
		[]interface{}{5, 5, 5, 6, 6, 6},
		[]interface{}{6, 7, 8, 9})

	it = list.GetBegin().MoveTo(10)
	insertAfter(list, it, 56, 56)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 4, 4, 4, 4},
		[]interface{}{5, 5, 5, 56, 56, 6, 6, 6},
		[]interface{}{6, 7, 8, 9})

	it = list.GetBegin().MoveTo(list.length - 1)
	insertAfter(list, it, 10, 11, 12, 13)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 4, 4, 4, 4},
		[]interface{}{5, 5, 5, 56, 56, 6, 6, 6},
		[]interface{}{6, 7, 8, 9, 10, 11, 12, 13})

	it = list.GetBegin().MoveTo(1)
	insertAfter(list, it, 23)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 23, 3, 4},
		[]interface{}{4, 4, 4, 4},
		[]interface{}{5, 5, 5, 56, 56, 6, 6, 6},
		[]interface{}{6, 7, 8, 9, 10, 11, 12, 13})

	it = list.GetBegin().MoveTo(14)
	insertAfter(list, it, 66)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 23, 3, 4},
		[]interface{}{4, 4, 4, 4},
		[]interface{}{5, 5, 5, 56},
		[]interface{}{56, 6, 66, 6, 6},
		[]interface{}{6, 7, 8, 9, 10, 11, 12, 13})
}

//====> PopFront tests
func TestPopFrontUnrolledList(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := createUnrolledForwardList(data)

	dataLen := len(data)
	for i := 1; i < dataLen; i++ {
		value := list.PopFront()
		checkValue(t, i, value)
		testUnrolledList(t, data[i:dataLen], list)
	}
}

func newCustomFilledUnrolledForwardList(values ...[]interface{}) *UnrolledForwardList {
	if len(values) == 0 {
		return &UnrolledForwardList{nil, 0}
	}

	head := newNode(nil)
	var node *node = nil

	length := 0
	for _, v := range values {
		if node == nil {
			node = head
		} else {
			node.next = newNode(nil)
			node = node.next
		}
		node.values = node.values[:len(v)]
		copy(node.values, v)
		length += len(v)
	}

	return &UnrolledForwardList{head, length}
}

func popNTimes(l *UnrolledForwardList, count int) {
	for i := 0; i < count; i++ {
		l.PopFront()
	}
}

func TestPopFrontFromFullUnrolledListInternals(t *testing.T) {
	list := newCustomFilledUnrolledForwardList(
		[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12, 13, 14, 15, 16})

	popNTimes(list, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12, 13, 14, 15, 16})

	popNTimes(list, 3)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{8},
		[]interface{}{9, 10, 11, 12, 13, 14, 15, 16})

	popNTimes(list, 1)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{9, 10, 11, 12, 13, 14, 15, 16})

	popNTimes(list, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{13, 14, 15, 16})

	popNTimes(list, 4)
	var expected *node = nil
	checkValue(t, expected, list.head)
}

func TestPopFrontFromHalfFilledUnrolledListInternals(t *testing.T) {
	list := newCustomFilledUnrolledForwardList(
		[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18},
		[]interface{}{19})

	popNTimes(list, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18},
		[]interface{}{19})

	popNTimes(list, 1)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{6, 7, 8, 9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18},
		[]interface{}{19})

	popNTimes(list, 5)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18},
		[]interface{}{19})

	popNTimes(list, 1)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{12, 13, 14, 15, 16, 17, 18},
		[]interface{}{19})

	popNTimes(list, 1)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{13, 14, 15, 16, 17, 18, 19})

	popNTimes(list, 6)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{19})

	popNTimes(list, 1)
	var expected *node = nil
	checkValue(t, expected, list.head)
}

//====> RemoveAfter tests
func TestRemoveAfterUnrolledList(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := createUnrolledForwardList(data)

	for it := list.GetBegin(); list.GetLength() != 1; {
		list.RemoveAfter(it)

		copy(data[1:], data[2:])
		data = data[:len(data)-1]

		testUnrolledList(t, data, list)
	}
}

func removeAfter(list *UnrolledForwardList, it *Iterator, step, count int) {

	for i := 0; i < count; i++ {
		list.RemoveAfter(it)
		it.MoveTo(step)
	}
}
func TestRemoveAfterFromFullUnrolledListInternals(t *testing.T) {
	list := newCustomFilledUnrolledForwardList(
		[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12, 13, 14, 15, 16})

	it := list.GetBegin()
	removeAfter(list, it, 1, 4)

	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 3, 5, 7},
		[]interface{}{9, 10, 11, 12, 13, 14, 15, 16})

	removeAfter(list, it, 1, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 3, 5, 7},
		[]interface{}{9, 11, 13, 15})

	list.RemoveAfter(list.GetBegin().MoveNext())
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 3, 7, 9, 11, 13, 15})

	removeAfter(list, list.GetBegin(), 1, 3)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 7, 11, 15})

	removeAfter(list, list.GetBegin(), 0, 3)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1})
}

func TestRemoveAfterFromHalfFilledUnrolledListInternals(t *testing.T) {
	list := newCustomFilledUnrolledForwardList(
		[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18},
		[]interface{}{19},
		[]interface{}{20})

	it := list.GetBegin().MoveTo(list.GetLength() - 3)
	list.RemoveAfter(it)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18},
		[]interface{}{20})

	list.RemoveAfter(it)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		[]interface{}{9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18})

	it = list.GetBegin().MoveNext()
	removeAfter(list, it, 0, 3)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 6, 7, 8},
		[]interface{}{9, 10, 11, 12},
		[]interface{}{13, 14, 15, 16, 17, 18})

	it = list.GetBegin().MoveTo(4)
	removeAfter(list, it, 0, 3)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 6, 7, 8},
		[]interface{}{12, 13, 14, 15, 16, 17, 18})

	it = list.GetBegin().MoveTo(4)
	list.RemoveAfter(it)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 2, 6, 7, 8},
		[]interface{}{13, 14, 15, 16, 17, 18})

	it = list.GetBegin()
	removeAfter(list, it, 0, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 13, 14, 15, 16, 17, 18})

	it = it.MoveTo(list.GetLength() - 2)
	list.RemoveAfter(it)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 13, 14, 15, 16, 17})

	it = list.GetBegin()
	removeAfter(list, it, 0, 4)
	testUnrolledForwardListInternals(t, list,
		[]interface{}{1, 17})
}
