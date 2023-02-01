package main

import (
	"fmt"
	"testing"
)

func TestGReduce(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := gReduce(nums, 0, func(result, elem int) int {
		return result + elem
	})
	fmt.Printf("Sum = %d \n", sum)
}
