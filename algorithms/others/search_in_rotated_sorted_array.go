package others

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
