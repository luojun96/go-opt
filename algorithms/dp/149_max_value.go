package algorithms

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
// input:
// [
//
//	[1,3,1],
//	[1,5,1],
//	[4,2,1]
//
// ]
// output: 12
func maxValue(grid [][]int) int {
	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	rows, columns := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[rows-1][columns-1]
}
