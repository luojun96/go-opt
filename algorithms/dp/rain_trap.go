package dp

func trap(height []int) int {
	n := len(height)
	res := 0
	for i := 0; i < n; i++ {
		lMax, rMax := 0, 0
		for j := i; j < n; j++ {
			rMax = maxValue(rMax, height[j])
		}
		for j := i; j >= 0; j-- {
			lMax = maxValue(lMax, height[j])
		}
		res += min(lMax, rMax) - height[i]
	}
	return res
}

func maxValue(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
