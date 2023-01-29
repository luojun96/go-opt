package main

import "fmt"

// 	type Shape interface {
// 		Sides() int
// 		Area() int
// 	}

type Square struct {
	len int
}

func (s Square) Sides() int {
	return 4
}

func main() {
	// var _ Shape = r*Square)(nil)
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}
