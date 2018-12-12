package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			fmt.Println(k)
			return k
		}
	}
	fmt.Println(len(nums))
	return len(nums)
}

func main() {
	arr := []int{1, 6, 8, 40}
	searchInsert(arr, 1)
	searchInsert(arr, 2)
	searchInsert(arr, 8)
	searchInsert(arr, 7)
	searchInsert(arr, 57)
}
