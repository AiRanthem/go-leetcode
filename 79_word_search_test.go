package leetcode

import "testing"

func Test79(t *testing.T) {
	exist([][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}, "abcb")
}
