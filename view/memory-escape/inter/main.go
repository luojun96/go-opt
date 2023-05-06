package main

import "fmt"

// build with `go build -gcflags="-m -m -l" main.go`
func main() {
	str := "Hello, World!"
	fmt.Println(str)
}
