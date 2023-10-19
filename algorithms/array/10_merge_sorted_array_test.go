package array

import (
	"reflect"
	"testing"
)

func Test_mergeTwoArray(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "merge success",
			args: args{
				num1: []int{1, 2, 3},
				num2: []int{1, 4, 5},
			},
			wantRes: []int{1, 1, 2, 3, 4, 5},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := mergeTwoArray(tt.args.num1, tt.args.num2); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("mergeTwoArray() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
