package leetcode

// 最长递增子序列
// dp[i]:		以下标i结尾的LIS长度

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
	}
	var result = 0
	for _, i := range dp {
		result = maxInt(result, i)
	}
	return result
}
