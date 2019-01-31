package btree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	row := len(matrix)
	column := len(matrix[0])

	low, high := 0, row * column - 1

	for low <= high {
		mid := (high + low) / 2
		num := matrix[mid/column][mid%column]
		if num == target {
			return true
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func TestSearchMatrix(t *testing.T) {

	Convey("test search", t, func() {
		data := [][]int{
			[]int{1, 3, 5, 7},
			[]int{10, 11, 16, 20},
			[]int{23, 30, 34, 50},
		}
    dut := searchMatrix(data, 3)
		So(dut, ShouldEqual, true)
	})
}
