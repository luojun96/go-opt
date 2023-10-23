package string

import (
	"testing"
)

func Test_countSeniors(t *testing.T) {
	type args struct {
		details []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "no seniors",
			args: args{
				details: []string{
					"1313579440F2036",
					"2921522980M5644",
				},
			},
			want: 0,
		}, {
			name: "seniors exists",
			args: args{
				details: []string{
					"7868190130M7522",
					"5303914400F9211",
					"9273338290F4010",
				},
			},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSeniors(tt.args.details); got != tt.want {
				t.Errorf("countSeniors() = %v, want %v", got, tt.want)
			}
		})
	}
}
