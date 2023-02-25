package leetcode

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n+1; i++ {
		for j := 1; j <= 2; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}
