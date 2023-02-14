package leetcode

//在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//
//现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
//
// nums1[i] == nums2[j]
//且绘制的直线不与任何其他连线（非水平线）相交。
//请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
//以这种方法绘制线条，并返回可以绘制的最大连线数。

/**
最长公共子序列

dp[i][j]: 前i，j位的最长公共子序列长度

如果 nums1[i] == nums2[j]，则dp[i][j] = dp[i-1][j-1] + 1

*/

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	var m, n = len(nums1), len(nums2)
	var dp = new2dSlice(m+1, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
