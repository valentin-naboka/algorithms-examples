package list

type node struct {
	next   *node
	values []interface{}
}

var maxChunkSize int = 8
var midOfChunk int = maxChunkSize / 2

func newNode(next *node) *node {
	return &node{next, make([]interface{}, 0, maxChunkSize)}
}

func (n *node) isFull() bool {
	return len(n.values) >= maxChunkSize
}

func canMergeNodes(n *node) bool {
	return n.next != nil &&
		(len(n.values)-1+len(n.next.values)) < maxChunkSize
}

func removeValueFromNode(node *node, idx int) {
	if canMergeNodes(node) {
		copy(node.values[idx:], node.values[idx+1:])

		currentNodeLen := len(node.values)
		node.values = node.values[:currentNodeLen-1+len(node.next.values)]
		copy(node.values[currentNodeLen-1:], node.next.values)
		node.next = node.next.next
	} else {
		copy(node.values[idx:], node.values[idx+1:])
		node.values = node.values[:len(node.values)-1]
	}
}

type Iterator struct {
	//TODO: can node be null? make check
	node       *node
	currentIdx int
}

func (it *Iterator) MoveNext() *Iterator {

	if it.currentIdx+1 < len(it.node.values) {
		it.currentIdx++
	} else {
		it.node = it.node.next
		it.currentIdx = 0
		if it.node == nil {
			return nil
		}
	}
	return it
}

func (it *Iterator) GetValue() interface{} {
	return it.node.values[it.currentIdx]
}

func (it *Iterator) isLastValue() bool {
	return it.currentIdx == len(it.node.values)-1
}

type UnrolledForwardList struct {
	head   *node
	lenght int
}

func NewUnrolledForwardList(values ...interface{}) *UnrolledForwardList {
	lenght := len(values)
	if lenght == 0 {
		return &UnrolledForwardList{nil, 0}
	}

	head := newNode(nil)
	return &UnrolledForwardList{head, lenght}
}

func (l *UnrolledForwardList) GetBegin() *Iterator {
	if l.head == nil {
		return nil
	}
	return &Iterator{l.head, 0}
}

func (l UnrolledForwardList) GetLength() int {
	return l.lenght
}

//TODO: tests
func (l *UnrolledForwardList) InsertAfter(it *Iterator, v interface{}) *Iterator {
	if it == nil {
		return nil
	}

	if it.node.isFull() {
		it.node.next = newNode(it.node.next)

		it.node.next.values = append(it.node.next.values, it.node.values[midOfChunk:]...)
		it.node.next.values = append(it.node.next.values, v)

		it.node.values = it.node.values[0:midOfChunk]
		it.node = it.node.next
	} else {
		nodeLen := len(it.node.values)
		it.node.values = it.node.values[:nodeLen+1]
		insertIdx := it.currentIdx + 1
		copy(it.node.values[insertIdx:], it.node.values[it.currentIdx:nodeLen])
		it.node.values[insertIdx] = v
	}

	l.lenght++
	return it.MoveNext()
}

//TODO: tests
func (l *UnrolledForwardList) PushFront(v interface{}) {
	if l.head == nil {
		l.head = newNode(nil)
		l.head.values = append(l.head.values, v)
	} else if l.head.isFull() {
		l.head = newNode(l.head)

		l.head.values = append(l.head.values, v)
		l.head.values = append(l.head.values, l.head.next.values[0:midOfChunk]...)

		copy(l.head.next.values, l.head.next.values[midOfChunk:])
		l.head.next.values = l.head.next.values[:midOfChunk]

	} else {
		l.head.values = append(l.head.values, 0)
		copy(l.head.values[1:], l.head.values)
		l.head.values[0] = v
	}
	l.lenght++
}

//TODO: tests:
// 1. nextNodeLen == 7, 8
// 2. head.next.next != nil
func (l *UnrolledForwardList) PopFront() interface{} {
	if l.head == nil {
		panic("pop on empty list")
	}

	result := l.head.values[0]
	removeValueFromNode(l.head, 0)
	return result
}

//TODO: tests
//Returns iterator to the previous element
func (l *UnrolledForwardList) RemoveAfter(it *Iterator) *Iterator {
	if it.isLastValue() {
		if it.node.next == nil {
			panic("attempt to remove after the last item")
		}
		removeValueFromNode(it.node.next, 0)
	} else {
		removeValueFromNode(it.node, it.currentIdx+1)
	}
	l.lenght--
	return it
}
