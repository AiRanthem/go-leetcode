package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("the sky is blue"))
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		st string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{"the sky is blue"},
			"blue is sky the",
		},
		{
			"test",
			args{"  hello world  "},
			"world hello",
		},
		{
			"test",
			args{"a good   example"},
			"example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.st); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
