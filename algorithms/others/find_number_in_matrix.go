package main

import "sort"

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)	
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}
