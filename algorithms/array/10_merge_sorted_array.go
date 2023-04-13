package algorithms

// https://leetcode.cn/problems/merge-sorted-array/
// input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// output: [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m + n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}
