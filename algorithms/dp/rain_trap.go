package dp

func trap(height []int) int {
	n := len(height)
	res := 0

	var max func(x, y int) int
	max = func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 0; i < n; i++ {
		lMax, rMax := 0, 0
		for j := i; j < n; j++ {
			rMax = max(rMax, height[j])
		}
		for j := i; j >= 0; j-- {
			lMax = max(lMax, height[j])
		}
		res += min(lMax, rMax) - height[i]
	}
	return res
}
