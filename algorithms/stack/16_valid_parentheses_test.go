package stack

import "testing"

func Test_isValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "(){}[]",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "{([])}",
			},
			want: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("isValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
