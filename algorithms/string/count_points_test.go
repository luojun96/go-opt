package string

import (
	"testing"
)

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case0",
			args: args{
				rings: "R0G0B0",
			},
			want: 1,
		},
		{
			name: "case1",
			args: args{
				rings: "B0B6G0R6R0R6G9",
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				rings: "B0R0G0R9R0B0G0",
			},
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.rings); got != tt.want {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPoints2(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case0",
			args: args{
				rings: "R0G0B0",
			},
			want: 1,
		},
		{
			name: "case1",
			args: args{
				rings: "B0B6G0R6R0R6G9",
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				rings: "B0R0G0R9R0B0G0",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints2(tt.args.rings); got != tt.want {
				t.Errorf("countPoints2() = %v, want %v", got, tt.want)
			}
		})
	}
}
