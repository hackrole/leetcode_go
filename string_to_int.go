package main

import (
	"fmt"
	"math"
)

func main() {
	b := math.Pow(2, 31)
	b = b*10 + 1

	fmt.Printf("%d\n", -int(math.Pow(2, 31)))
	fmt.Printf("%d\n", b)
}
