package algorithms

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}

// the array contains dupliated elements
func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[l]
}
