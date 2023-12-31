package sort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{6, 1, 3, 5, 7, 2, 4, 9, 11, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("quickSort = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
