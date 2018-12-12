package list

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("SFA", func() {
		g.It("nil nil", func() {
			a := SFA(nil)
			g.Assert(a).Equal(nil)
		})
		g.It("empt nil", func() {
			a := SFA([]int{})
			g.Assert(a).Equal(nil)
		})
		g.It("arr tree", func() {
			a := SFA([]int{1, 2})
			g.Assert(a.length()).Equal(2)
		})
	})
}
