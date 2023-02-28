package leetcode

func numDistinct(s string, t string) int {
	// dp[i][j]: s[:i] t[:j]
	// if s[i] == t[j]:  dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// if s[i] != t[j]:  dp[i][j] = dp[i-1][j]
	dp := make([]int, len(t))
	for i := 0; i < len(s); i++ {
		for j := len(t) - 1; j > 0; j-- {
			if i >= j && s[i] == t[j] {
				dp[j] += dp[j-1]
			}
		}
		if s[i] == t[0] {
			dp[0]++
		}
	}
	return dp[len(dp)-1]
}
