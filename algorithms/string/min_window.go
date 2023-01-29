package main

import "math"

func minWindow(s string,  t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	maxLen := math.MaxInt
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}

		for check() && l <= r {
			if r - l + 1 < maxLen {
				maxLen = r - l + 1
				ansL, ansR = 1, l + maxLen
			}	

			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}
