package leetcode

import (
	"fmt"
	"testing"
)

func Test494(t *testing.T) {
	ways := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println(ways)
}
