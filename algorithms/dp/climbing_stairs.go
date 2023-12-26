package dp

// https://leetcode.cn/problems/climbing-stairs/

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}

	a := make([]int, 2)
	a[0] = 1
	a[1] = 2
	for i := 3; i <= n; i++ {
		sum := a[0] + a[1]
		a[0] = a[1]
		a[1] = sum
	}
	return a[1]
}
