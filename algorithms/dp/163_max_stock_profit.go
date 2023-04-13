package algorithms

// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
// input: [7,1,5,3,6,4]
// output: 5
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}
