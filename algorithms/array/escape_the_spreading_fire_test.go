package array

import "testing"

func Test_maximumMinutes(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				grid: [][]int{
					{0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 2, 2, 1, 0},
					{0, 2, 0, 0, 1, 2, 0},
					{0, 0, 2, 2, 2, 0, 2},
					{0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				grid: [][]int{
					{0, 0, 0},
					{2, 2, 0},
					{1, 2, 0},
				},
			},
			want: 1e9,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumMinutes(tt.args.grid); got != tt.want {
				t.Errorf("maximumMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
