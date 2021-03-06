package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// TestSpec ...
func TestSpec(t *testing.T) {
	Convey("Given some integer with a starting  value", t, func() {
		x := 1
		Convey("when the integer is incremented", func() {
			x++
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
