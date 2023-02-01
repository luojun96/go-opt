package main

func gReduce[T1 any, T2 any] (arr []T1, init T2, f func(T2, T1) T2) T2 {
    result := init
    for _, elem := range arr {
        result = f(result, elem) 
    }
    return result
}
