package main

// https://leetcode.cn/problems/count-the-number-of-consistent-strings/solutions/

func countConsistentStrings(allowed string, words []string) int {
	res := 0
	mask := 0
	for _, c := range allowed {
		mask |= 1 << (c - 'a')
	}

	for _, word := range words {
		mask1 := 0
		for _, c = range word {
			mask1 |= 1 << (c - 'a')
		}
		if mask1|mask == mask {
			res++
		}
	}
	return res
}
