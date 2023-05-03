package main

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s Square) Sides() int {
	return 4
}

func (s Square) Area() int {
	return s.len * s.len
}

func main() {
	var _ Shape = (*Square)(nil)
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}
