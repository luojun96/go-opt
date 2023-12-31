package sort

func quickSort(nums []int) {
	qSort(nums, 0, len(nums)-1)
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

func qSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	var partition func(left, right int) int
	partition = func(left, right int) int {
		p := nums[right]
		i := left - 1
		for j := left; j < right; j++ {
			if nums[j] <= p {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i+1], nums[right] = nums[right], nums[i+1]
		return i + 1
	}

	pivot := partition(low, high)
	qSort(nums, low, pivot-1)
	qSort(nums, pivot+1, high)
}
