package tree

import (
	"github.com/algorithms-examples/tree-traverse/queue"
	"github.com/algorithms-examples/tree-traverse/stack"
)

//Visitor interface to visit nodes while traversing the tree.
type Visitor interface {
	visit(Value interface{})
}

// TraverseRecursivelyNLR is a pre-order traverse.
// Access the data part of the current node.
// Traverse the left subtree by recursively calling the pre-order function.
// Traverse the right subtree by recursively calling the pre-order function.
// The pre-order traversal is a topologically sorted one, because a parent node is processed before any of its child nodes is done.
func TraverseRecursivelyNLR(n *Node, visitor Visitor) {
	if n != nil {
		visitor.visit(n.Value)
		TraverseRecursivelyNLR(n.Left, visitor)
		TraverseRecursivelyNLR(n.Right, visitor)
	}
}

// TraverseRecursivelyLRN is a post-order traverse.
// Traverse the left subtree by recursively calling the post-order function.
// Traverse the right subtree by recursively calling the post-order function.
// Access the data part of the current node.
func TraverseRecursivelyLRN(n *Node, visitor Visitor) {
	if n != nil {
		TraverseRecursivelyLRN(n.Left, visitor)
		TraverseRecursivelyLRN(n.Right, visitor)
		visitor.visit(n.Value)
	}
}

// TraverseRecursivelyLNR is a in-order traverse.
// Traverse the left subtree by recursively calling the in-order function.
// Access the data part of the current node.
// Traverse the right subtree by recursively calling the in-order function.
// In a binary search tree ordered such that in each node the key is greater than all keys in its left subtree and less than all keys in its right subtree,
// in-order traversal retrieves the keys in ascending sorted order.
func TraverseRecursivelyLNR(n *Node, visitor Visitor) {
	if n != nil {
		TraverseRecursivelyLNR(n.Left, visitor)
		visitor.visit(n.Value)
		TraverseRecursivelyLNR(n.Right, visitor)
	}
}

func toNode(v interface{}) *Node {
	node, ok := v.(*Node)
	if !ok {
		panic("value is not *Node")
	}
	return node
}

// TraverseNLR is a pre-order traverse using stack instead of recursion.
func TraverseNLR(n *Node, visitor Visitor) {
	stack := stack.Stack{}
	currNode := n
	for currNode != nil {
		if currNode.Right != nil {
			stack.Push(currNode.Right)
		}
		if currNode.Left != nil {
			stack.Push(currNode.Left)
		}
		visitor.visit(currNode.Value.(int))

		if !stack.IsEmpty() {
			currNode = toNode(stack.Pop())
		} else {
			currNode = nil
		}
	}
}

func toNodeWrapper(v interface{}) *nodeWrapper {
	n, ok := v.(*nodeWrapper)
	if !ok {
		panic("value is not *Node")
	}
	return n
}

type nodeWrapper struct {
	n     *Node
	ready bool
}

func traverse(n *Node, visitor Visitor, nodeInserter func(currNodeWrap *nodeWrapper, stack *stack.Stack)) {
	stack := stack.Stack{}
	stack.Push(&nodeWrapper{n, false})
	for !stack.IsEmpty() {
		currNodeWrap := toNodeWrapper(stack.Pop())
		if currNodeWrap.ready || currNodeWrap.n.Left == nil && currNodeWrap.n.Right == nil {
			visitor.visit(currNodeWrap.n.Value.(int))
		} else {
			nodeInserter(currNodeWrap, &stack)
		}
	}
}

// TraverseLRN is a post-order traverse using stack instead of recursion.
func TraverseLRN(n *Node, visitor Visitor) {
	traverse(n, visitor, func(currNodeWrap *nodeWrapper, stack *stack.Stack) {
		currNodeWrap.ready = true
		stack.Push(currNodeWrap)
		if currNodeWrap.n.Right != nil {
			stack.Push(&nodeWrapper{currNodeWrap.n.Right, false})
		}

		if currNodeWrap.n.Left != nil {
			stack.Push(&nodeWrapper{currNodeWrap.n.Left, false})
		}
	})
}

// TraverseLNR is a in-order traverse using stack instead of recursion.
func TraverseLNR(n *Node, visitor Visitor) {
	traverse(n, visitor, func(currNodeWrap *nodeWrapper, stack *stack.Stack) {
		if currNodeWrap.n.Right != nil {
			stack.Push(&nodeWrapper{currNodeWrap.n.Right, false})
		}

		currNodeWrap.ready = true
		stack.Push(currNodeWrap)

		if currNodeWrap.n.Left != nil {
			stack.Push(&nodeWrapper{currNodeWrap.n.Left, false})
		}
	})
}

//TraverseBFS is a breadth-first search, which visits every node on a level before going to a lower level.
func TraverseBFS(n *Node, visitor Visitor) {
	queue := queue.Queue{}
	currNode := n
	for currNode != nil {
		if currNode.Left != nil {
			queue.Push(currNode.Left)
		}
		if currNode.Right != nil {
			queue.Push(currNode.Right)
		}
		visitor.visit(currNode.Value.(int))

		if !queue.IsEmpty() {
			currNode = toNode(queue.Pop())
		} else {
			currNode = nil
		}
	}
}
