package design

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Stack struct {
	capacity int
	len      int
	data     []int
}

func New(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
		len:      0,
		data:     make([]int, capacity),
	}
}

func NewFromArray(arr []int) *Stack {
	cap := cap(arr)
	stack := New(cap)
	for _, v := range arr {
		stack.push(v)
	}
	return stack
}

func (s *Stack) pop() (int, error) {
	if s.len == 0 {
		return -1, errors.New("empty")
	}
	v := s.data[s.len-1]
	s.len--
	return v, nil
}

func (s *Stack) push(v int) error {
	if s.len == s.capacity {
		return errors.New("full")
	}
	s.data[s.len] = v
	s.len++
	return nil
}

func (s *Stack) String() string {
	return fmt.Sprintf("Stack: %v", s.data[:s.len])
}

func TestStack(t *testing.T) {
	Convey("test array stack works", t, func() {
		Convey("test new ok", func() {
			stack := New(2)
			So(stack.capacity, ShouldEqual, 2)
			So(stack.len, ShouldEqual, 0)
			So(len(stack.data), ShouldEqual, 0)
		})
	})
}
