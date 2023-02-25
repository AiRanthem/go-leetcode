package leetcode

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s))
	for j := 1; j < len(s); j++ {
		if s[j] == '(' {
			continue
		}
		n := dp[j-1]
		i := j - n - 1
		if i < 0 {
			continue
		}
		if s[i] == '(' {
			dp[j] = dp[j-1] + 2
			if i > 1 {
				dp[j] += dp[i-1]
			}
		}
	}
	result := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
