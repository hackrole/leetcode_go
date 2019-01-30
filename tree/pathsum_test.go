package tree

// Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
// Note: A leaf is a node with no children.

import (
  "testing"
  "github.com/smartystreets/goconvey/convey"
)

func hasPathSum(root *TreeNode, sum int) bool{
  if root == nil {
    return false
  }

  if root.Left == nil && root.Right == nil {
    return root.Val == sum
  }

  return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}


func TestHasPathSum(t *testing.T){
  Convey("test recursive has path sum right", t, func(){
    tree := &TreeNode{
      Val: 5,
      Left: &TreeNode{
        Val: 4,
      },
      Right: &TreeNode{
        Val: 8,
        Left: &TreeNode{
          Val: 13,
        },
        Right: &TreeNode{
          Val: 4,
        },
      },
    }
    So(hasPathSum(tree, 17), ShouldEqual, true)
    So(hasPathSum(tree, 18), shoulEqual, false)
  })
}
