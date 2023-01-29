package main

import "testing"

func Test_replaceWords_01(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				dictionary: []string{"cat", "bat", "rat"},
				sentence:   "ca",
			},
			want: "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords_01(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWords_01() = %v, want %v", got, tt.want)
			}
		})
	}
}
