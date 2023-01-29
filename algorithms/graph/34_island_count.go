package main

// https://leetcode.cn/problems/number-of-islands/solutions/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				// use bfs to search the grid, and set the found island value to '0'
				bfs(grid, i, j)
			}
		}
	}
	return num
}

type island struct {
	row int
	col int
}

func bfs(grid [][]byte, row, col int) {
	grid[row][col] = 0
	queue := []*island{
		{row, col},
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			currIsland := queue[0]
			queue = queue[1:]

			// up
			if currIsland.row > 0 && grid[currIsland.row-1][currIsland.col] == '1' {
				grid[currIsland.row-1][currIsland.col] = '0'
				queue = append(queue, &island{currIsland.row - 1, currIsland.col})
			}
			// down
			if currIsland.row < len(grid)-1 && grid[currIsland.row+1][currIsland.col] == '1' {
				grid[currIsland.row+1][currIsland.col] = '0'
				queue = append(queue, &island{currIsland.row + 1, currIsland.col})
			}
			// right
			if currIsland.col < len(grid[0])-1 && grid[currIsland.row][currIsland.col+1] == '1' {
				grid[currIsland.row][currIsland.col+1] = '0'
				queue = append(queue, &island{currIsland.row, currIsland.col + 1})
			}
			// left
			if currIsland.col > 0 && grid[currIsland.row][currIsland.col-1] == '1' {
				grid[currIsland.row][currIsland.col-1] = '0'
				queue = append(queue, &island{currIsland.row, currIsland.col - 1})
			}
		}
	}
}
