package linklist

import "testing"

func isPalindrome(head *ListNode) bool {
	tmp := head
	mid := head

	for tmp != nil {
		mid = mid.Next
		if mid == nil {
			// 奇数长度list, false
			return false
		}
		tmp = tmp.Next.Next
	}

	rvmid := reverseLinkList(mid)

	tmp = head
	for rvmid != nil {
		if tmp.Val != rvmid.Val {
			return false
		}
		tmp = tmp.Next
		rvmid = rvmid.Next
	}
	return true
}

func reverseLinkList(head *ListNode) *ListNode {
	tmp := head
	prev := tmp
	next := tmp.Next

	for tmp != nil {
		prev = tmp
		tmp = next
		next = tmp.Next
		tmp.Next = prev
	}
	return tmp
}

func TestisPalindrome(t *testing.T) {
	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	dut := isPalindrome(ll)
	if dut != false {
		t.Fatal("fail")
	}
}

func TestisPalindrome2(t *testing.T) {
	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	dut := isPalindrome(ll)
	if dut != true {
		t.Fatal("fail")
	}
}

func TestisPalindrome3(t *testing.T) {
	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	dut := isPalindrome(ll)
	if dut != false {
		t.Fatal("fail")
	}
}
