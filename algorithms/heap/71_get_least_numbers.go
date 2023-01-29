package main

// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/description/?languageTags=golang
// input: arr = [3,2,1], k = 2
// output: [1,2] or [2,1]
func getLeastNumbers(arr []int, k int) []int {
	var quickSelect func(nums []int, left, right int) []int
	quickSelect = func(nums []int, left, right int) []int {
		if left > right {
			return nil
		}

		i, j, pivot := left, right, nums[left]
		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}

			for i < j && nums[i] <= pivot {
				i++
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[left] = nums[left], nums[i]
		quickSelect(nums, left, i-1)
		quickSelect(nums, i+1, right)
		return nums

	}
	nums := quickSelect(arr, 0, len(arr)-1)
	return nums[:k]
}
