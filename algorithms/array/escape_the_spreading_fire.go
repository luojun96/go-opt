package array

// https://leetcode.cn/problems/escape-the-spreading-fire/description/?envType=daily-question&envId=2023-11-09

const inf int = 1e9

var mdirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fireTime := make([][]int, m)
	for i := 0; i < m; i++ {
		fireTime[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fireTime[i][j] = inf
		}
	}

	mbfs(grid, fireTime, m, n)

	res := -1
	l, h := 0, m*n
	for l <= h {
		mid := l + (h-l)>>1
		if check(mid, m, n, grid, fireTime) {
			res = mid
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	if res >= m*n {
		res = inf
	}
	return res
}

func mbfs(grid [][]int, fireTime [][]int, m int, n int) {
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
				fireTime[i][j] = 0
			}
		}
	}

	time := 1
	for len(q) > 0 {
		tmp := q
		q = [][]int{}
		for _, p := range tmp {
			cx, cy := p[0], p[1]
			for _, d := range mdirs {
				nx, ny := cx+d[0], cy+d[1]
				if nx >= 0 && ny >= 0 && nx < m && ny < n {
					if grid[nx][ny] == 2 || fireTime[nx][ny] != inf {
						continue
					}
					q = append(q, []int{nx, ny})
					fireTime[nx][ny] = time
				}
			}
		}
		time++
	}
}

func check(stayTime int, m int, n int, grid [][]int, fireTime [][]int) bool {
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	q := [][]int{}
	q = append(q, []int{0, 0, stayTime})
	for len(q) > 0 {
		tmp := q
		q = [][]int{}

		for _, p := range tmp {
			cx, cy, time := p[0], p[1], p[2]
			for _, d := range mdirs {
				nx, ny := cx+d[0], cy+d[1]
				if nx >= 0 && ny >= 0 && nx < m && ny < n {
					if visit[nx][ny] || grid[nx][ny] == 2 {
						continue
					}
					if nx == m-1 && ny == n-1 {
						return fireTime[nx][ny] >= time+1
					}
					if fireTime[nx][ny] > time+1 {
						q = append(q, []int{nx, ny, time + 1})
						visit[nx][ny] = true
					}
				}
			}
		}
	}
	return false
}
