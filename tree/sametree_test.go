package tree

import "testing"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return sym(root.Left, root.Right)
}

func sym(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && sym(l.Left, r.Right) && sym(l.Right, r.Left)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val == q.Val && isLeaf(p) && isLeaf(q) {
		return true
	}

	if p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}

func isLeaf(p *TreeNode) bool {
	if p.Left == nil && p.Right == nil {
		return true
	}
	return false
}

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		desc string
		q    *TreeNode
		p    *TreeNode
		dut  bool
	}{
		{
			desc: "left is not right",
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			p:    &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			dut:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := isSameTree(tC.q, tC.p)
			if res != tC.dut {
				t.Fatal("what's possible")
			}
		})
	}
}

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		desc string
		tree *TreeNode
		dut  bool
	}{
		{
			desc: "tree symmetric",
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			dut: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := isSymmetric(tC.tree)
			if res != tC.dut {
				t.Fatal("not impossible")
			}
		})
	}
}
