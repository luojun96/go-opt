package algorithms

// https://leetcode.cn/problems/permutations/
// input: nums = [1,2,3]
// output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// use backtracking to solve this problem
var res [][]int
func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track := []int{}
	backtrack(nums, track)
	return res
}

func backtrack(nums []int, track []int) {
	if len(track) == len(nums) {
		res = append(res, append([]int{}, track...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		} 
		track = append(track, nums[i])
		backtrack(nums, track)
		track = track[:len(track)-1]
	}
}

func contains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
