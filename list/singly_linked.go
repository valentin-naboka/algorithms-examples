package list

type Node struct {
	next  *Node
	value interface{}
}

func (n *Node) MoveNext() *Node {
	if n == nil {
		return nil
	}

	return n.next
}

func (n *Node) GetValue() interface{} {
	return n.value
}

type SignlyLinkedList struct {
	head   *Node
	lenght int
}

func NewSinglyLinkedList(values ...interface{}) *SignlyLinkedList {
	lenght := len(values)
	if lenght == 0 {
		return &SignlyLinkedList{nil, 0}
	}

	head := &Node{nil, values[0]}
	current := head
	for _, v := range values[1:lenght] {
		current.next = &Node{nil, v}
		current = current.MoveNext()
	}

	return &SignlyLinkedList{head, lenght}
}

func (l *SignlyLinkedList) GetHead() *Node {
	return l.head
}

func (l SignlyLinkedList) GetLength() int {
	return l.lenght
}

func (l *SignlyLinkedList) InsertAfter(n *Node, v interface{}) *Node {
	if n == nil {
		l.head = &Node{nil, v}
		return l.head
	}
	n.next = &Node{n.next, v}
	l.lenght++
	return n.next
}

func (l *SignlyLinkedList) PushFront(v interface{}) {
	newNode := &Node{nil, v}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.lenght++
}

func (l *SignlyLinkedList) RemoveAfter(n *Node) {
	if l.lenght == 1 && n == nil {
		n = nil
	} else {
		n.next = n.next.next
	}
	l.lenght--
}
