package main

import (
	"sort"
)

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

func getLeastNumbersBySort(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[:k]
}

func getLeastNumbersByQuickSelect(arr []int, k int) []int {
	var quickSelect func(a []int, l, r, index int) int
	var partition func(a []int, l, r int) int

	quickSelect = func(a []int, l, r, index int) int {
		q := partition(a, l, r)
		if q == index {
			return a[q]
		} else if q < index {
			return quickSelect(a, q+1, r, index)
		} else {
			return quickSelect(a, l, q-1, index)
		}
	}

	partition = func(a []int, l, r int) int {
		pivot := a[r]
		i := l
		for j := l; j < r; j++ {
			if a[j] <= pivot {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[r] = a[r], a[i]
		return i
	}

	index := quickSelect(arr, 0, len(arr)-1, k)
	return arr[:index]
}
