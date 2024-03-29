package search

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)>>1

	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := []int{}

	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}

	for ; j <= end; j++ {
		tmp = append(tmp, nums[i])
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}

	return cnt
}
