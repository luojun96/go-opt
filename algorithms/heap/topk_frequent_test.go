package heap

import (
	"reflect"
	"testing"
)

func Test_topKFrequentByHeap(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 0, 1, 0},
				k:    1,
			},
			want: []int{0},
		},
		{
			name: "case3",
			args: args{
				nums: []int{4, 1, -1, 2, -1, 2, 3},
				k:    2,
			},
			want: []int{-1, 2},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
