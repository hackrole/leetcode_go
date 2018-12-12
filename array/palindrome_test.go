package array

import (
	"strconv"
	"testing"
)

func isPalindrome2(x int) bool {
	s := strconv.Itoa(x)
	slen := len(s)
	for i := 0; i <= slen/2; i++ {
		if s[i] != s[slen-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	reverse := 0
	for num != 0 {
		reverse = reverse*10 + num%10
		num = num / 10
	}
	if reverse != x {
		return false
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	testcase := map[int]bool{
		121:  true,
		-121: false,
		10:   false,
	}

	for k, v := range testcase {
		dut := isPalindrome(k)
		if dut != v {
			t.Errorf("k: %d, expect: %t, indeed: %t", k, v, dut)
		}
	}
}
