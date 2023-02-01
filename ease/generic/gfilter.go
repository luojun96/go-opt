package main

func gFilter[T any](arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}

func gFilterIn[T any](arr []T, f func(T) bool) []T {
	return gFilter(arr, true, f)
}

func gFilterOut[T any](arr []T, f func(T) bool) []T {
	return gFilter(arr, false, f)
}
