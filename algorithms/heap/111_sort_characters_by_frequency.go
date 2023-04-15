package heap

import "bytes"

// https://leetcode.cn/problems/sort-characters-by-frequency/
// input: s = "tree"
// output: "eert"
func frequencySort(s string) string {
	cnt := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cnt {
		buckets[c] = append(buckets[c], ch)
	}

	res := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			res = append(res, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
