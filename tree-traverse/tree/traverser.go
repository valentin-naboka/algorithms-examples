package tree

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
// Pre-order (NLR)
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

// func printNR(n *Node) {
// 	stack := make([]*Node, 0)
// 	ok := true
// 	currNode := n
// 	for ok {
// 		if currNode != nil {
// 			if currNode.Right != nil {
// 				stack = append(stack, currNode.Right)
// 			}
// 			if currNode.Left != nil {
// 				stack = append(stack, currNode.Left)
// 			}
// 			// stack = append(stack, currNode)
// 			println(currNode.Value.(int))
// 		}
// 		ok = len(stack) > 0
// 		if ok {
// 			currNode = stack[len(stack)-1]
// 			stack = stack[0 : len(stack)-1]
// 		}
// 	}
// }

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
