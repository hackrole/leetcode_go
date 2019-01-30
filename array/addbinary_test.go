package array

// import (
// 	"strconv"
// )

// // addBinary ...
// func addBinary(a string, b string) string {
// 	r := []uint{}

// 	l1 := len(a)
// 	l2 := len(b)
// 	ll := min(l1, l2)
// 	l := max(l1, l2)
// 	d := 0
// 	for i := 0; i <= l-1; i++ {
// 		if l1 <= i {
// 			ai := 0
// 		} else {
// 			ai := int(a[l1-1-i]) - '0'
// 		}
// 		if l2 <= i {
// 			bi := 0
// 		} else {
// 			bi := int(b[l2-1-i]) - '0'
// 		}
// 		c := ai + bi + d
// 		if c >= 2 {
// 			c = c - 2
// 			d = 1
// 		} else {
// 			d = 0
// 		}
// 		r := append(r, strconv.Itoa(c))
// 	}

// 	for i := 0; i <= len(r)/2; i++ {
// 		r[i], r[len(r)-i] = r[len(r)-1], r[i]
// 	}
// 	return string(r)
// }

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
