package algorithms

import "fmt"

type node struct {
	r int
	c int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				// use bfs to search the grid, and set the found islands value to '0'
				bfs(grid, i, j)
			}
		}
	}
	return num
}

func bfs(grid [][]byte, r, c int) {
	grid[r][c] = 0
	q := []*node{
		{r, c},
	}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cNode := q[0]
			q = q[1:]

			// up
			if cNode.r > 0 && grid[cNode.r-1][cNode.c] == '1' {
				grid[cNode.r-1][cNode.c] = '0'
				q = append(q, &node{cNode.r - 1, cNode.c})
			}

			// down
			if cNode.r < len(grid)-1 && grid[cNode.r+1][cNode.c] == '1' {
				grid[cNode.r+1][cNode.c] = '0'
				q = append(q, &node{cNode.r + 1, cNode.c})
			}

			// left
			if cNode.c > 0 && grid[cNode.r][cNode.c-1] == '1' {
				grid[cNode.r][cNode.c-1] = '0'
				q = append(q, &node{cNode.r, cNode.c - 1})
			}

			// right
			if cNode.c < len(grid[0])-1 && grid[cNode.r][cNode.c+1] == '1' {
				grid[cNode.r][cNode.c+1] = '0'
				q = append(q, &node{cNode.r, cNode.c + 1})
			}
		}
	}
}
func algorithms() {
	fmt.Println("vim-go")
}
