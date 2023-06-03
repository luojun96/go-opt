package dp

// https://zh.wikipedia.org/wiki/%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98
// 279. 完全平方数
// https://leetcode-cn.com/problems/perfect-squares/
// 给定正整数 n，找到若干个完全平方数（比如 1,4,9,16,...）使得它们的和等于 n。
// 你需要让组成和的完全平方数的个数最少。
// 示例 1:
// 输入: n = 12
// 输出: 3
// 解释: 12 = 4 + 4 + 4.
// 示例 2:
// 输入: n = 13
// 输出: 2
// 解释: 13 = 4 + 9.
// 解题思路：
// 1. dp[i]表示组成i的完全平方数的最小个数
// 2. dp[i] = min(dp[i], dp[i-j*j]+1)
// 3. dp[0] = 0
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = i
	}
	for i := 2; i < n+1; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
