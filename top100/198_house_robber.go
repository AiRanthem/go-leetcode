package top100

func rob(nums []int) int {
	dp := make([]int, len(nums)+1) // i个房子最多偷多少
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = maxInt(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[len(nums)]
}
