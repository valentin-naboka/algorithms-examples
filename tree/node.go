package tree

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

func (n *Node) AddLeftChild(Value interface{}) *Node {
	n.Left = &Node{Value: Value}
	return n.Left
}

func (n *Node) AddRightChild(Value interface{}) *Node {
	n.Right = &Node{Value: Value}
	return n.Right
}
