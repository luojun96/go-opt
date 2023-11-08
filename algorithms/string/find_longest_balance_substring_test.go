package string

import "testing"

func Test_findTheLongestBalancedSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "01000111",
			},
			want: 6,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestBalancedSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestBalancedSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
