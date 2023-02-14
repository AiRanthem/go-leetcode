package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	var m, n = len(text1), len(text2)
	var dp = new2dSlice(m+1, n+1)
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
