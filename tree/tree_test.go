package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treefromarr(arr []int, root *TreeNode, i int) *TreeNode {
	l := len(arr)
	for ; i <= (l-1)/2; i++ {
		node := &TreeNode{
			Val: arr[i],
		}
		root = node
		root.Left = treefromarr(arr, root.Left, 2*i+1)
		root.Right = treefromarr(arr, root.Right, 2*i+2)
	}

	return root
}

func treefromarr_test(t *testing.T) {
	Convey("test tree from arr", t, func() {
		Convey("test normal", func() {
			arr := []int{0, 1, 2, 3, 4}
			dut := treefromarr(arr, nil, 0)
			So(dut.Val, ShouldEqual, 0)
			So(dut.Left.Val, ShouldEqual, 1)
			So(dut.Right.Val, ShouldEqual, 2)
			So(dut.Left.Left.Val, ShouldEqual, 3)
			So(dut.Left.Right.Val, ShouldEqual, 4)
		})
	})
}
