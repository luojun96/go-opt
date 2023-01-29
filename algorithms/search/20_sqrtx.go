package main

// https://leetcode.cn/problems/sqrtx/
// input: x = 4
// output: 2
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1

	for l <= r {
		mid := l + (r-l)>>1
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	
	return ans
}
