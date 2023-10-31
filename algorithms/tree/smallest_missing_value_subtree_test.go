package tree

import (
	"reflect"
	"testing"
)

func Test_smallestMissingValueSubtree(t *testing.T) {
	type args struct {
		parents []int
		nums    []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				parents: []int{-1, 0, 1, 0, 3, 3},
				nums:    []int{5, 4, 6, 2, 1, 3},
			},
			want: []int{7, 1, 1, 4, 2, 1},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestMissingValueSubtree(tt.args.parents, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestMissingValueSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
