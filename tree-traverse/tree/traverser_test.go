package tree

import (
	"reflect"
	"testing"

	"github.com/algorithms-examples/tree-traverse/testutil"
)

// Input tree visualization
//     1
//     ├── 3
//     │   ├── 7
//     │   │   └── 12
//     │   └── 6
//     └── 2
//         ├── 5
//         │   ├── 11
//         │   └── 10
//         └── 4
//             ├── 9
//             └── 8
func makeTree() *Node {
	root := Node{1, nil, nil}
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

type visitor struct {
	result []interface{}
}

func newvisitor() *visitor {
	return &visitor{make([]interface{}, 0)}
}

func (v *visitor) visit(val interface{}) {
	v.addResult(val)
}

func (v *visitor) addResult(val interface{}) {
	v.result = append(v.result, val)
}

func testTree(t *testing.T, expected []interface{}, traverser func(*Node, Visitor)) {
	root := makeTree()
	visitor := newvisitor()
	traverser(root, visitor)
	if !reflect.DeepEqual(visitor.result, expected) {
		testutil.PrintCaller(2)
		t.Errorf("expected: %v, got: %v\n\n", expected, visitor.result)
	}
}

func TestTraverseRecursevely(t *testing.T) {
	testTree(t,
		[]interface{}{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 7, 12},
		TraverseRecursivelyNLR)

	testTree(t,
		[]interface{}{8, 4, 9, 2, 5, 10, 11, 6, 3, 12, 7},
		TraverseRecursivelyLNR)
}
