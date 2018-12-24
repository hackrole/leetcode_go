package array

import (
	"fmt"
	"testing"
)

// plugone ...
func plugone(digists []int) []int {
	l := len(digists)
	digists[l-1] = digists[l-1] + 1
	if digists[l-1] < 10 {
		return digists
	}

	digists[l-1] = 0
	for i := l - 2; i >= 0; i-- {
		digists[i] = digists[i] + 1
		if digists[i] < 10 {
			return digists
		}
		digists[i] = 0
	}

	a := make([]int, l+1)
	a[0] = 1
	for i := 1; i <= l; i++ {
		a[i] = 0
	}
	return a
}

func TestPlugone(t *testing.T) {
	a := []int{1, 2, 3}
	b := plugone(a)
	if b[2] != 4 {
		t.Fatal("error")
	}
	a = []int{1, 2, 9}
	b = plugone(a)
	if b[2] != 0 || b[1] != 3 {
		fmt.Println(b)
		t.Fatal("error")
	}
	a = []int{9, 9, 9}
	b = plugone(a)
	if b[2] != 0 || b[1] != 0 || len(b) != 4 {
		t.Fatal("error")
	}
}
