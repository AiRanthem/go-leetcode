package leetcode

func findMaxForm(strs []string, m int, n int) int {
	ones, zeroes := make([]int, len(strs)), make([]int, len(strs))
	for i, str := range strs {
		one, zero := 0, 0
		for j := 0; j < len(str); j++ {
			if str[j] == '0' {
				zero++
			} else {
				one++
			}
		}
		ones[i], zeroes[i] = one, zero
	}
	dp := new2dSlice(m+1, n+1)
	// dp[m][n] = max(dp[m][n], dp[m-zeroes[i]][n-ones[i]])
	for i := 0; i < len(strs); i++ {
		for z := m; z >= zeroes[i]; z-- {
			for o := n; o >= ones[i]; o-- {
				dp[z][o] = maxInt(dp[z][o], dp[z-zeroes[i]][o-ones[i]]+1)
			}
		}
	}
	return dp[m][n]
}
