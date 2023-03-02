package leetcode

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])+1
	left := make([]int, n)
	row := make([]int, n)
	mono := Stack{-1}
	result := 0
	for i := 0; i < m; i++ {
		matrix[i] = append(matrix[i], '0')
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				row[j] += 1
			} else {
				row[j] = 0
			}
			for len(mono) > 1 && row[mono.Top()] >= row[j] {
				toBalance := mono.Pop()
				area := (j - left[toBalance] - 1) * row[toBalance]
				if area > result {
					result = area
				}
			}
			left[j] = mono.Top()
			mono.Push(j)
		}
		mono.Pop()
	}
	return result
}
