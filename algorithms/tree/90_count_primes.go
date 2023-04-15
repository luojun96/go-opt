package tree

// https://leetcode.cn/problems/count-primes/
// input: n = 10
// output: 4
func countPrimes(n int) int {
	var res int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes1(n int) int {
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
