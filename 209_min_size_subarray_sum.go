package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	i, sum := 0, 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		if sum >= target {
			for sum >= target {
				sum -= nums[i]
				i++
			}
			if length := j - i + 2; length < result {
				result = length
			}
		}
	}
	if result != math.MaxInt {
		return result
	} else {
		return 0
	}
}
