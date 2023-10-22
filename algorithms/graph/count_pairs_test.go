package graph

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"no non-pairs", args{3, [][]int{{0, 1}, {0, 2}, {1, 2}}}, 0},
		{name: "has non-pairs",
			args: args{
				n: 7,
				edges: [][]int{
					{0, 2},
					{0, 5},
					{2, 4},
					{4, 5},
					{1, 6},
				},
			},
			want: 14,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
