package array

// https://leetcode.cn/problems/k-diff-pairs-in-an-array/
// input: nums = [3, 1, 4, 1, 5], k = 2
// output: 2
func findPairs(nums []int, k int) int {
	visited := map[int]struct{}{}
	res := map[int]struct{}{}

	for _, num := range nums {
		if _, ok := visited[num-k]; ok {
			res[num-k] = struct{}{}
		}
		if _, ok := visited[num+k]; ok {
			res[num] = struct{}{}
		}
		visited[num] = struct{}{}
	}
	return len(res)
}
