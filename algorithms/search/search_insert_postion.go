package search

func searchInsert(nums []int, target int) int {
	var res int
	var search func(l, h int) int
	search = func(l, h int) int {
		for l <= h {
			mid := l + (h-l)>>1
			if nums[mid] >= target {
				if mid > 0 && nums[mid-1] < target {
					return mid
				} else if mid == 0 {
					return 0
				} else {
					h = mid - 1
				}
			} else {
				l = mid + 1
			}
		}
		return len(nums)
	}
	res = search(0, len(nums)-1)
	return res
}
