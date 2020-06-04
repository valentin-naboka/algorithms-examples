package main

import (
	"fmt"

	"github.com/algorithms-examples/tree-traverse/tree"
)

func makeTree() *tree.Node {
	root := tree.Node{1, nil, nil}
	{
		l := root.AddLeftChild(2)
		{
			l1 := l.AddLeftChild(4)
			{
				l1.AddLeftChild(8)
				l1.AddRightChild(9)
			}
		}
		{
			r1 := l.AddRightChild(5)
			{
				r1.AddLeftChild(10)
				r1.AddRightChild(11)
			}
		}
	}
	{
		r := root.AddRightChild(3)
		{
			r.AddLeftChild(6)
			r2 := r.AddRightChild(7)
			{
				r2.AddRightChild(12)
			}
		}
	}
	return &root
}

func main() {
	root := makeTree()
	visitor := func(v interface{}) {
		value, ok := v.(int)
		if !ok {
			panic("unexpected node value type")
		}
		fmt.Printf("%d ", value)
	}

	tree.TraverseRecursivelyLRN(root, visitor)
	print("\n")
	tree.TraverseRecursivelyNLR(root, visitor)
}
