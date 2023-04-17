package cmath

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal case",
			args: args{123},
			want: 321,
		}, {
			name: "with leading case",
			args: args{10200},
			want: 201,
		}, {
			name: "navigate number",
			args: args{-12301},
			want: -10321,
		},
		{
			name: "max",
			args: args{1534236469},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
