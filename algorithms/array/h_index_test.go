package array

import (
	"testing"
)

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				citations: []int{
					3, 0, 6, 1, 5,
				},
			},
			want: 3,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hIndex2(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				citations: []int{
					0, 1, 3, 5, 6,
				},
			},
			want: 3,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex2(tt.args.citations); got != tt.want {
				t.Errorf("hIndex2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hIndex3(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				citations: []int{
					0, 1, 3, 5, 6,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex3(tt.args.citations); got != tt.want {
				t.Errorf("hIndex3() = %v, want %v", got, tt.want)
			}
		})
	}
}
