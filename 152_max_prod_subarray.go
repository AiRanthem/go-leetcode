package leetcode

import "math"

// max[i], min[i]
// max[i] = min{min[i]} * nums[i] if nums[i] < 0
// vise versa

func maxProduct(nums []int) int {
	n := len(nums)
	max, min := make([]int, n), make([]int, n)
	max[0] = nums[0]
	min[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			if min[i-1] <= 0 {
				max[i] = min[i-1] * nums[i]
			} else {
				max[i] = nums[i]
			}
			if max[i-1] > 0 {
				min[i] = max[i-1] * nums[i]
			} else {
				min[i] = nums[i]
			}
		} else {
			if max[i-1] <= 0 {
				max[i] = nums[i]
			} else {
				max[i] = max[i-1] * nums[i]
			}
			if min[i-1] > 0 {
				min[i] = nums[i]
			} else {
				min[i] = min[i-1] * nums[i]
			}
		}
	}
	result := math.MinInt
	for i := 0; i < n; i++ {
		if max[i] > result {
			result = max[i]
		}
	}
	return result
}
