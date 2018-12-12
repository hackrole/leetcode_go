package tree

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	return result
}

func treeToArr(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	return result
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
