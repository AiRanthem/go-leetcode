package leetcode

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestDirRing(t *testing.T) {
	direction := ring.New(4)
	direction.Value = R
	direction = direction.Next()
	direction.Value = D
	direction = direction.Next()
	direction.Value = L
	direction = direction.Next()
	direction.Value = U
	direction = direction.Next()

	for i := 0; i < 8; i++ {
		fmt.Println(direction.Value)
		direction = direction.Next()
	}
}

func TestDraw(t *testing.T) {
	matrix := generateMatrix(5)
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, i := range row {
			fmt.Printf("%2d ", i)
		}
		fmt.Println()
	}
}
