package linklist

import (
	"math"
	"testing"
)

// ListNode list node for linklist type
type ListNode struct {
	Val  int
	Next *ListNode
}

func makelist(arr []int) *ListNode {
	var l, tmp *ListNode
	for _, v := range arr {
		if l == nil {
			tmp = &ListNode{Val: v}
			l = tmp
		} else {
			tmp.Next = &ListNode{Val: v}
			tmp = l.Next
		}
	}
	return l
}

func list2int(l *ListNode) int {
	ll := l
	var result int
	for i := 0; ll != nil; i++ {
		result += ll.Val * int(math.Pow10(i))
		ll = ll.Next
	}
	return result
}

// add_two_number ...
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2, next, v int
	var l, tmp *ListNode
	for (l1 != nil || l2 != nil) || next != 0 {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		v = (v1 + v2 + next) % 10
		next = (v1 + v2 + next) / 10
		if tmp == nil {
			tmp = &ListNode{Val: v}
			l = tmp
		} else {
			tmp.Next = &ListNode{Val: v}
			tmp = tmp.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return l
}

func TestInt(t *testing.T) {
	l1 := makelist([]int{1, 2, 3})
	if l1.Val != 1 {
		t.Errorf("l1 error: %+v", l1)
	}
	result := list2int(l1)
	if result != 321 {
		t.Errorf("result error: %+v", result)
	}
}

func TestSum(t *testing.T) {
	l1 := makelist([]int{5})
	l2 := makelist([]int{5})
	l := addTwoNumbers(l1, l2)
	result := list2int(l)
	if result != 10 {
		t.Errorf("test error: l1: %v, l2: %v, result: %d", l1, l2, result)
	}
}
