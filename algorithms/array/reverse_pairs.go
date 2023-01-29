package main

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// e: input = [7, 5, 6, 4]
// 	  output: count = 5: [7, 5], [7, 6]. [7, 4], [5, 4], [6, 4]
// merge sort
// 	   [7,5,6,4]
// 	[7,5]	  [6,4]
// [7] [5]	 [6] [4]
//   [5,7]	  [4,6]		// reversed count = 2	[7,5], [6,4]
// 	   [4, 5, 6, 7]		// [5,4], [7,4], [7,6]
// begin, end := 0, len(nums) - 1
//
var count int

func reversePairs(nums []int) int {
	count = 0
	if len(nums) < 2 {
		return 0
	}
	mergeSort(nums, 0, len(nums)-1)
	return count
}

func mergeSort(nums []int, begin, end int) {
	if begin >= end {
		return
	}
	mid := (begin + end) / 2
	mergeSort(nums, begin, mid)
	mergeSort(nums, mid+1, end)
	merge(nums, begin, mid, end)

}

func merge(nums []int, begin, mid, end int) {
	i, j, k := begin, mid+1, 0
	temp := make([]int, end-begin+1)
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			temp[k] = nums[j]
			count += (mid - i + 1)
			// for index := i; index <= mid; index++ {
			// 	if nums[index] > nums[j] {
			// 		count++
			// 	}
			// }
			j++
		} else {
			temp[k] = nums[i]
			i++
		}
		k++
	}
	if i <= mid {
		for i <= mid {
			temp[k] = nums[i]
			i++
			k++
		}
	} else {
		for j <= end {
			temp[k] = nums[j]
			j++
			k++
		}
	}

	for i := 0; i < end-begin+1; i++ {
		nums[begin+i] = temp[i]
	}
}
