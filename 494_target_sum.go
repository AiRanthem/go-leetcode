package leetcode

// x + y = sum
// x - y = target
// 2x = sum + target
// x = (sum+target)/2

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum || -target > sum || (sum-target)%2 == 1 {
		return 0
	}
	c := (sum - target) / 2
	dp := make([]int, c+1) // 多少种装满的方法
	dp[0] = 1              // 记住就行
	for i := 0; i < len(nums); i++ {
		for j := c; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[c]
}
