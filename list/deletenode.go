package list

//type ListNode struct {
//  Val  int
//  Next *ListNode
//}

func deleteNode(ListNode node) {
	if node != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}

func removeElements(root *ListNode, val int) {
	pre := &ListNode{
		Next: root,
	}
	cur := root
	for cur != nil {
		if cur.Val == val {
			pre.next = now.next
			now = pre.next
		} else {
			pre = pre.next
			now = now.next
		}
	}
	return pre.Next
}
