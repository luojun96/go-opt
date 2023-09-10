package array

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
// input: nums = [1,1,1,2,2,3]
// output: 5, nums = [1,1,2,2,3]
func removeDuplicatesThree(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
