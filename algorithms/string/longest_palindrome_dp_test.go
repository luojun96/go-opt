package algorithms

import "testing"

func Test_longestPalindromeByDP(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{"cbbd"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeByDP(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeByDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
