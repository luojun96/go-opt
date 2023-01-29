package main

// https://leetcode.cn/problems/maximal-network-rank/description/
// input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
// output: 4
func maximalNetworkRank(n int, roads [][]int) int {
	res := 0
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	cities := make([]int, n)
	for i := 0; i < len(roads); i++ {
		road := roads[i]
		matrix[road[0]][road[1]] = 1
		matrix[road[1]][road[0]] = 1
		cities[road[0]]++
		cities[road[1]]++
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = max(res, cities[i]+cities[j]-matrix[i][j])
		}
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
