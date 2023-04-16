package graph

// https://leetcode.cn/problems/maximal-network-rank/
// n cities, roads, and roads[i] = [ai, bi]
// n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
// 1. use adjustce table to store the roads and cities, matrix := make([][]int, n)
// 2. loop the roads data, and set value in matrix, like matrix[0][1] = 1 and matrix[0][1] = 1
// 3. when loop the roads data, record the count of each city, cities ;= make([]int, n)
// 4. loop all cities, and found the max connection count
func maximalNetworkRank(n int, roads [][]int) int {
	cnt := 0
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
			cnt = max(cnt, cities[i]+cities[j]-matrix[i][j])
		}
	}
	return cnt
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
