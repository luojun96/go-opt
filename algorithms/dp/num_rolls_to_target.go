package dp

func numRollsToTarget(n int, k int, target int) int {
	mod := int(1e9 + 7)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, target+1)
	}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			for x := 1; x <= k; x++ {
				if j-x >= 0 {
					f[i][j] = (f[i][j] + f[i-1][j-x]) % mod
				}
			}
		}
	}
	return f[n][target]
}
