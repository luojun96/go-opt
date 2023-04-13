package algorithms

func getLeasNumbers(arr []int, k int) []int {
	i := quickSelect(arr, 0, len(arr)-1, k)
	return arr[:i]
}

func quickSelect(arr []int, l, r, k int) int {
	q := partition1(arr, l, r)
	if q == k {
		return q
	} else if q < k {
		return quickSelect(arr, q+1, r, k)
	} else {
		return quickSelect(arr, l, q-1, k)
	}
}

func partition1(arr []int, l, r int) int {
	p := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] <= p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
