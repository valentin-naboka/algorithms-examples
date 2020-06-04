package tree

import "fmt"

//Post-order
func TraverseRecursivelyLRN(n *Node, visitor func(Value interface{})) {
	if n != nil {
		TraverseRecursivelyLRN(n.Left, visitor)
		TraverseRecursivelyLRN(n.Right, visitor)
		visitor(n.Value)
	}
}

func printTreeRecursive(n *Node, h int) {
	if n != nil {
		fmt.Printf("%*s%d", h*3, "", n.Value.(int))
		fmt.Printf("\n")
		printTreeRecursive(n.Left, h+1)
		printTreeRecursive(n.Right, h+1)
	}
}

func printNR(n *Node) {
	stack := make([]*Node, 0)
	ok := true
	currNode := n
	for ok {
		if currNode != nil {
			if currNode.Right != nil {
				stack = append(stack, currNode.Right)
			}
			if currNode.Left != nil {
				stack = append(stack, currNode.Left)
			}
			// stack = append(stack, currNode)
			println(currNode.Value.(int))
		}
		ok = len(stack) > 0
		if ok {
			currNode = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
	}
}

func printTree(n *Node) {
	queue := make([]*Node, 0)
	ok := true
	currNode := n
	for ok {
		if currNode != nil {
			println(currNode.Value.(int))
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			// stack = append(stack, currNode)
		}
		ok = len(queue) > 0
		if ok {
			currNode = queue[0]
			queue = queue[1:len(queue)]
		}
	}
}
