package algorithms

// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/
// input: n = 5
// output: 5
func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
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
