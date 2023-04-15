package string

// https://leetcode.cn/problems/longest-palindromic-substring/
// 动态规划方法
// 状态： dp[i][j]表示子串s[i..j]是否为回文子串
// 状态转移方程：dp[i][j] = s[i] == s[j] && dp[i+1][j-1] == true
// 边界条件：j - 1 - (i + 1) + 1 < 2, 整理得出 j - i < 3 等价于 j - i + 1 < 4 (s[i..j]长度为2或者3时，不用检查是否回文)
func longestPalindromeByDP(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := l + i - 1
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]

}
