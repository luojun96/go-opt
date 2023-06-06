package dp

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/?envType=study-plan-v2&envId=dynamic-programming
func maxProfitWithFee(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}

	return dp[n-1][0]
}
