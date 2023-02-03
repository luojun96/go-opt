package main

import "testing"

func TestGFilterIn(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	odds := gFilterIn(nums, func(elem int) bool {
		return elem%2 == 1
	})
	print(odds)
}
