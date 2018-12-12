package tree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preiter(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		fmt.Println(node.Val)
	}

	preiter(node.Left)
	preiter(node)
	preiter(node.Right)
}

func TestPreiter(t *testing.T) {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 12,
				},
			},
		},
	}

	preiter(node)
}
