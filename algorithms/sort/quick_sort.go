package main

func quickSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, low, high int) {
	if low >= high {
		return
	}

	var partition func(nums []int, low, high int) int
	partition = func(nums []int, low, high int) int {
		x := nums[high]
		i := low - 1
		for j := low; j < high; j++ {
			if nums[j] <= x {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i+1], nums[high] = nums[high], nums[i+1]
		return i + 1
	}

	q := partition(nums, low, high)
	sort(nums, low, q-1)
	sort(nums, q+1, high)
}
