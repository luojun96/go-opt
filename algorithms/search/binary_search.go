package main

import "fmt"

// array doesn't contain duplicate item
func search(array []int, value int) int {
	low, high := 0, len(array) - 1
	for low <= high {
		mid := low + (high - low) / 2
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

// array contain duplicated items, and look for the element which bigger or equal to value
func search1(array int[], value int) {
	low, high := 0, len(array) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if array[mid] >= value {
			if mid == 0 || array[mid - 1] < value {
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
func search2(array []int, value int) {
	low, high := 0, len(array) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if array[mid] <= value {
			if mid == len(array) - 1 || array[mid + 1] > value {
				return mid
			}	
		} else {
			high = mid - 1
		}
	}
}

// look for the first element which is equal to value
func search3(array []int, value int) {
	low, high := 0, n - 1
	for low <= high {
		mid := low + (high - low) / 2
		if array[mid] > value {
			high = mid
		} else if array[mid] < value {
			low = mid
		} else {
			if mid == 0 || a[mid - 1] < value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
}
func main() {
	fmt.Println("vim-go")
}
