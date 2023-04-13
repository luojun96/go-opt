package algorithms

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weight []int
		days   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				weight: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				days:   5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weight, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
