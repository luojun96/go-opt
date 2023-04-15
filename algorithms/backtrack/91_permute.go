package backtrack

// https://leetcode.cn/problems/permutations/
// input: nums = [1,2,3]
// output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backtrack(nums, len(nums), []int{})
	return res
}

func backtrack(nums []int, numslen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}

	for i := 0; i < numslen; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backtrack(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}
