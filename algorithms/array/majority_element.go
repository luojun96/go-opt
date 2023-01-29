package main

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mid := len(nums) / 2
	start := 0
	end := len(nums) - 1
	index := partition(nums, start, end)

	for index != mid {
		if index > mid {
			end = index - 1
			index = partition(nums, start, end)
		} else {
			start = index + 1
			index = partition(nums, start, end)
		}
	}

	res := nums[mid]
	// check more than half
	if check(nums, res) {
		res = 0
	}
	return res
}

func partition(nums []int, begin, end int) int {
	pivot := nums[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i + 1
}

func check(nums []int, v int) bool {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == v {
			c++
		}
	}
	res := false
	if 2 * c >= len(nums) {
		res = true
	}
	return res
}
