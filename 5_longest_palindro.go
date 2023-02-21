package leetcode

func longestPalindrome(s string) string {
	ss := []byte(s)
	l := len(ss)
	maxLen, maxI, maxJ := 1, 0, 0
	dp := new2dSlice(l, l)
	/**
	dp[i][j] = palindrome len of s[k:j] (0 if not)
	dp[i][j] = 	dp[k+1][j-1] + 2 (s[k] == s[j])
				0 (s[k] != s[j])
	*/
	for i := 0; i < l; i++ {
		dp[i][i] = 1
	}
	for k := 1; k < l; k++ {
		for i := 0; i < l-k; i++ {
			j := i + k
			if ss[i] == ss[j] && (dp[i+1][j-1] > 0 || i > j-2) {
				dp[i][j] = dp[i+1][j-1] + 2
			}
			if maxLen < dp[i][j] {
				maxLen = dp[i][j]
				maxI = i
				maxJ = j
			}
		}
	}
	return string(ss[maxI : maxJ+1])
}
