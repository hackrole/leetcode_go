package searchtree

type SearchTreeNode struct {
	Val   int
	Left  *SearchTreeNode
	Right *SearchTreeNode
}

func NewFromArray(arr []int) *SearchTreeNode {
	root := &SearchTreeNode{
		Val: arr[0],
	}

	for i := 1; i < len(arr); i++ {
		root.Insert(arr[i])
	}
	return root
}

func (root *SearchTreeNode) Insert(val int) {
	node := &SearchTreeNode{
		Val: val,
	}
	if root == nil {
		root = node
	}

	tmp := root
	for tmp != nil {
		if val <= tmp.Val {
			if tmp.Left == nil {
				tmp.Left = node
				break
			} else {
				tmp = tmp.Left
				continue
			}
		} else if val > tmp.Val {
			if tmp.Right == nil {
				tmp.Right = node
				break
			} else {
				tmp = tmp.Right
				continue
			}
		}
	}
}

func (root *SearchTreeNode) Search(val int) bool {
	if root == nil {
		return false
	}

	tmp := root
	for tmp != nil {
		if val == tmp.Val {
			return true
		}
		if val > tmp.Val {
			tmp = tmp.Right
			continue
		} else {
			tmp = tmp.Left
			continue
		}
	}
	return false
}
