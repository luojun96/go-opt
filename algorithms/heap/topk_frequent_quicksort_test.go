package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequentByQuickSelect(t *testing.T) {
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
			name: "normal case",
			args: args{[]int{5, 3, 1, 1, 1, 3, 5, 73, 1}, 3},
			want: []int{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentByQuickSelect(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentByQuickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
