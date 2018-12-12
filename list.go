package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type Head *Node

func (h Node) String() string {
	return fmt.Sprintf("%#v", h)
}

func main() {
	a := Node{Value: 1, Next: nil}
	fmt.Println(a)
	head := &a
	fmt.Println(head)
	fmt.Printf("%T", a)
}
