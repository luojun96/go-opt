package tree

import (
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
			got := connect(tt.args.root)
			compareXTrees(got, tt.want, t)
		})
	}
}

func compareXTrees(a, b *XNode, t *testing.T) {
	if a == nil && b == nil {
		return
	}
	if !(a != nil && b != nil && a.Val == b.Val) {
		t.Errorf("different node: a = %+v, b = %+v", a, b)
	}

	if (a.Next == nil && b.Next != nil) || (a.Next != nil && b.Next == nil) ||
		(a.Next != nil && b.Next != nil && a.Next.Val != b.Next.Val) {
		t.Errorf("different next node: a.next = %+v, b.next = %+v", a.Next, b.Next)
	}
	compareXTrees(a.Left, b.Left, t)
	compareXTrees(a.Right, b.Right, t)
}
