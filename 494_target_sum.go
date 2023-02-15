package leetcode

// 2x + t = sum; x = (sum - t) / 2

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum-target)/2%2 == 1 {
		return 0
	}
	c := (sum - target) / 2 // capacity of the bag
	dp := make([]int, c+1)  // 全部装满有多少种装法
	for i := 0; i < len(nums); i++ {
		for j := c; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[c]
}
