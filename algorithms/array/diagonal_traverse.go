package array

// https://leetcode.cn/problems/diagonal-traverse/
// input: mat = [[1,2,3],[4,5,6],[7,8,9]]
// output: [1,2,4,7,5,3,6,8,9]
// 类似于图的广度优先遍历，以mat[0][0]开始，往右或下寻找它的临近节点，过程中需要反转本层新加的节点，此外节点可能被多次添加

func findDiagonalOrderNormal(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	res := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			x := max(i-n+1, 0)
			y := min(i, n-1)
			for x < m && y >= 0 {
				res = append(res, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := max(i-m+1, 0)
			for x >= 0 && y < n {
				res = append(res, mat[x][y])
				x--
				y++
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findDiagonalOrder(mat [][]int) []int {
	r, c := len(mat), len(mat[0])
	res := make([]int, 0, r*c)

	q := [][2]int{[2]int{0, 0}}
	reversed := false
	isRightDown := true
	for len(q) > 0 {
		size := len(q)
		a := [][2]int{}
		m := map[[2]int]bool{}
		for i := 0; i < size; i++ {
			e := q[0]
			res = append(res, mat[e[0]][e[1]])
			q = q[1:]
			if isRightDown {
				// right
				if e[1] < c-1 {
					n := [2]int{e[0], e[1] + 1}
					if !m[n] {
						a = append(a, n)
						m[n] = true
					}
				}
				// down
				if e[0] < r-1 {
					n := [2]int{e[0] + 1, e[1]}
					if !m[n] {
						a = append(a, n)
						m[n] = true
					}
				}
			} else {
				// down
				if e[0] < r-1 {
					n := [2]int{e[0] + 1, e[1]}
					if !m[n] {
						a = append(a, n)
						m[n] = true
					}
				}
				// right
				if e[1] < c-1 {
					n := [2]int{e[0], e[1] + 1}
					if !m[n] {
						a = append(a, n)
						m[n] = true
					}
				}
			}
		}
		if reversed {
			for i := len(a) - 1; i >= 0; i-- {
				q = append(q, a[i])
			}
		} else {
			q = append(q, a...)
			reversed = true
		}
		isRightDown = !isRightDown
	}
	return res
}
