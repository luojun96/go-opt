package dp

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/
// input: s = "00110"
// output: 1
func minFlipsMonoIncr(s string) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	dp0, dp1 := 0, 0
	for _, c := range s {
		dp0New, dp1New := dp0, min(dp0, dp1)
		if c == '1' {
			dp0New++
		} else {
			dp1New++
		}
		dp0, dp1 = dp0New, dp1New
	}
	return min(dp0, dp1)
}
