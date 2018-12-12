// 二叉搜索书两节点最小公共祖先
package tree

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode{
  if p.Val > q.Val {
    p, q = q, p
  }

  for root.Val < p.Val || root.Val > q.Val {
    if root.Val < p.Val {
      root = root.Right
    }
    if root.Val > q.Val {
      root = root.Left
    }
  }
  return root
}
