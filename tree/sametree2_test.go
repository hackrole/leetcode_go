package tree

func isSameTree(p *NodeTree, q *NodeTree) bool {
  if p == nil && q == nil {
    return true
  }
  if (p != nil && q == nil) || (p == nil && q != nil){
    return false
  }

  if p.Val != q.Val {
    return False
  }

  return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}