package main

import "fmt"

func main() {
	num := 42

	defer func() {
		fmt.Println(num)
	}()

	num = 13
}
