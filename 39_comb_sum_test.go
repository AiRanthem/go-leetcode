package leetcode

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var a []int
	ChangeSliceContent(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func ChangeSliceContent(a []int) {
	a = append(a, 10)
	a = append(a, 20)
}

func Test39(t *testing.T) {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}
