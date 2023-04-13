package algorithms

// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/discuss/2131489/Java-backtracking-solution
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%k != 0 {
		return false
	}

	used := make([]bool, len(nums))
	target := sum / k

	var backtrack func(k int, bucket int, nums []int, start int, used []bool, target int) bool
	backtrack = func(count, bucket int, _ []int, start int, used []bool, target int) bool {
		if k == 0 {
			return true
		}

		return false
	}

	return backtrack(k, 0, nums, 0, used, target)

}
