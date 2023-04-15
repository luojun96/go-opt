package search

// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/description/
// input: m = 2, n = 3, k = 1
// output: 3
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	q := make([][]int, 0)
	dx := []int{0, 1}
	dy := []int{1, 0}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	q = append(q, []int{0, 0})
	visited[0][0] = true

	res := 1
	for len(q) > 0 {
		cell := q[0]
		q = q[1:]

		x := cell[0]
		y := cell[1]

		for i := 0; i < 2; i++ {
			tx := dx[i] + x
			ty := dy[i] + y

			if tx < 0 || tx >= m || ty < 0 || ty >= n || visited[tx][ty] || get(tx)+get(ty) > k {
				continue
			}
			q = append(q, []int{tx, ty})
			visited[tx][ty] = true
			res++
		}
	}
	return res
}

func get(x int) int {
	res := 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}
