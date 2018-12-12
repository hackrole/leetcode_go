package main

import (
	"fmt"
	"strconv"
)

func rev(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func reverse(x int) int {
	ss := strconv.Itoa(x)
	s := []byte(ss)
	if s[0] == '-' {
		rev(s[1:])
	} else {
		rev(s[0:])
	}
	tmp, _ := strconv.Atoi(string(s))
	return tmp
}

func main() {
	a := 123
	fmt.Println(reverse(a))
	a = -123
	fmt.Println(reverse(a))
	a = 10
	fmt.Println(reverse(a))
}
