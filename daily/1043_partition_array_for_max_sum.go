package daily

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		maxNum := arr[i]
		for j := i; j >= 0 && j > i-k; j-- {
			// [j,i] is the last partition
			if arr[j] > maxNum {
				maxNum = arr[j]
			}
			dp[i+1] = maxInt(dp[i+1], maxNum*(i-j+1)+dp[j])
		}
	}
	return dp[len(arr)]
}
