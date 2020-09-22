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

type Iterator struct {
	node       *node
	currentIdx int
}

func (it *Iterator) MoveNext() *Iterator {

	if it.currentIdx+1 < len(it.node.values) {
		it.currentIdx++
		return it
	}

	if it.node.next == nil {
		return nil
	}
	it.node = it.node.next
	it.currentIdx = 0
	return it
}

func (it *Iterator) MoveTo(shift int) *Iterator {

	if shift < len(it.node.values)-it.currentIdx {
		it.currentIdx += shift
		return it
	}

	for step := len(it.node.values) - it.currentIdx; shift >= step; step = len(it.node.values) {
		shift -= step
		if it.node.next == nil {
			panic("shift is out of range")
		}
		it.node = it.node.next
	}

	it.currentIdx = shift
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
	length int
}

func NewUnrolledForwardList() *UnrolledForwardList {
	return &UnrolledForwardList{nil, 0}
}

func (l *UnrolledForwardList) GetBegin() *Iterator {
	if l.head == nil || len(l.head.values) == 0 {
		return nil
	}
	return &Iterator{l.head, 0}
}

func (l UnrolledForwardList) GetLength() int {
	return l.length
}

func insertValue(values []interface{}, value interface{}, pos int) []interface{} {
	nodeLen := len(values)
	values = values[:nodeLen+1]
	copy(values[pos+1:], values[pos:nodeLen])
	values[pos] = value
	return values
}

func (l *UnrolledForwardList) InsertAfter(it *Iterator, v interface{}) {
	if it == nil {
		panic("insert after nil iterator")
	}

	l.length++
	if !it.node.isFull() {
		it.node.values = insertValue(it.node.values, v, it.currentIdx+1)
		return
	}

	it.node.next = newNode(it.node.next)

	if it.currentIdx < midOfChunk {
		it.node.next.values = append(it.node.next.values, it.node.values[midOfChunk:]...)
		it.node.values = insertValue(it.node.values[0:midOfChunk], v, it.currentIdx+1)
		return
	}

	it.node.next.values = append(it.node.next.values, it.node.values[midOfChunk:it.currentIdx+1]...)
	insertPos := it.currentIdx - midOfChunk + 1

	if insertPos < midOfChunk {
		it.node.next.values = append(it.node.next.values, v)
		it.node.next.values = append(it.node.next.values, it.node.values[it.currentIdx+1:]...)
	} else {
		it.node.next.values = append(it.node.next.values, v)
	}

	it.node.values = it.node.values[0:midOfChunk]
	it.node = it.node.next
	it.currentIdx -= midOfChunk
}

func (l *UnrolledForwardList) PushFront(v interface{}) {
	l.length++
	if l.head == nil {
		l.head = newNode(nil)
		l.head.values = append(l.head.values, v)
		return
	}
	if l.head.isFull() {
		l.head = newNode(l.head)

		l.head.values = append(l.head.values, v)
		l.head.values = append(l.head.values, l.head.next.values[0:midOfChunk]...)

		copy(l.head.next.values, l.head.next.values[midOfChunk:])
		l.head.next.values = l.head.next.values[:midOfChunk]
		return
	}

	l.head.values = append(l.head.values, 0)
	copy(l.head.values[1:], l.head.values)
	l.head.values[0] = v
	return
}

func canMergeNodes(n *node) bool {
	return n.next != nil &&
		//NOTE: -1 to make node up to 7 elements -> make test
		(len(n.values)-1+len(n.next.values)) < maxChunkSize
}

func removeValueFromNode(node *node, idx int) *node {
	if !canMergeNodes(node) {
		nodeLen := len(node.values)
		if nodeLen == 1 {
			return node.next
		}

		copy(node.values[idx:], node.values[idx+1:])
		node.values = node.values[:nodeLen-1]
		return node
	}

	copy(node.values[idx:], node.values[idx+1:])
	currentNodeLen := len(node.values)
	node.values = node.values[:currentNodeLen-1+len(node.next.values)]
	copy(node.values[currentNodeLen-1:], node.next.values)
	node.next = node.next.next
	return node
}

func (l *UnrolledForwardList) PopFront() interface{} {
	if l.head == nil {
		panic("pop on empty list")
	}

	result := l.head.values[0]
	l.head = removeValueFromNode(l.head, 0)
	return result
}

//TODO: tests
//TODO: test l.length in all cases
func (l *UnrolledForwardList) RemoveAfter(it *Iterator) {
	if !it.isLastValue() {
		//TODO: assing node
		removeValueFromNode(it.node, it.currentIdx+1)
		l.length--
		return
	}

	if it.node.next == nil {
		panic("attempt to remove after the last item")
	}
	//TODO: assing node
	removeValueFromNode(it.node.next, 0)
	l.length--
}
