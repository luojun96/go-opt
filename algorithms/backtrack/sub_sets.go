package algorithms

// https://leetcode.cn/problems/subsets/
// input: nums = [1,2,3]
// output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	var res [][]int
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return res
}

// func subsets(nums []int) [][]int {
// 	res = [][]int{}
// 	track := []int{}
// 	backtrack(nums, 0, track)
// 	return res
// }

// func backtrack(nums []int, start int, track []int) {
// 	res = append(res, track)
// 	for i := 0; i < len(nums); i++ {
// 		track = append(track, nums[i])
// 		backtrack(nums, i+1, track)
// 		track = track[:len(track)-1]
// 	}
// }
