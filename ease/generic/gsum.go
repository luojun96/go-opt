package main

// Sumable is a type constraint that matches any sumable type.
type Sumable interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64
}

func gSum[T any, U Sumable](arr []T, f func(T) U) U {
	var sum U
	for _, elem := range arr {
		sum += f(elem)
	}
	return sum
}
