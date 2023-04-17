package string

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/
// input: s = "00110"
// output: 1 (00111)
// find the postion of the first one, then use backtracking to find the min flips count
var nFlips int

func minFlipsMonoIncr(s string) int {
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

func minFlips(s string) int {
	res := 0
	foIndex := 0
	for ; foIndex < len(s); foIndex++ {
		if s[foIndex] == '1' {
			break
		}
	}
	if foIndex == len(s) || foIndex == len(s)-1 {
		return 0
	}

	nZero := 0
	for i := foIndex + 1; i < len(s); i++ {
		if s[i] == '0' {
			nZero++
		}
	}

	newStr := s[foIndex+1:]
	count := minFlips(newStr)
	if nZero <= count {
		res = nZero
	} else {
		res = count + 1
	}

	return res
}
