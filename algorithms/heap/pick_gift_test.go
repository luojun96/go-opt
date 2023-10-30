package heap

import "testing"

func Test_pickGifts(t *testing.T) {
	type args struct {
		gifts []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{
				gifts: []int{25, 64, 9, 4, 100},
				k:     4,
			},
			want: 29,
		},
		{
			name: "case2",
			args: args{
				gifts: []int{1, 1, 1, 1},
				k:     4,
			},
			want: 4,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickGifts(tt.args.gifts, tt.args.k); got != tt.want {
				t.Errorf("pickGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}
