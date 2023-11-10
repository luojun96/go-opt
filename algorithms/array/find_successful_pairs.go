package array

import (
	"slices"
	"sort"
)

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/?envType=daily-question&envId=2023-11-10

func findSuccessfulPairs(spells []int, potions []int, success int64) []int {
	pairs := make([]int, len(spells))
	slices.Sort(potions)
	for i, spell := range spells {
		pairs[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/spell+1)
	}
	return pairs
}
