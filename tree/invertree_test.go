// this func InvertTree. for example below
// ```
//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// ```
//
// to
//
// ```
//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1
// ```

package tree

import (
	"testing"
)

func InvertTree(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTree(root.Left)
	InvertTree(root.Right)
}

func TestInvertTree(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	InvertTree(tree)

	if tree.Val != 4 {
		t.Fatal("fail")
	}
	if tree.Right.Val != 2 {
		t.Fatal("fail")
	}
	if tree.Left.Val != 7 {
		t.Fatal("fail")
	}
	if tree.Left.Left.Val != 9 {
		t.Fatal("fail")
	}
	if tree.Left.Right.Val != 6 {
		t.Fatal("fail")
	}
}
