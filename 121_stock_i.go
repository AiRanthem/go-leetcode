package leetcode

func maxProfit(prices []int) int {
	result := 0
	dp := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		dp[i] = dp[i-1] + prices[i] - prices[i-1]
		if dp[i] < 0 {
			dp[i] = 0
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
