package tree

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *XNode
	}
	tests := []struct {
		name string
		args args
		want *XNode
	}{
		{
			name: "case1",
			args: args{
				root: &XNode{
					Val: 1,
					Left: &XNode{
						Val: 2,
						Left: &XNode{
							Val: 4,
						},
						Right: &XNode{
							Val: 5,
						},
					},
					Right: &XNode{
						Val: 3,
						Right: &XNode{
							Val: 7,
						},
					},
				},
			},
			want: &XNode{
				Val: 1,
				Left: &XNode{
					Val: 2,
					Left: &XNode{
						Val: 4,
						Next: &XNode{
							Val: 5,
						},
					},
					Right: &XNode{
						Val: 5,
						Next: &XNode{
							Val: 7,
						},
					},
					Next: &XNode{
						Val: 3,
						Right: &XNode{
							Val: 7,
						},
					},
				},
				Right: &XNode{
					Val: 3,
					Right: &XNode{
						Val: 7,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
