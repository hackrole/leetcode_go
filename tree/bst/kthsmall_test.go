// TODO not finish, the array not modify in place
package bst

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bst2Arr(root *TreeNode, arr []int) {
	if root == nil {
		return
	}

	bst2Arr(root.Left, arr)
	arr = append(arr, root.Val)
	bst2Arr(root.Right, arr)
}

func kthSmallest(root *TreeNode, k int) int {
	var arr []int

	bst2Arr(root, arr)
	return arr[k-1]
}

func TestKthSmallest(t *testing.T) {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 2}
	n1.Left = n2
	n1.Right = n3
	n2.Right = n4

	dut := kthSmallest(n1, 1)
	if dut != 1 {
		t.Fatal("failed")
	}
}
