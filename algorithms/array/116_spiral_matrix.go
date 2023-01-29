package main

// https://leetcode.cn/problems/spiral-matrix-ii/
// input: n = 3
// output: [[1,2,3],[8,9,4],[7,6,5]]
type pair struct {
	x int
	y int
}

var dirs = []pair{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}
