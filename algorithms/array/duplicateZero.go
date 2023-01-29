package main

// https://leetcode.cn/problems/duplicate-zeros/
// input: arr = [1,0,2,3,0,4,5,0]
// output: null, arr = [1,0,0,2,3,0,0,4]

func duplicateZeros(arr []int) {
	n := len(arr)
	i := 0
	for i < n - 1 {
		if arr[i] == 0{
			i++
			prev := arr[i]
			arr[i] = 0
			for j := i + 1; j < n; j++ {
				temp := arr[j]
				arr[j] = prev
				prev = temp
			}
		}
		i++
	}
}

func duplicateZeros01(arr []int) {
	n := len(arr)
	top := 0
	i := -1
	for top < n {
		i++
		if arr[i] != 0 {
			top++
		} else {
			top += 2
		}
	}
	j := n - 1
	if top == n + 1 {
		arr[j] = 0
		j--
		i--
	}
	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = arr[i]
			j--
		}
		i--
	}
}
