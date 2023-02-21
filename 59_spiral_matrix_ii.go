package leetcode

import "container/ring"

const (
	R = iota
	D
	L
	U
)

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	direction := ring.New(4)
	direction.Value = R
	direction = direction.Next()
	direction.Value = D
	direction = direction.Next()
	direction.Value = L
	direction = direction.Next()
	direction.Value = U
	direction = direction.Next()

	x, y := 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		nx, ny := nextCoord(x, y, direction.Value.(int))
		if !canWrite(nx, ny, matrix) {
			direction = direction.Next()
			nx, ny = nextCoord(x, y, direction.Value.(int))
		}
		x, y = nx, ny
	}
	return matrix
}

func nextCoord(x, y, direct int) (int, int) {
	switch direct {
	case U:
		return x - 1, y
	case D:
		return x + 1, y
	case L:
		return x, y - 1
	default: // R
		return x, y + 1
	}
}

func canWrite(x, y int, matrix [][]int) bool {
	n := len(matrix)
	return x < n && x >= 0 && y < n && y >= 0 && matrix[x][y] == 0
}
