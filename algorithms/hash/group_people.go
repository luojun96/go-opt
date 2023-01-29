package main

// https://leetcode.cn/problems/group-the-people-given-the-group-size-they-belong-to/

func groupTHePeople(groupSizes []int) [][]int {
	res := make([][]int, 0)
	groups := map[int][]int{}
	
	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}

	for size, people := range groups {
		for i := 0; i < len(people); i += size {
			res = append(res, people[i:i+size])
		}
	}

	return res
}
