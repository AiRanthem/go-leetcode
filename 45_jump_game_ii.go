package leetcode

/*
dp[i] 从i跳到底的最少次数
dp[i] = 1 + min(dp[k])
*/
// func jump(nums []int) int {
// 	if len(nums) <= 1 {
// 		return 0
// 	}
// 	dp := make([]int, len(nums)-1)
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		if i+nums[i] >= len(nums)-1 {
// 			dp[i] = 1
// 		} else if nums[i] != 0 {
// 			dp[i] = dp[i+1]
// 			for k := 2; k <= nums[i]; k++ {
// 				if dp[i] > dp[i+k] {
// 					dp[i] = dp[i+k]
// 				}
// 			}
// 			dp[i]++
// 		} else {
// 			dp[i] = 1001
// 		}
// 	}
// 	return dp[0]
// }

func jump(nums []int) int {
	var l, r, nextR, step int
	for r < len(nums)-1 {
		nextR = r
		for l <= r {
			if l+nums[l] > nextR {
				nextR = l + nums[l]
			}
			l++
		}
		step++
		r = nextR
	}
	return step
}
