package dp

// https://leetcode.cn/problems/fibonacci-number/

func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib1(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}
