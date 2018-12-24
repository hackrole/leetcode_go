package lru

import (
	"errors"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type ListNode struct {
	val  int
	key  int
	prev *ListNode
	next *ListNode
}

func (head *ListNode) String() string {
	arr := make([]string, 8)
	for head != nil {
		s := strings.Join([]string{string(head.key), string(head.val)}, "-")
		arr = append(arr, s)
		head = head.next
	}
	return strings.Join(arr, "->")
}

type LRU struct {
	data map[int]*ListNode
	head *ListNode
	tail *ListNode

	capatial int
	length   int
}

// New LRU instance
func New(capatial int) (*LRU, error) {
	if capatial < 0 {
		return nil, errors.New("error")
	}

	obj := &LRU{
		capatial: capatial,
		length:   0,
		head:     &ListNode{},
		tail:     &ListNode{},
		data:     make(map[int]*ListNode, capatial),
	}
	obj.head.next = obj.tail
	obj.tail.prev = obj.head
	return obj, nil
}

// Set value with key in cache
func (c *LRU) Set(key int, value int) {
	node, ok := c.data[key]
	// key exists
	if ok {
		node.val = value
		// move node to tail
		node.prev.next = node.next
		node.next.prev = node.prev
		node.prev = c.tail.prev
		c.tail.prev.next = node
		node.next = c.tail
		c.tail.prev = node
		return
	}
	// not exists and not full
	if c.length != c.capatial {
		node := &ListNode{
			val: value,
			key: key,
		}
		// add to tail
		node.prev = c.tail.prev
		c.tail.prev.next = node
		node.next = c.tail
		c.tail.prev = node

		c.length++
		c.data[key] = node
		return
	}

	// full and not exists
	node = c.head.next
	// remove from header and dict
	c.head.next = node.next
	node.next.prev = c.head
	delete(c.data, node.key)

	node = &ListNode{
		key: key,
		val: value,
	}
	// add to tail
	node.prev = c.tail.prev
	c.tail.prev.next = node
	node.next = c.tail
	c.tail.prev = node
	c.data[key] = node
	return
}

// Get value with key in cache
func (c *LRU) Get(key int) (int, error) {
	node, ok := c.data[key]
	if !ok {
		return -1, errors.New("not exists")
	}

	// move node to tail after access
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = c.tail.prev
	c.tail.prev.next = node
	node.next = c.tail
	c.tail.prev = node

	return node.val, nil
}

func TestLRU(t *testing.T) {
	Convey("test lru ok", t, func() {
		Convey("test New ok", func() {
			lru, _ := New(4)
			So(lru.capatial, ShouldEqual, 4)
			So(lru.length, ShouldEqual, 0)

			Convey("set/get value ok", func() {
				lru.Set(1, 1)
				So(lru.capatial, ShouldEqual, 4)
				So(lru.length, ShouldEqual, 1)
				value, _ := lru.Get(1)
				So(value, ShouldEqual, 1)
				lru.Set(1, 11)
				value, _ = lru.Get(1)
				So(value, ShouldEqual, 11)
				_, error := lru.Get(2)
				So(error, ShouldNotEqual, nil)
			})
			Convey("full set ok", func() {
				value, error := lru.Get(1)
				So(value, ShouldEqual, -1)
				So(error, ShouldNotEqual, nil)
				lru.Set(1, 1)
				lru.Set(2, 2)
				lru.Set(3, 3)
				lru.Set(4, 4)
				So(lru.length, ShouldEqual, 4)
				lru.Set(5, 5)
				So(lru.length, ShouldEqual, 4)
				value, error = lru.Get(1)
				So(value, ShouldEqual, -1)
				So(error, ShouldNotEqual, nil)
				lru.Get(2)
				lru.Set(6, 6)
				value, error = lru.Get(3)
				So(value, ShouldEqual, -1)
				So(error, ShouldNotEqual, nil)
				value, error = lru.Get(2)
				So(value, ShouldEqual, 2)
			})
		})
	})
}
