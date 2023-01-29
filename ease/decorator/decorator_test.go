package main

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var fn func(s string)
	fn = func(s string) {
		fmt.Println(s)
	}
	decorator(fn)("Hello, World:")
}
