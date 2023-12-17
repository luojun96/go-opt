package stack

import "testing"

func Test_removeDuplicateKChars(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "deeedbbcccbdaa",
				k: 3,
			},
			want: "aa",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateKChars(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("removeDuplicateKChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
