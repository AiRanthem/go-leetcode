package leetcode

/*
dp[i]: i个房屋的最大收益
偷i，则	dp[i] = nums[i]+dp[i-2]
否则 dp[i] = dp[i-1]
*/

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[len(nums)-1]
}
