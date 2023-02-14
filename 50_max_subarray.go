package leetcode

import "math"

func maxSubArray(nums []int) int {
	result := math.MinInt64
	s := 0
	for _, num := range nums {
		s += num
		result = maxInt(result, s)
		if s < 0 {
			s = 0
		}
	}
	return result
}
