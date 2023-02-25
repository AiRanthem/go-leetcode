package leetcode

import "math"

func numSquares(n int) int {
	var square []int
	for i := 0; i <= n; i++ {
		fact := i * i
		if fact <= n {
			square = append(square, fact)
		} else {
			break
		}
	}
	dp := make([]int, n+1) // dp[i]: 和为i的最小个数; dp[i] = dp[i - square[k]] + 1
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i < len(square); i++ {
		for j := 0; j < n+1; j++ {
			if j >= square[i] {
				dp[j] = minInt(dp[j], dp[j-square[i]]+1)
			}
		}
	}
	return dp[n]
}
