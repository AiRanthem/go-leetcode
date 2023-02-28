package leetcode

/**
dp[i][j]: 两单词前几位的编辑距离，不妨设word1短。
if w1[i] == w2[j]:	dp[i][j] =min{dp[i-1][j-1], 1 + dp[i][j-1]}
if w1[i] != w2[j]:	dp[i][j] =min{1 + dp[i-1][j-1], 1 + dp[i][j-1]}
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := new2dSlice(m+1, n+1)
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			tail := 1
			if word1[i-1] == word2[j-1] {
				tail = 0
			}
			dp[i][j] = minTInt(tail+dp[i-1][j-1], 1+dp[i][j-1], 1+dp[i-1][j])
		}
	}
	return dp[m][n]
}
