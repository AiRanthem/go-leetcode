package weekly

import "sort"

func maxScore(nums []int) (score int) {
	sort.Ints(nums)
	var prefix int
	for i := len(nums) - 1; i >= 0; i-- {
		prefix += nums[i]
		if prefix > 0 {
			score++
		} else {
			break
		}
	}
	return
}
