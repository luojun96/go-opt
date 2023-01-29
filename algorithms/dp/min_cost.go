package main

// https://leetcode.cn/problems/JEj789/
// input: costs = [[17,2,17],[16,16,5],[14,3,19]]
// output: 10
// 动态规划
func minCost(costs [][]int) int {
	dp := costs[0]
	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)
		for j, c := range cost {
			dpNew[j] = min(dp[(j+1)%3], dp[(j+2)%3]) + c
		}
		dp = dpNew
	}
	return min(min(dp[0], dp[1]), dp[2])
}

// func min(x, y int) int {
// 	if x > y {
// 		return y
// 	}
// 	return x
// }
