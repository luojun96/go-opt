package dp

// https://leetcode.cn/problems/coin-change/
// input: coins = [1, 2, 5], amount = 11
// output: 3
func coinChange(coins []int, amount int) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
