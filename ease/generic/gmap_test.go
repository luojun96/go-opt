package main

import (
	"strings"
	"testing"
)

func TestGMap(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	squares := gMap(nums, func(elem int) int {
		return elem * elem
	})
	print(squares)

	strs := []string{"Luo", "Jun"}
	upstrs := gMap(strs, func(s string) string {
		return strings.ToUpper(s)
	})
	print(upstrs)

	dict := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	strs = gMap(nums, func(elem int) string {
		return dict[elem]
	})
	print(strs)
}
