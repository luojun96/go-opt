package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("before filter, set: %v\n", intset)
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("after filter, set: %v\n", out)
	var expect = []int{1, 3, 5, 7, 9}
	res := reflect.DeepEqual(expect, out)
	if !res {
		t.Errorf("expecting %v, got %v\n", expect, out)
	}
}
