package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {
	f1 := 3.14
	f2 := 1.14

	c := NewRefContainer(reflect.TypeOf(f1), 16)

	if err := c.Put(f1); err != nil {
		panic(err)
	}

	if err := c.Put(f2); err != nil {
		panic(err)
	}

	g := 0.0

	if err := c.Get(&g); err != nil {
		panic(err)
	}

	fmt.Printf("%v (%T)\n", g, g)
	fmt.Println(c.s.Index(0))
}
