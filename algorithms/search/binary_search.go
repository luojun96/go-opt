package search

// Binary search:
// A integer array is already sorted in ascending order, and it hasn't duplicate element.
func bSearch(a []int, v int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := l + (h-l)>>1
		if a[m] == v {
			return m
		} else if a[m] > v {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

// Binary search using recursion
func bSearchInternally(a []int, l, h, v int) int {
	if l > h {
		return -1
	}

	m := l + (h-l)>>1
	if a[m] == v {
		return m
	} else if a[m] > v {
		return bSearchInternally(a, l, m-1, v)
	} else {
		return bSearchInternally(a, m+1, h, v)
	}
}

// array contain duplicated items, and look for the element which is bigger or equal to value
func binarySearchForLastBigger(array []int, value int) int {
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

// look for the last element which is less or equal to value
func binarySearchForLastLess(array []int, value int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := low + (high-low)/2
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

// look for the first element which is equal to value
func search3(array []int, value int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] > value {
			high = mid
		} else if array[mid] < value {
			low = mid
		} else {
			if mid == 0 || array[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
