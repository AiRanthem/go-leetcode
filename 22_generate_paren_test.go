package leetcode

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T) {
	for _, s := range generateParenthesis(3) {
		fmt.Println(s)
	}
}
