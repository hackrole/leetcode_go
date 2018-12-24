package linklist

import "testing"

func isMatch(s string, p string) bool {
	return false
}

// TestIsMatch test match
func TestIsMatch(t *testing.T) {
	testcase := map[[2]string]bool{
		[2]string{"aa", "a"}:  false,
		[2]string{"aa", "a*"}: true,
		[2]string{"ab", ".*"}: true,
	}
	for k, v := range testcase {
		s, p := k[0], k[1]
		dut := isMatch(s, p)
		if dut != v {
			t.Errorf("s: %s, p: %s, dut: %t, v: %t", s, p, dut, v)
		}
	}
}
