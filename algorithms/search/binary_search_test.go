package search

import (
	"testing"
)

func Test_bSearch(t *testing.T) {
	type args struct {
		a []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal binary search",
			args: args{
				a: []int{
					8, 11, 19, 23, 27, 33, 45, 55, 67, 98,
				},
				v: 19,
			},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearch(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("bSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bSearchInternally(t *testing.T) {
	type args struct {
		a []int
		l int
		h int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal binary search using recursion",
			args: args{
				a: []int{
					8, 11, 19, 23, 27, 33, 45, 55, 67, 98,
				},
				l: 0,
				h: 9,
				v: 19,
			},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchInternally(tt.args.a, tt.args.l, tt.args.h, tt.args.v); got != tt.want {
				t.Errorf("bSearchInternally() = %v, want %v", got, tt.want)
			}
		})
	}
}
