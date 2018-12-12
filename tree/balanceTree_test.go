package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func balance(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if abs(lh-rh) > 1 {
		return false
	}
	if !balance(root.Left) {
		return false
	}
	if !balance(root.Right) {
		return false
	}
	return true
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func testBalance(t *testing.T) {
	Convey("test balance tree", t, func() {
	})
}
