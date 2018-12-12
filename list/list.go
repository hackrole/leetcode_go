package list

type SingleList struct {
	Val  int
	Next *SingleList
}

type DoubleList struct {
	Val  int
	Prev *DoubleList
	Next *DoubleList
}

func SFA(a []int) *SingleList {
	if a == nil || len(a) == 0 {
		return nil
	}

	head := &SingleList{Val: a[0]}
	tmp := head
	for i := 1; i < len(a); i++ {
		node := &SingleList{Val: a[i]}
		tmp.Next = node
		tmp = node
	}
	return head
}

func (s *SingleList) length() int {
	if s == nil {
		return 0
	}
	a := 0
	tmp := s
	for tmp != nil {
		a += 1
		tmp = tmp.Next
	}
	return a
}
