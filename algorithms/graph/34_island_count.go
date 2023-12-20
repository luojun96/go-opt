package graph

// https://leetcode.cn/problems/number-of-islands/
type Island struct {
	Row int
	Col int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var bfs func(row, col int)
	bfs = func(row, col int) {
		grid[row][col] = 0
		q := []*Island{{
			Row: row,
			Col: col,
		}}
		for len(q) > 0 {
			size := len(q)
			for i := 0; i < size; i++ {
				curr := q[0]
				q = q[1:]

				// up
				if curr.Row > 0 && grid[curr.Row-1][curr.Col] == '1' {
					grid[curr.Row-1][curr.Col] = '0'
					q = append(q, &Island{
						Row: curr.Row - 1,
						Col: curr.Col,
					})
				}
				// down
				if curr.Row < len(grid)-1 && grid[curr.Row+1][curr.Col] == '1' {
					grid[curr.Row+1][curr.Col] = '0'
					q = append(q, &Island{
						Row: curr.Row + 1,
						Col: curr.Col,
					})
				}
				// right
				if curr.Col < len(grid[0])-1 && grid[curr.Row][curr.Col+1] == '1' {
					grid[curr.Row][curr.Col+1] = '0'
					q = append(q, &Island{
						Row: curr.Row,
						Col: curr.Col + 1,
					})
				}
				// left
				if curr.Col > 0 && grid[curr.Row][curr.Col-1] == '1' {
					grid[curr.Row][curr.Col-1] = '0'
					q = append(q, &Island{
						Row: curr.Row,
						Col: curr.Col - 1,
					})
				}
			}
		}
	}

	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				// use bfs to search the grid, and set the found island value to '0'
				bfs(i, j)
			}
		}
	}
	return num
}

type island struct {
	row int
	col int
}

func bfs1(grid [][]byte, row, col int) {
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
