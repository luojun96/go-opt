package array

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// input: nums = [1,2,3,4]
// output: [1,3,2,4]
func exchange(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	for _, num := range nums {
		if num%2 == 1 {
			res[left] = num
			left++
		} else {
			res[right] = num
			right--
		}
	}
	return res
}

func doExchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && (nums[l]&1) == 1 {
			l++
		}

		for l < r && (nums[r]&1) == 0 {
			r--
		}

		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
