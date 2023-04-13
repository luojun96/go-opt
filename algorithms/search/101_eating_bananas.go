package algorithms

// https://leetcode.cn/problems/koko-eating-bananas/
// input: piles = [3,6,7,11], h = 8
// output: 4
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	k := max
	l, r := 1, max
	for l < r {
		mid := l + (r-l)>>1
		time := getTime(piles, mid)
		if time <= int64(h) {
			k = mid
			r = mid
		} else {
			l = mid + 1
		}
	}
	return k
}

func getTime(piles []int, speed int) int64 {
	time := int64(0)
	for _, p := range piles {
		curTime := (p + speed - 1) / speed
		time += int64(curTime)
	}
	return time
}
