package greedy

import (
	"slices"
)

// https://leetcode.cn/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	var res int
	slices.Sort(g)
	slices.Sort(s)
	index := len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			res++
			index--
		}
	}
	return res
}
