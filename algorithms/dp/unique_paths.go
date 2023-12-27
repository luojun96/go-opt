package dp

// https://leetcode.cn/problems/unique-paths/

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || j == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}
