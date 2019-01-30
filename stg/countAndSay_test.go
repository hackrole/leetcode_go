package stg

// TODO not finish

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	a := "1"
	for i := 1; i <= n; i++ {
		a = cs(a)
	}

	return a

}

func cs(a string) string {
  if a == "1" {
    return "11"
  }

	cur := a[0]
	count := 0
	result := []byte{}

	for i := 0; i < len(a); i++ {
		if cur == a[i] {
			count += 1
			continue
		} else {
			result = append(result, byte(count))
			result = append(result, cur)
			cur = a[i]
			count = 0
		}
  }
  result = append(result, byte(count))
  result = append(result, cur)
	return string(result)
}

func TestCS(t *testing.T) {
	Convey("test countAndSay", t, func() {
		a := countAndSay(4)
		So(a, ShouldEqual, "1211")
	})
}
