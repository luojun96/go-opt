package array

//https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
//input: nums = [0,0,1,1,1,2,2,3,3,4]
//output: 5, nums = [0,1,2,3,4]
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
