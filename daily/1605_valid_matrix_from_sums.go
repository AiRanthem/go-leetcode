package daily

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	result := new2dSlice(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = minInt(rowSum[i], colSum[j])
			rowSum[i] -= result[i][j]
			colSum[j] -= result[i][j]
		}
	}
	return result
}
