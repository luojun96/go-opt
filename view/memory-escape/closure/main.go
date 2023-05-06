package main

import (
	"fmt"
)

func Increase() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// build with: `go build -gcflags="-m -m -l" main.go`
func main() {
	in := Increase()
	fmt.Println(in())
}
