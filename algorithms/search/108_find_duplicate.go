package search

// https://leetcode.cn/problems/find-the-duplicate-number/
// input: nums = [1,3,4,2,2]
// output: 2
func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1

	res := -1
	for l <= r {
		mid := l + (r-l)>>1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}

		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			res = mid
		}
	}
	return res
}

func findDuplicateBySlowAndFastPointers(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
