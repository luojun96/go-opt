package main

import (
	"sort"
)

// https://leetcode.cn/problems/custom-sort-string/solutions/
// input: order = "cba", s = "abcd"
// output: "cbad"

func customSortString(order string, s string) string {
	val := map[byte]int{}

	for i, c := range order {
		val[byte(c)] = i + 1
	}

	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return val[t[i]] < val[t[j]] })
	return string(t)
}
