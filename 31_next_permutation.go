package leetcode

// a b 1 2 3
//

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i, j = i-1, j-1
	}
	if i >= 0 {
		for nums[k] <= nums[i] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	} else {
		j = 0
	}
	k = len(nums) - 1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j, k = j+1, k-1
	}
}
