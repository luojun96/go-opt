package main

import "math"

// https://leetcode.cn/problems/triangle/
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}

	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, f[n-1][i])
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
