package top100

/**
dp[i][j]: 两单词前几位的编辑距离，不妨设word1短。
if w1[i] == w2[j]:	dp[i][j] =min{dp[i-1][j-1], 1 + dp[i][j-1]}
if w1[i] != w2[j]:	dp[i][j] =min{1 + dp[i-1][j-1], 1 + dp[i][j-1]}
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = minInt(dp[i-1][j], dp[i][j-1]) + 1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = minInt(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = minInt(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}
