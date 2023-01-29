package main

import (
	"fmt"
	"math/rand"
	"time"
)

type item struct {
	value     int
	frequency int
}

func topKFrequentByQuickSelect(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	items := []*item{}
	for k, v := range m {
		items = append(items, &item{value: k, frequency: v})
	}

	res := make([]int, 0, k)
	// iTopK := quickSelectByFrequent(items, 0, len(items)-1, k, res)

	fmt.Println("frequent:")
	for i := 0; i < len(items); i++ {
		fmt.Printf("[%d: %d]\n", items[i].value, items[i].frequency)
	}

	// for i := iTopK; i < len(items); i++ {
	// 	res = append(res, items[i].value)
	// }
	return res
}

func quickSelectByFrequent(items []*item, l, r, index int, res []int) int {
	q := partitionByFrequent(items, l, r, res)
	if q == index {
		return q
	} else if q < index {
		return quickSelectByFrequent(items, q+1, r, index, res)
	}
	return quickSelectByFrequent(items, l, q-1, index, res)
}

func partitionByFrequent(items []*item, l, r int, res []int) int {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(r-l+1) + l
	pivot := items[picked].frequency
	i := l
	for j := l; j < r; j++ {
		if items[i].frequency >= pivot {
			fmt.Printf("before switch: i = %d, value = [%d, %d]; j = %d, value = [%d, %d]\n", i, items[i].value, items[i].frequency, j, items[j].value, items[j].frequency)
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	fmt.Printf("before switch: i = %d, value = [%d, %d]; r = %d, value = [%d, %d]\n", i, items[i].value, items[i].frequency, r, items[r].value, items[r].frequency)
	items[i], items[r] = items[r], items[i]
	return i
}
