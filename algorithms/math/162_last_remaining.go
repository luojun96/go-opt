package algorithms

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// input: n = 5, m = 3
// output: 3
func lastRealgorithmsing(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := lastRealgorithmsing(n-1, m)
	return (m + x) % n
}
