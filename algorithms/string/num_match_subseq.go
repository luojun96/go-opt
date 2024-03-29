package string

import "sort"

// https://leetcode.cn/problems/number-of-matching-subsequences/description/
// input: s = "abcde", words = ["a","bb","acd","ace"]
// output: 3
func numMatchingSubseq(s string, words []string) int {
	pos := [26][]int{}

	for i, c := range s {
		pos[c-'a'] = append(pos[c-'a'], i)
	}

	ans := len(words)
	for _, w := range words {
		if len(w) > len(s) {
			ans--
			continue
		}

		p := -1
		for _, c := range w {
			ps := pos[c-'a']
			j := sort.SearchInts(ps, p+1)
			if j == len(ps) {
				ans--
				break
			}
			p = ps[j]
		}
	}
	return ans
}
