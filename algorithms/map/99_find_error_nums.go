package algorithms

func findErrorNums(nums []int) []int {
	cnt := map[int]int{}

	for _, v := range nums {
		cnt[v]++
	}

	res := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if c := cnt[i]; c == 2 {
			res[0] = i
		} else if c == 0 {
			res[1] = i
		}
	}

	return res
}
