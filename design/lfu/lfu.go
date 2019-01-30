package lfu

import (
	"errors"
)

// DoubleListNode douible-list-node
type doubleListNode struct {
	Val  int
	Prev *doubleListNode
	Next *doubleListNode
	head *lfuLinkNode
}

// LFULinkNode 频率和时间相关linklist
type lfuLinkNode struct {
	Val   int
	Prev  *lfuLinkNode
  Next  *lfuLinkNode
  NodeHead *doubleListNode
  NodeTail *doubleListNode
}

type lfu struct {
	data     map[int]*doubleListNode
	linkhead *lfuLinkNode
	linktail *lfuLinkNode
	capital  int
	length   int
}

func NewLFU(capital int) *lfu {
	head := &lfuLinkNode{}
	tail := &lfuLinkNode{}
	head.Next = tail
	tail.Prev = head

	return &lfu{
		capital:  capital,
		data:     make(map[int]*doubleListNode, capital),
		linkhead: head,
		linktail: tail,
	}
}

func (l *lfu) Get(key int) (int, error) {
	node, ok := l.data[key]
	if !ok {
		return -1, errors.New("not found")
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	head := node.head
	if head.Next.Val-head.Val != 1 {
		tmp := &lfuLinkNode{
			Val:  head.Val + 1,
			Prev: head,
			Next: head.Next,
		}
		head.Next.Prev = tmp
    head.Next = tmp
    head.NodeHead := &doubleListNode{
      Next: node,
      Prev: node,
    }

	}else{
    head.
  }
}

func (l *lfu) Set(key int, value int) {

}
