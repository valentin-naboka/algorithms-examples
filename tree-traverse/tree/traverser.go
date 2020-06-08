package tree

import "github.com/algorithms-examples/tree-traverse/stack"

type visitor func(Value interface{})

// TraverseRecursivelyLRN is a post-order traverse.
// Traverse the left subtree by recursively calling the post-order function.
// Traverse the right subtree by recursively calling the post-order function.
// Access the data part of the current node.
func TraverseRecursivelyLRN(n *Node, visitor visitor) {
	if n != nil {
		TraverseRecursivelyLRN(n.Left, visitor)
		TraverseRecursivelyLRN(n.Right, visitor)
		visitor(n.Value)
	}
}

// TraverseRecursivelyNLR is a pre-order traverse.
// Access the data part of the current node.
// Traverse the left subtree by recursively calling the pre-order function.
// Traverse the right subtree by recursively calling the pre-order function.
// The pre-order traversal is a topologically sorted one, because a parent node is processed before any of its child nodes is done.
func TraverseRecursivelyNLR(n *Node, visitor visitor) {
	if n != nil {
		visitor(n.Value)
		TraverseRecursivelyNLR(n.Left, visitor)
		TraverseRecursivelyNLR(n.Right, visitor)
	}
}

// TraverseRecursivelyLNR is a in-order traverse.
// Traverse the left subtree by recursively calling the in-order function.
// Access the data part of the current node.
// Traverse the right subtree by recursively calling the in-order function.
// In a binary search tree ordered such that in each node the key is greater than all keys in its left subtree and less than all keys in its right subtree,
// in-order traversal retrieves the keys in ascending sorted order.
func TraverseRecursivelyLNR(n *Node, visitor visitor) {
	if n != nil {
		TraverseRecursivelyNLR(n.Left, visitor)
		visitor(n.Value)
		TraverseRecursivelyNLR(n.Right, visitor)
	}
}

func toNode(v interface{}) *Node {
	node, ok := v.(*Node)
	if !ok {
		panic("value is not *Node")
	}
	return node
}

//TODO: add comment
func TraverseLNR(n *Node, visitor visitor) {
	stack := stack.Stack{}
	notEmpty := true
	currNode := n
	for notEmpty {
		if currNode != nil {
			if currNode.Right != nil {
				stack.Push(currNode.Right)
			}
			if currNode.Left != nil {
				stack.Push(currNode.Left)
			}
			visitor(currNode.Value.(int))
		}
		notEmpty = !stack.IsEmpty()
		if notEmpty {
			currNode = toNode(stack.Pop())
		}
	}
}

//TODO: implement
// func printTree(n *Node) {
// 	queue := make([]*Node, 0)
// 	ok := true
// 	currNode := n
// 	for ok {
// 		if currNode != nil {
// 			println(currNode.Value.(int))
// 			if currNode.Right != nil {
// 				queue = append(queue, currNode.Right)
// 			}
// 			if currNode.Left != nil {
// 				queue = append(queue, currNode.Left)
// 			}
// 			// stack = append(stack, currNode)
// 		}
// 		ok = len(queue) > 0
// 		if ok {
// 			currNode = queue[0]
// 			queue = queue[1:len(queue)]
// 		}
// 	}
// }
