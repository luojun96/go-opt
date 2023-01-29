package main

import "testing"

func Test_reversePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal case",
			args: args{nums: []int{7, 5, 6, 4}}, // mid := (len(nums) - 1 - 0) / 2 = (4 - 1) / 2 = 1
			// a := [7, 5], b := [6, 4]
			// for [7, 5], after merge sort, output = [5, 7], count = 1 ([7, 5])
			// for [6, 4], after merge sort, output = [4, 6], count = 1 ([6, 4])
			// for [5, 7, 4, 6], i = 0, mid = 1, j = 2
			// i = 0, j = 3, tmp = [4], count + 1 ([5, 4])
			// i = 0, j = 3, tmp = [4, 5]
			// for index := i ; index <= mid, index ++ {
			//		 if nums[index] > nums[j] {
			//			 count ++
			// 		 }
			// } then j++
			// j = 1, j = 3
			// i = 1, j = 4, tmp = [4, 5, 6], count + 1 ([7, 6])
			//
			want: 5,
		},
		{
			name: "case 1",
			args: args{nums: []int{4, 5, 8, 6, 7}}, // mid := (len(nums) - 1 - 0) / 2 = (5 - 1) / 2 = 2;
			// a := [4, 5, 8], b := [6, 7], i = 0, j = 0, k = 0, count = 0
			// i = 1, j = 0, k = 1, tmp = [4]
			// i = 2, j = 0, k = 2, tmp = [4, 5]
			// i = 2, j = 1, k = 3, tmp = [4, 5, 6], count = 1 ([8, 6])
			// i = 2, j = 2, k = 4, tmp = [4, 5, 6, 7], count = 2 ([8,6], [8,7])
			// since j > end (=1), end loop
			// since i <= mid,
			// i = 3, k = 5, tmp = [4,5,6,7,8]
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
