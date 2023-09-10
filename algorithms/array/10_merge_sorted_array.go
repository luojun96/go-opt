package array

// https://leetcode.cn/problems/merge-sorted-array/
// input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// output: [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
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

func merge1(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if p1 == -1 {
			nums1[i] = nums2[p2]
			p2--
			continue
		}

		if p2 == -1 {
			nums1[i] = nums1[p1]
			p1--
			continue
		}

		if nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}
