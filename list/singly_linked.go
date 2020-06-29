package list

type Node struct {
	Next  *Node
	Value interface{}
}

type SignlyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList() *SignlyLinkedList {
	return &SignlyLinkedList{}
}

func NewSinglyLinkedListWithHead(h *Node) *SignlyLinkedList {
	return &SignlyLinkedList{h}
}

func (l *SignlyLinkedList) Head() *Node {
	return l.head
}

func InsertAfter(n *Node, v interface{}) *Node {
	if n == nil {
		n = &Node{nil, v}
		return n
	}
	n.Next = &Node{n.Next, v}
	return n.Next
}
