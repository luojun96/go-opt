package main

import "fmt"

type stack[T any] []T

func (s *stack[T]) push(elem T) {
	*s = append(*s, elem)
}

func (s *stack[T]) pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *stack[T]) top() *T {
	if len(*s) > 0 {
		return &(*s)[len(*s)-1]
	}
	return nil
}

func (s *stack[T]) len() int {
	return len(*s)
}

func (s *stack[T]) print() {
	for _, elem := range *s {
		fmt.Print(elem)
		fmt.Print(" ")
	}
	fmt.Println("")
}
