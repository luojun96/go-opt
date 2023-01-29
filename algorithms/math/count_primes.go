package main

// https://leetcode.cn/problems/count-primes/
// input: n = 10,
// output: 4, [2, 3, 5, 7]
func countPrimes(n int) int {
	isPrim := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrim[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrim[i] {
			for j := i * i; j < n; j += i {
				isPrim[j] = false
			}
		}
	}
	res := 0
	for i := 2; i < n; i++ {
		if isPrim[i] {
			res++
		}
	}
	return res
}
