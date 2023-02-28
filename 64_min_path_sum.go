package leetcode

/*
dp[i][j]: 最小路径长度
dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
单行优化：dp[j] = min(dp[j], dp[j-1]) + grid[i][j], j := 1~n
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = grid[0][j] + dp[j-1]
	}
	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			if dp[j] > dp[j-1] {
				dp[j] = dp[j-1]
			}
			dp[j] += grid[i][j]
		}
	}
	return dp[n-1]
}
