package leetcode

// 最长公共连续子序列
func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := new2dSlice(m+1, n+1)
	result := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				result = maxInt(result, dp[i][j])
			}
		}
	}
	return result
}
