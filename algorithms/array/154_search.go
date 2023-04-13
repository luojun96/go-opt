package algorithms

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// input: nums = [5,7,7,8,8,10], target = 8
// output: 2
func search(nums []int, target int) int {
	var bsearch func(array []int, value int, lower bool) int
	bsearch = func(array []int, value int, lower bool) (index int) {
		low, high := 0, len(array)-1
		for low <= high {
			mid := low + (high-low)>>1
			if nums[mid] > target {
				high = mid - 1
			} else if nums[mid] < target {
				low = mid + 1
			} else {
				if lower {
					if mid == 0 || array[mid-1] != value {
						return mid
					} else {
						high = mid - 1
					}
				} else {
					if mid == len(array)-1 || array[mid+1] != value {
						return mid
					} else {
						low = mid + 1
					}

				}
			}
		}
		return -1
	}

	leftIndex := bsearch(nums, target, true)
	rightIndex := bsearch(nums, target, false)

	if leftIndex != -1 && rightIndex != -1 && leftIndex <= rightIndex {
		return rightIndex - leftIndex + 1
	} else {
		return 0
	}
}
