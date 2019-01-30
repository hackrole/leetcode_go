package tree

import (
	"github.com/smartstreets/goconvey/convey"
	"testing"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return
	}
	result := make([]int)

}

func btp(root *TreeNode, path, string, s []string) {
	if root == nil {
		return s
	}

	path += s.Val

	if root.Left == nil && root.Right == nil {
		s = append(s, path)
	}

	btp(root.Left, path, s)
	btp(root.Right, path, s)

	return s
}
