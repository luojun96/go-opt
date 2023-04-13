package algorithms

import "fmt"

func solveNQueens(n int) [][]string {
	res := make([]int, n)
	calNQueens(0, res)
	return printQueen(res)
}

func calNQueens(row int, res []int) {
	n := len(res)
	if row == n {
		return
	}

	for column := 0; column < n; column++ {
		if isOk(row, column, res) {
			res[row] = column
			calNQueens(row+1, res)
		}
	}
}

func isOk(row, column int, res []int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if res[i] == column {
			return false
		}
		if leftup >= 0 && res[i] == leftup {
			return false
		}
		if leftup < len(res) && res[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

func printQueen(a []int) [][]string {
	n := len(a)
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		res[i] = make([]string, n)
	}
	for row := 0; row < n; row++ {
		for column := 0; column < n; column++ {
			if a[row] == column {
				res[row][column] = "Q"
			} else {
				res[row][column] = "*"
			}
		}
	}
	return res
}

func algorithms() {
	res := solveNQueens(8)
	for _, v := range res {
		fmt.Println(v)
	}
}
