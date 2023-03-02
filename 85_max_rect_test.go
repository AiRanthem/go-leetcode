package leetcode

import (
	"fmt"
	"testing"
)

func Test85(t *testing.T) {
	var input [][]byte
	for i := 0; i < 100; i++ {
		input = append(input, []byte{'1', '1', '1', '1', '1'})
	}
	n := maximalRectangle(input)
	fmt.Println(n)
}
