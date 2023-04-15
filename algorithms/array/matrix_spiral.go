package array

// https://leetcode.cn/problems/spiral-matrix/

func printBySpiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	rows := len(matrix)
	columns := len(matrix[0])
	res := make([]int, 0, rows*columns)

	start := 0

	for columns > start*2 && rows > start*2 {
		endCol := columns - 1 - start
		endRow := rows - 1 - start

		// from left to right
		for i := start; i <= endCol; i++ {
			res = append(res, matrix[start][i])
		}

		// from up to down
		if start < endRow {
			for i := start + 1; i <= endRow; i++ {
				res = append(res, matrix[i][endCol])
			}
		}

		// from right to left
		if start < endCol && start < endRow {
			for i := endCol - 1; i >= start; i-- {
				res = append(res, matrix[endRow][i])
			}
		}

		// from down to up
		if start < endCol && start < endRow-1 {
			for i := endRow - 1; i >= start+1; i-- {
				res = append(res, matrix[i][start])
			}
		}
	}
	return res
}
