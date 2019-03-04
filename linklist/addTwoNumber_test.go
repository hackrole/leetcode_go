package linklist


import (
  "testing"
)

type Stack struct{
  data []int
  l int
}

func getLen(l *ListNode) int{
  i := 0
  for l != nil {
      i++
      l = l.Next
  }
  return i
}

func makeStackFromList(l *ListNode) *Stack{
  a := make([]int, getLen(l))
  for l != nil {
      a = append(a, l.Val)
      l = l.Next
  }
  s := &Stack{data: a, l: len(a)}
  return s
}


func (s *Stack) pop() int{
  if s.l == 0 {
      return 0
  }
  s.l -= 1
  return s.data[s.l]
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
  a := makeStackFromList(l1)
  b := makeStackFromList(l2)
  l := a.l

  if l < b.l {
      l = b.l
  }

  c := 0

  result := make([]int, l + 1)
  for i := 0; i < l; i++ {
      a0, b0 := a.pop(), b.pop()
      n := (a0 + b0 + c) % 10
      c = (a0 + b0 + c) / 10
      result = append(result, n)
  }
  if c != 0 {
      result = append(result, c)
  }

  rr := &ListNode{Val: result[len(result) - 1]}
  tmp := rr
  for i := len(result) - 2; i >= 0; i-- {
      node := &ListNode{Val: result[i]}
      tmp.Next = node
  }
  return rr
}


func TestSum2(t *testing.T) {
	l1 := makelist([]int{5})
	l2 := makelist([]int{5})
	l := addTwoNumbers2(l1, l2)
	result := list2int(l)
	if result != 10 {
		t.Errorf("test error: l1: %v, l2: %v, result: %d", l1, l2, result)
	}
}
