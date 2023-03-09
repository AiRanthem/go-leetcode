package leetcode

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[j+1] < dp[j] {
				dp[j+1] = dp[j]
			}
			dp[j+1] += grid[i][j]
		}
	}
	return dp[n]
}
