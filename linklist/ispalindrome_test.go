package linklist

import "testing"

func isPalindrome(head *ListNode) bool {
	tmp := head
	mid := head

	for tmp != nil {
		mid = mid.Next
		if tmp.Next == nil {
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
	var prev, next *ListNode
	next = head

	for next != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func TestIsPalindrome(t *testing.T) {
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

func TestIsPalindrome2(t *testing.T) {
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

func TestIsPalindrome3(t *testing.T) {
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

func TestIsPalindrome4(t *testing.T) {
	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
			},
		},
	}

	dut := isPalindrome(ll)

	if dut != false {
		t.Fatal("fail")
	}
}

func TestIsPalindrome5(t *testing.T) {
	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	dut := isPalindrome(ll)

	if dut != false {
		t.Fatal("fail")
	}
}
