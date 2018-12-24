package design

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

func New() *ListNode {
	return new(ListNode)
}

func NewFromArray(arr []int) *ListNode {
	head := New()
	head.Val = arr[0]

	tmp := head
	for i := 1; i < len(arr); i++ {
		nd := &ListNode{
			Val:  arr[i],
			Prev: tmp,
		}
		tmp.Next = nd
		tmp = nd
	}
	tmp.Next = head
	head.Prev = tmp
	return head
}

func (head *ListNode) Append(val int) {
	tail := head.Prev

	nd := &ListNode{
		Val:  val,
		Prev: tail,
		Next: head,
	}
	tail.Next = nd
	head.Prev = nd
}

func (head *ListNode) Insert(val int, index int) *ListNode {
	nd := &ListNode{
		Val: val,
	}

	if index == 0 {
		nd.Next = head
		nd.Prev = head.Prev
		head.Prev = nd
		return nd
	}

	tmp := head
	for index != 0 {
		tmp = tmp.Next
		index--
	}
	nd.Next = tmp
	nd.Prev = tmp.Prev
	tmp.Prev = nd
	return head
}

func (head *ListNode) Length() int {
	tmp := head
}

func (head *ListNode) ToArray() []int {
	arr := make([]int, 8)

	tmp := head

	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
}

func (head *ListNode) String() string {

}
