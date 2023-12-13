package search

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 6,
			},
			want: 5,
		},
		{
			name: "case3",
			args: args{
				nums:   []int{2, 3, 4, 5},
				target: 1,
			},
			want: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
