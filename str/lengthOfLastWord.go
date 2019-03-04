package str

import (
	"testing"
	"strings"
)

func lengthOfLastWord(s string) int {
	a := strings.Split(strings.Trim(s, " "), " ")
	return len(a[len(a)-1])
}

// TestLengthOfLastWord Export
func TestLengthOfLastWord(t *testing.T) {
  a := lengthOfLastWord("a ")
  if a != 1 {
    t.Fatal("error")
  }
}
