package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	ss := stack[string]{}
	ss.push("hello")
	ss.push("jun")
	ss.push("luo")
	fmt.Printf("stack top is - %v\n", *(ss.top()))
	ss.pop()
	ss.pop()
	ss.print()

	ns := stack[int]{}
	ns.push(10)
	ns.push(20)
	ns.print()
	ns.pop()
	ns.print()
	*ns.top() += 1
	ns.print()
	ns.pop()
	fmt.Printf("stack top is - %v\n", ns.top())
}
