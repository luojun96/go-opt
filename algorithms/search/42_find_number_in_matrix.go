package algorithms

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
// input:
// [
//
//	[1,   4,  7, 11, 15],
//	[2,   5,  8, 12, 19],
//	[3,   6,  9, 16, 22],
//	[10, 13, 14, 17, 24],
//	[18, 21, 23, 26, 30]
//
// ]
// target = 5
// output: true
func findNumberIn2DArray(matrix [][]int, target int) bool {
	var search func(array []int) int
	search = func(array []int) int {
		low, high := 0, len(array)-1
		for low <= high {
			mid := low + (high-low)>>1
			if array[mid] == target {
				return mid
			} else if array[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return -1
	}

	for _, row := range matrix {
		i := search(row)
		if i < len(row) && i >= 0 && row[i] == target {
			return true
		}
	}
	return false
}
