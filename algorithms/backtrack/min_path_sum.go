package backtrack

import "math"

// https://leetcode.cn/problems/minimum-path-sum/
// input: grid = [[1,3,1],[1,5,1],[4,2,1]]
// output: 7
func minPathSum(grid [][]int) int {
	rows, columns := len(grid), len(grid[0])
	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, columns)
	}
	var backtrackSum func(grid [][]int, r, c int) int
	backtrackSum = func(grid [][]int, r, c int) int {
		if r == 0 && c == 0 {
			return grid[0][0]
		}

		if res[r][c] != 0 {
			return res[r][c]
		}
		v := math.MaxInt
		if r > 0 {
			v = grid[r][c] + backtrackSum(grid, r-1, c)
		}
		if c > 0 {
			v = min(v, grid[r][c]+backtrackSum(grid, r, c-1))
		}
		res[r][c] = v
		return res[r][c]
	}

	return backtrackSum(grid, rows-1, columns-1)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
