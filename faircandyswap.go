package main

import "fmt"
import "sort"

func fairCandySwap(A []int, B []int) []int {
	sort.Ints(A)
	sort.Ints(B)
	s1 := sum(A)
	s2 := sum(B)
	s := s1 - s2
	if s%2 != 0 {
		return nil
	}
	s = s / 2

	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i]-B[j] < s {
			i += 1
			continue
		} else if A[i]-B[j] > s {
			j += 1
			continue
		} else {
			return []int{A[i], B[j]}
		}
	}
	return nil
}

func sum(arr []int) int {
	r := 0
	for i := 0; i < len(arr); i++ {
		r += arr[i]
	}
	return r
}

func main() {
	A := []int{1, 2, 5}
	B := []int{2, 4}

	fmt.Printf("%v\n", fairCandySwap(A, B))
}
