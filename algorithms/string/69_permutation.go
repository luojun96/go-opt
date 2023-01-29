package main

import "sort"

// https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/
// input: s = "abc"
// output: ["abc","acb","bac","bca","cab","cba"]
func permutation(s string) []string {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var res []string
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			res = append(res, string(perm))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return res
}
