package main

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// input: n = 7
// output: 21
func numWays(n int) int {
	const mod int = 1e9 + 7
	f1, f2, f3 := 0, 0, 1
	for i := 1; i <= n; i++ {
		f1 = f2
		f2 = f3
		f3 = (f1 + f2) % mod
	}
	return f3
}
