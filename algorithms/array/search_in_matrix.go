package main

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	r, c := 0, len(matrix[0]) - 1
	for r <= len(matrix) - 1 && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false

}
