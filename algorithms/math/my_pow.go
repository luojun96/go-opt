package math

// https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
// intput: x = 2.00000, n = 10
// output: 1024.00000
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
