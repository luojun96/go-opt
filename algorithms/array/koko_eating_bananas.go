package algorithms

// https://leetcode.cn/problems/koko-eating-bananas/
// input: piles = [3, 6, 7, 11], h = 8
// output: 4
func minEatingSpeed(piles []int, h int) int {
	low, high := 1, getMax(piles)
	for low <= high {
		mid := low + (high-low)/2
		if canFinish(piles, mid, h) {
			if mid == low || !canFinish(piles, mid-1, h) {
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

func canFinish(piles []int, speed int, h int) bool {
	t := 0
	for _, p := range piles {
		t += timeOf(p, speed)
	}
	return t <= h
}

func timeOf(n int, speed int) int {
	if n%speed == 0 {
		return n / speed
	} else {
		return n/speed + 1
	}
}

func getMax(piles []int) int {
	max := 0
	for _, n := range piles {
		max = maxInt(max, n)
	}
	return max
}

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
