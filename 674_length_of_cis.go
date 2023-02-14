package leetcode

func findLengthOfLCIS(nums []int) int {
	var dp = make([]int, len(nums))
	var result = 1
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		result = maxInt(result, dp[i])
	}
	return result
}
