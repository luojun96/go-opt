package dp

func cuttingRope(n int) int {
	const mod int = 1e9 + 7
	dp := make([]int, n+1)

	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax % mod
	}
	return dp[n]
}
