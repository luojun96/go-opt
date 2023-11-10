package array

import (
	"reflect"
	"testing"
)

func Test_findSuccessfulPairs(t *testing.T) {
	type args struct {
		spells  []int
		potions []int
		success int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				spells:  []int{5, 1, 3},
				potions: []int{1, 2, 3, 4, 5},
				success: 7,
			},
			want: []int{4, 0, 3},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSuccessfulPairs(tt.args.spells, tt.args.potions, tt.args.success); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSuccessfulPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
