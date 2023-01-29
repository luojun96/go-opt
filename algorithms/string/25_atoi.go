package main

import "math"

// https://leetcode.cn/problems/string-to-integer-atoi/
// input: s = "42"
// output: 42
func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	if i >= n {
		return 0
	}

	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	default:
	}

	for ; i < n; i++ {
		if s[i] < 48 || s[i] > 57 {
			break
		}

		result = result*10 + int(s[i]-'0')
		if sign*result < math.MinInt32 {
			return math.MinInt32
		}
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return sign * result
}
