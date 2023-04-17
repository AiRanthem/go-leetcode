package top100

import "math"

func maxCoins(nums []int) int {
	nums = append(nums, 1, 1)
	var balloons []int
	balloons = append(balloons, 1)
	balloons = append(balloons, nums...)
	balloons = append(balloons, 1)
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 2; j < len(nums); j++ {
			dp[i][j] = math.MinInt
			for k := i + 1; k < j; k++ {
				s := balloons[i]*balloons[k]*balloons[j] + dp[i][k] + dp[k][j]
				if s > dp[i][j] {
					dp[i][j] = s
				}
			}
		}
	}
	return dp[0][len(nums)-1]
}
