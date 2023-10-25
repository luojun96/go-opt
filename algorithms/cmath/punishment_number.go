package cmath

import "strconv"

func punishmentNumber(n int) int {
	var dfs func(string, int, int, int) bool
	dfs = func(s string, pos, tot, target int) bool {
		if pos == len(s) {
			return tot == target
		}
		sum := 0
		for i := pos; i < len(s); i++ {
			sum = sum*10 + int(s[i]-'0')
			if sum+tot > target {
				break
			}
			if dfs(s, i+1, sum+tot, target) {
				return true
			}
		}
		return false
	}

	res := 0
	for i := 1; i <= n; i++ {
		if dfs(strconv.Itoa(i*i), 0, 0, i) {
			res += i * i
		}
	}
	return res
}
