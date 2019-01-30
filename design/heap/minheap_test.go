// XXX copy from internet, no test
// TODO implement myself.
package heap

import "math"

type HeapNode struct {
	heap  *MinHeap
	index int
}

type MinHeap struct {
	arr []int
}

func NewMinHeap(values ...int) *MinHeap {
	h := &MinHeap{
		arr: append([]int{}, values...),
	}

	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}

func (h *MinHeap) Count() int {
	return len(h.arr)
}

func (h *MinHeap) Push(value int) {
	h.arr = append(h.arr, value)
	h.up(len(h.arr) - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}

	val := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.down(0)

	return val, true
}

func (h *MinHeap) up(idx int) {
	for {
		parentIdx := (idx - 1) / 2
		if idx == 0 || h.arr[parentIdx] <= h.arr[idx] {
			break
		}
		h.arr[idx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[idx]
		idx = parentIdx
	}
}

func (h *MinHeap) down(idx int) {
	for {
		childIdx := idx*2 + 1
		if childIdx >= len(h.arr) || childIdx < 0 {
			break
		}
		rightIdx := childIdx + 1
		if rightIdx < len(h.arr) && h.arr[childIdx] >= h.arr[rightIdx] {
			childIdx = rightIdx
		}

		if h.arr[childIdx] >= h.arr[idx] {
			break
		}
		h.arr[idx], h.arr[childIdx] = h.arr[childIdx], h.arr[idx]
		idx = childIdx
	}
}

func (h *MinHeap) Root() *HeapNode {
	return h.nodeAt(0)
}

func (h *MinHeap) Height() int {
	return int(math.Floor(math.Log2(float64(len(h.arr)))))
}

func (h *MinHeap) nodeAt(index int) *HeapNode {
	if index < 0 || index >= len(h.arr) {
		return nil
	}

	return &HeapNode{heap: h, index: index}
}
