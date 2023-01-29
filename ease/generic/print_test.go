package main

import (
	"testing"
)

func TestGenericPrint(t *testing.T) {
	strs := []string{"Hello", "World", "Generic"}
	decs := []float64{3.14, 1.14, 1.618}
	nums := []int{2, 4, 6, 8}

	print(strs)
	print(decs)
	print(nums)
}
