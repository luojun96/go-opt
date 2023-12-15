package string

import "testing"

func Test_lengthOfLongestSubstring1(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring1() = %v, want %v", got, tt.want)
			}
		})
	}
}
