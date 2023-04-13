package algorithms

// https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
// input: nums = [4,3,2,7,8,2,3,1]
// output: [5,6]
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := []int{}
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}

	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}
	return res
}
