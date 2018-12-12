package tree

func maxDepth(root *NodeTree) int {
	if root == nil {
		return 0
	}

	ldp := maxDepth(root.Left)
	rdp := maxDepth(root.Right)

	if ldp >= rdp {
		return ldp + 1
	} else {
		return rdp + 1
	}
}
