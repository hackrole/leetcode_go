package design

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	data     []int
	capacity int
	length   int
	head     int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		capacity: capacity,
		data:     make([]int, capacity),
	}
}

func (q *ArrayQueue) Shift() (int, error) {
	if q.length == 0 {
		return 0, errors.New("empty")
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.length--
	return v, nil
}

type FullError struct {
	len   int
	value int
}

func (e *FullError) Error() string {
	return fmt.Sprintf("len: %d, value: %d full", e.len, e.value)
}

func (q *ArrayQueue) Unshift(value int) error {
	if q.length >= q.capacity {
		return errors.New("array full")
	}
	index := (q.head + q.length) % q.capacity
	q.data[index] = value
	q.length++
	return nil
}
