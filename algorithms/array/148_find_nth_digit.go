package algorithms

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
// input: n = 3
// output: 3
func findNthDigit(n int) int {
	var totalDigits func(length int) int
	totalDigits = func(length int) int {
		var digits int
		for curLength, curCount := 1, 9; curLength <= length; curLength++ {
			digits += curLength * curCount
			curCount *= 10
		}
		return digits
	}

	d := 1 + sort.Search(8, func(length int) bool {
		return totalDigits(length+1) >= n
	})
	prevDigits := totalDigits(d - 1)
	index := n - prevDigits - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}
