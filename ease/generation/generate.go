package main

import "fmt"

//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func generateUint32Example() {
	var u uint32 = 42
	c := NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateExample: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./template/container.tmp.go gen string container
func generateStringExample() {
	var s string = "Hello"
	container := NewStringContainer()
	container.Put(s)
	v := container.Get()
	fmt.Printf("generateExample: %s (%s)\n", v, v)
}
