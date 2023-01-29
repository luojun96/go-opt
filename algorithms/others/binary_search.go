package main

func search_in_non_duplidated_array(array []int, value int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] == value {
			return mid
		} else if array[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func search_bigger_or_equal_element_in_duplidated_array(array []int, value int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] >= value {
			if mid == 0 || array[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search_less_or_equal_element_in_dupliated_array(array []int, value int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] <= value {
			if mid == len(array)-1 || array[mid+1] > value {
				return mid
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
