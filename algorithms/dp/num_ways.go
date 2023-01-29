package main

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// input: n = 7
// output: 21
const mod int = 1e9 + 7

func numWays(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
