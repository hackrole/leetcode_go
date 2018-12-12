package main

import "fmt"

type obj int

type op *obj

func (o obj) String() string {
	return fmt.Sprintf("hello")
}

func (o op) String() string {
	return fmt.Sprintf("world")
}

func main() {
	a := obj(8)
	fmt.Println(a)
	fmt.Println(&a)
}
