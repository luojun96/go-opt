package array

import "fmt"

// https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/
// input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5
// output: 15
func shipWithinDays(weight []int, days int) int {
	low, high := getMax(weight), getSum(weight)
	for low <= high {
		mid := low + (high-low)/2
		if canShip(weight, days, mid) {
			if mid == low || !canShip(weight, days, mid-1) {
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

func canShip(weight []int, days int, cap int) bool {
	fmt.Printf("cap = %d\n", cap)
	i := 0
	for d := 1; d <= days; d++ {
		maxCap := cap
		maxCap -= weight[i]
		for maxCap >= 0 {
			i++
			if i == len(weight) {
				fmt.Printf("day %d: i = %d\n", d, i)
				return true
			}
			maxCap -= weight[i]
		}
		fmt.Printf("day %d: i = %d\n", d, i)
	}
	return false
}

func getSum(weight []int) int {
	total := 0
	for _, w := range weight {
		total += w
	}
	return total
}
