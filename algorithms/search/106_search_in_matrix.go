package main

import "sort"

// https://leetcode.cn/problems/search-a-2d-matrix/
// input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// output: true
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {
		mid := low + (high-low)>>1
		x := matrix[mid/n][mid%n]
		if x < target {
			low = mid + 1
		} else if x > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
